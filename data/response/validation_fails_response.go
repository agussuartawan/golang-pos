package response

type ValidationFailsResponse struct {
	Field string `json:"field"`
	Message string `json:"message"`
}