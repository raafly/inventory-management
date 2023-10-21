package middleware

import (
	"net/http"

	"github.com/raafly/inventory-management/helper"
	"github.com/raafly/inventory-management/model"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		// ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := model.WebResponse {
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToRequestBody(writer, webResponse)
	}
}