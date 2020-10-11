package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"hacathon/model"
	"hacathon/modelApi"
	"hacathon/tools"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	host = "0.0.0.0"
	port = "8081"

	grpcHost = "localhost"
	grpcPort = "50051"

	predictURL = "https://gw.hackathon.vtb.ru/vtb/hackathon/car-recognize"
	getCarList = "https://gw.hackathon.vtb.ru/vtb/hackathon/marketplace"
	calculqteURL = "https://gw.hackathon.vtb.ru/vtb/hackathon/calculate"
	paymentGraphURL = "https://gw.hackathon.vtb.ru/vtb/hackathon/payments-graph"
	settingsURL = "https://gw.hackathon.vtb.ru/vtb/hackathon/settings"
	carloanURL = "https://gw.hackathon.vtb.ru/vtb/hackathon/carloan"
)

type Server struct {
	carList *model.CarListInternal
	Predictor *modelApi.Predictor
}

func newServer() *Server {
	predictor := modelApi.NewPredictor(grpcHost, grpcPort)
	return &Server{Predictor: predictor}
}

func (s *Server) getCarByFilter(key string) (model.ReturnCar, bool) {
	keys := strings.Split(strings.ToUpper(key), " ")
	if len(keys) < 2 {
		return model.ReturnCar{}, false
	}

	car, ok := s.carList.Cars[keys[0]]
	if !ok {
		return model.ReturnCar{}, false
	}

	carModel, ok := car.Models[keys[1]]
	if !ok {
		return model.ReturnCar{}, false
	}

	modelTitle := strings.ToUpper(carModel.Title)
	if strings.Index(modelTitle, "СЕРИИ") != -1 || strings.Index(modelTitle, "СЕРИЯ") != -1 {
		modelTitle = strings.Split(modelTitle, " ")[0]
	}

	returnCar := model.ReturnCar{
		Absentee: car.Absentee,
		Alias:    car.Alias,
		Country:  car.Country,
		Logo:     car.Logo,
		Models:   carModel,
		Title:    fmt.Sprintf("%s %s %s", car.Title, carModel.Prefix, carModel.Title),
		TitleRus: car.TitleRus,
		Info: model.CarI.Info[strings.ToUpper(fmt.Sprintf("%s %s", car.Title, modelTitle))],
	}

	return returnCar, true
}

func (s *Server) createRequest(url string, method string, reader *strings.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, reader)
	s.addHeader(req)

	return req
}

func (s *Server) constructorURL(url string, params map[string]string) string {
	var paramsString []string

	for key, value := range params {
		paramsString = append(paramsString, fmt.Sprintf("%s=%s", key, value))
	}

	query := fmt.Sprintf("%s?%s", url, strings.Join(paramsString, "&"))

	return query
}

func (s *Server) predictModel(w http.ResponseWriter, r *http.Request) {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &tools.Request{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(500)
		return
	}

	jsonReq := body.(*tools.Request)
	defer closeFunc()

	reader := tools.NewRecognitionReader(jsonReq.Content)
	req := s.createRequest(predictURL, "POST", reader)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// TODO: check code
		log.Printf("Error to get predict: %s\n", err)
		w.WriteHeader(500)
		return
	}

	if res.StatusCode != 200 {
		w.WriteHeader(res.StatusCode)
		return
	}

	carResp := &model.CarResponse{}
	if err = json.NewDecoder(res.Body).Decode(carResp); err != nil {
		log.Printf("Error to parse json predict: %s\n", err)
		w.WriteHeader(500)
		return
	}

	defer res.Body.Close()

	predictCar := model.GetBestProb(carResp)
	filteredCars, ok := s.getCarByFilter(predictCar)
	if !ok {
		w.WriteHeader(404)
		return
	}
	cars := model.PredictToiOS(&filteredCars)
	response, _ := json.Marshal(cars)

	w.Write(response)
}

func (s *Server) bestPredictModel(w http.ResponseWriter, r *http.Request) {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &tools.Request{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(400)
		w.Write(model.NewJsonErrorMessage("Smth wrong"))
		return
	}

	jsonReq := body.(*tools.Request)
	defer closeFunc()

	res, err := s.Predictor.CarDetector(jsonReq.Content)
	if err != nil {
		if errors.Is(err, model.ValidationError) {
			w.WriteHeader(422)
			w.Write(model.NewJsonErrorMessage(err.Error()))
		}
		w.WriteHeader(400)
		w.Write(model.NewValidationErrorMessageJson(err.Error()))
	}

	resp, _ := json.Marshal(res)

	w.Write(resp)
}

func (s *Server) calculate(w http.ResponseWriter, r *http.Request)  {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &model.CalculateRequest{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(500)
		return
	}

	defer closeFunc()

	jsonReq := body.(*model.CalculateRequest)
	jsonData, err := json.Marshal(jsonReq)
	if err != nil {
		log.Printf("Can`t marshal json body: %s", err)
		w.WriteHeader(500)
		return
	}

	reader := strings.NewReader(string(jsonData))
	req := s.createRequest(calculqteURL, "POST", reader)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// TODO: check code
		log.Printf("Error to get predict: %s\n", err)
		w.WriteHeader(500)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		w.WriteHeader(res.StatusCode)
		return
	}

	calcResp := &model.CalculateResponse{}
	if err = json.NewDecoder(res.Body).Decode(calcResp); err != nil {
		log.Printf("Error to parse json predict: %s\n", err)
		w.WriteHeader(500)
		return
	}

	calciOS := model.CalcResponseToCalciOS(*calcResp)

	response, _ := json.Marshal(calciOS)

	w.Write(response)
}

