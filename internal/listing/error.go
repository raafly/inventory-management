package listing

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/raafly/inventory-management/pkg/helper"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok:= err.(NotFoundError)	
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := ErrorResponse{
			Code: http.StatusBadRequest,
			Message: exception.Error,
		}

		helper.WriteToRequestBody(w, webResponse)
		
		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := ErrorResponse{
			Code: http.StatusBadRequest,
			Message: exception.Error(),
		}

		helper.WriteToRequestBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := ErrorResponse {
		Code:   http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
	}

	helper.WriteToRequestBody(writer, webResponse)
}