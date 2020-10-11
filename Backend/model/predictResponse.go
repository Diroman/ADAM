package model

import (
	pb "hacathon/internal/api"
)

type CarResponse struct {
	Probabilities map[string]float32 `json:"probabilities"`
}

func ModelResponseToCarResponse(modelResp *pb.Classes) CarResponse {
	return CarResponse{Probabilities: modelResp.Classes}
}