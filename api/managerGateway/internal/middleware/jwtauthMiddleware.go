package middleware

import (
	"net/http"
	"strings"
)

type JwtAuthMiddleware struct {
	// ParseFunc 解析并校验token，返回是否通过
	ParseFunc func(token string) bool
}

func NewJwtAuthMiddleware(parse func(string) bool) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{ParseFunc: parse}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "missing or invalid authorization", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		if token == "" || m.ParseFunc == nil || !m.ParseFunc(token) {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
