package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Request struct{
	Content string `json:"content"`
}

func ReadRequestBodyJson(r *http.Request, jsonReq interface{}) (interface{}, func() error, error) {
	err := json.NewDecoder(r.Body).Decode(jsonReq)
	if err != nil {
		return nil, nil, err
	}

	return jsonReq, r.Body.Close, nil
}

func NewRecognitionReader(photo string) *strings.Reader {
	photoJson := fmt.Sprintf("{\"content\":\"%s\"}", photo)
	reader := strings.NewReader(photoJson)

	return reader
}
