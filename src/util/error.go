package util

import (
	"fmt"

	"github.com/api_rest_go/src/model"
)

func HandlerError(e *error, code int, message string) model.Response {
	fmt.Println(e)

	response := model.Response{Code: code, Message: message}

	return response
}
