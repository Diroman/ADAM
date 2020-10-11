package model

import (
	"encoding/json"
	"errors"
)

var ValidationError = errors.New("ValidationError")

type ErrorMessage struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func NewJsonErrorMessage(message string) []byte {
	errMessage := ErrorMessage{
		Error:   true,
		Message: message,
	}
	res, _ := json.Marshal(errMessage)

	return res
}

type ValidationErrorMessage struct {
	Detail []Detail `json:"detail"`
}

type Detail struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

func NewValidationErrorMessageJson(msg string) []byte {
	validMessage := ValidationErrorMessage{
		Detail: []Detail{{
			Loc:  []string{"This is error"},
			Msg:  "This is error",
			Type: "This is type",
		}},
	}
	res, _ := json.Marshal(validMessage)

	return res
}
