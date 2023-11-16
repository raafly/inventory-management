package middleware

import "net/http"

type AuthMiddleware struct {
  Handler http.Handler
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  r.Get.Cookie("auth")
}
