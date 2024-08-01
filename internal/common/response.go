package common

import (
	"encoding/json"
	"time"
)

type responseBase struct {
	Status    int    `json:"status"`
	Timestamp string `json:"timestamp"`
}

type ResponseSucess[T any] struct {
	responseBase
	Data T `json:"data"`
}

type ResponseSucessString struct {
  responseBase
  Message string `json:"message"`
}

type ResponseError struct {
	responseBase
	Error string `json:"error"`
}

func ToJsonSucessString(status int, data string) string {
	makeResponse := ResponseSucessString{
		responseBase: responseBase{
			Status:    status,
			Timestamp: time.Now().String(),
		},
		Message: data,
	}

	return toJson(makeResponse)
}
func ToJsonSucess[T any](status int, data T) string {
	makeResponse := ResponseSucess[T]{
		responseBase: responseBase{
			Status:    status,
			Timestamp: time.Now().String(),
		},
		Data: data,
	}

	return toJson(makeResponse)
}

func ToJsonError(status int, error string) string {
	makeResponse := ResponseError{
		responseBase: responseBase{
			Status:    status,
			Timestamp: time.Now().String(),
		},
		Error: error,
	}

	return toJson(makeResponse)
}

func toJson[T any](data T) string {
	responseJson, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return string(responseJson)
}
