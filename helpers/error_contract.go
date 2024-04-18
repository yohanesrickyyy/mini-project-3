package helpers

import (
	"mini-project-3/dto"
	"net/http"
)

func ErrBadRequest(detail any) dto.ErrorResponse {
	return dto.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
		Detail:  detail,
	}
}

func ErrUnauthorized(detail any) dto.ErrorResponse {
	return dto.ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized access",
		Detail:  detail,
	}
}

func ErrInternalServer(detail any) dto.ErrorResponse {
	return dto.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
		Detail:  detail,
	}
}

func ErrNotFound(detail any) dto.ErrorResponse {
	return dto.ErrorResponse{
		Code:    http.StatusNotFound,
		Message: "Error not found",
		Detail:  detail,
	}
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
