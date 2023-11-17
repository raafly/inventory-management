package middleware

import (
  "net/http"
  "github.com/raafly/inventory-management/internal/listing"
  "github.com/raafly/inventory-management/pkg/helper"
)

type AuthMiddleware struct {
  Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  _, err := r.Cookie("auth")
  if err != nil {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusUnauthorized)

      webResponse := listing.WebResponse {
        Code: http.StatusUnauthorized,
      }

      helper.WriteToRequestBody(w, webResponse)
  } else {
    m.Handler.ServeHTTP(w, r)
  }
}  
