package response

import "net/http"

type Response struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Page      int         `json:"page,omitempty"`
	Limit     int         `json:"limit,omitempty"`
	TotalData *int64      `json:"totalData,omitempty"`
	TotalPage *int        `json:"totalPage,omitempty"`
	NextPage  *int        `json:"nextPage,omitempty"`
	PrevPage  *int        `json:"prevPage,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
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
