package response

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

func OK(data interface{}) Response {
	return Response{
		Status: http.StatusOK,
		Message: "Success",
		Data: data,
	}
}