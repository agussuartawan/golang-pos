package response

import "net/http"

type Response struct {
	Status     int                 `json:"status"`
	Message    string              `json:"message"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
	Data       interface{}         `json:"data,omitempty"`
	Error      *Error              `json:"error,omitempty"`
}

type Error struct {
	Message    string                    `json:"message"`
	Validation []ValidationFailsResponse `json:"validation,omitempty"`
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
