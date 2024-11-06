package response

import "net/http"

type Response struct {
	Status     int                 `json:"status"`
	Message    string              `json:"message"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
	Data       interface{}         `json:"data,omitempty"`
	Errors     interface{}         `json:"errors,omitempty"`
}

type IDResponse struct {
	ID uint `json:"id"`
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