func (s *Server) paymentGraph(w http.ResponseWriter, r *http.Request) {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &model.PaymentRequest{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(500)
		return
	}

	defer closeFunc()

	jsonReq := body.(*model.PaymentRequest)
	jsonData, err := json.Marshal(jsonReq)
	if err != nil {
		log.Printf("Can`t marshal json body: %s", err)
		w.WriteHeader(500)
		return
	}

	reader := strings.NewReader(string(jsonData))
	req := s.createRequest(paymentGraphURL, "POST", reader)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// TODO: check code
		log.Printf("Error to get predict: %s\n", err)
		w.WriteHeader(500)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		w.WriteHeader(res.StatusCode)
		return
	}

	payResp := &model.PaymentResponse{}
	if err = json.NewDecoder(res.Body).Decode(payResp); err != nil {
		log.Printf("Error to parse json predict: %s\n", err)
		w.WriteHeader(500)
		return
	}

	payiOS := model.PaymentRespToPaymentiOS(*payResp)
	response, _ := json.Marshal(payiOS)

	w.Write(response)
}

func (s *Server) getSettings(w http.ResponseWriter, r *http.Request) {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &model.CalculateRequest{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(500)
		return
	}

	defer closeFunc()

	jsonReq := body.(*model.CalculateRequest)
	jsonData, err := json.Marshal(jsonReq)
	if err != nil {
		log.Printf("Can`t marshal json body: %s", err)
		w.WriteHeader(500)
		return
	}

	reader := strings.NewReader(string(jsonData))

	query := r.URL.Query()
	queryString := s.constructorURL(settingsURL,
		map[string]string{"name": query.Get("name"), "language": query.Get("language")})
	req := s.createRequest(queryString, "GET", reader)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// TODO: check code
		log.Printf("Error to get predict: %s\n", err)
		w.WriteHeader(500)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		w.WriteHeader(res.StatusCode)
		return
	}

	settingsResp := &model.SettingsResponse{}
	if err = json.NewDecoder(res.Body).Decode(settingsResp); err != nil {
		log.Printf("Error to parse json predict: %s\n", err)
		w.WriteHeader(500)
		return
	}

	response, _ := json.Marshal(settingsResp)

	w.Write(response)
}

func (s *Server) carLoan(w http.ResponseWriter, r *http.Request) {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &model.CalculateRequest{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(500)
		return
	}

	defer closeFunc()

	jsonReq := body.(*model.CarLoanRequest)
	jsonData, err := json.Marshal(jsonReq)
	if err != nil {
		log.Printf("Can`t marshal json body: %s", err)
		w.WriteHeader(500)
		return
	}

	reader := strings.NewReader(string(jsonData))
	req := s.createRequest(carloanURL, "POST", reader)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// TODO: check code
		log.Printf("Error to get predict: %s\n", err)
		w.WriteHeader(500)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		w.WriteHeader(res.StatusCode)
		return
	}

	loanResp := &model.CarLoanResponse{}
	if err = json.NewDecoder(res.Body).Decode(loanResp); err != nil {
		log.Printf("Error to parse json predict: %s\n", err)
		w.WriteHeader(500)
		return
	}

	response, _ := json.Marshal(loanResp)

	w.Write(response)
}

func (s *Server) addHeader(req *http.Request) {
	req.Header.Add("x-ibm-client-id", "1823d72e90f4be9a700ee2619a21b37f")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")
}

func (s *Server) downloadCarList() {
	req, _ := http.NewRequest("GET", getCarList, nil)
	s.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		log.Printf("Error to read body car list: %s\n", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error to read body car list: %s\n", err)
		return
	}

	var list model.CarList
	if err := json.Unmarshal(body, &list); err != nil {
		log.Printf("Error to unmarshal car list: %s\n", err)
	}

	s.carList = model.CarToInternal(list)
	jsn, _ := json.Marshal(s.carList.Cars)
	println(jsn)
	println("Ok")
}

func initRouter(r *mux.Router, s *Server) {
	//r.HandleFunc("/marketplace", s.getCarList)
	r.HandleFunc("/recognition", s.predictModel)
	r.HandleFunc("/calculate", s.calculate)
	r.HandleFunc("/payments-graph", s.paymentGraph)
	r.HandleFunc("/settings", s.getSettings)
	r.HandleFunc("/carloan", s.carLoan)
	r.HandleFunc("/bestrecog", s.bestPredictModel)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func startServer(r *mux.Router) {
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%s", host,  port),
	}
	log.Printf("Server run on http://%s:%s\n", host, port)

	log.Fatal(srv.ListenAndServe())
}

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	s := newServer()
	s.downloadCarList()

	model.InitCarInfo()

	initRouter(r, s)
	startServer(r)
}

//decPhoto, err := base64.StdEncoding.DecodeString(photo)
//photoEncoded := base64.StdEncoding.EncodeToString(decPhoto)
