package middleware

import "net/http"

type AuthMiddleware struct {
  Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  err := r.Cookie("auth")
  if err != nil {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusUnauthorized)

      webResponse := web.WebResponse {
        Code: http.StatusUnauthorized,
        Status: "UNAUTHORIZED",
      }

      helper.WriteToRequestBody(w, webResponse)
  } else {
    middleware.Handler.ServeHTTP(w, r)
  }
} 
