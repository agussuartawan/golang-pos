package response

import "net/http"

type Response struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Page      *int        `json:"page,omitempty"`
	Limit     *int        `json:"limit,omitempty"`
	TotalData *int64      `json:"totalData,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
}

func OK(data interface{}) Response {
	return Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}
}

func Created(data interface{}) Response {
	return Response{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    data,
	}
}
