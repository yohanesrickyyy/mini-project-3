package dto

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  any    `json:"detail,omitempty"`
}

func (e ErrorResponse) Error() int {
	return e.Code
}
