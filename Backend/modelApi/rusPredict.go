package modelApi

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hacathon/model"
	"log"

	pb "hacathon/internal/api"
)

type Predictor struct {
	host string
	port string
	Client *pb.CarDetectorClient
}

func NewPredictor(host, port string) *Predictor {
	predictor := &Predictor{
		host:   host,
		port:   port,
	}

	return predictor
}

func (p *Predictor) CarDetector(image string) (model.CarResponse, error) {
	address := fmt.Sprintf("%s:%s", p.host, p.port)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		return model.CarResponse{}, err
	}
	defer conn.Close()

	client := pb.NewCarDetectorClient(conn)

	imageGrpc := &pb.Image{Image: image}

	classes, err := client.Predict(context.Background(), imageGrpc)
	if err != nil {
		log.Printf("Error to get predict from best model: %s", err)
		return model.CarResponse{}, model.ValidationError
	}

	return model.ModelResponseToCarResponse(classes), nil
}
