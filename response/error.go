package response

type ErrorResponse struct {
	Error string `json:"error" xml:"error" form:"error"`
}