package middleware

import (
	commonmw "akatm/common/middleware"
	"net/http"
)

type SignCheckMiddleware struct {
	inner func(http.HandlerFunc) http.HandlerFunc
}

func NewSignCheckMiddleware(salt string, skewSeconds int64) *SignCheckMiddleware {
	cmw := commonmw.NewSignCheckMiddleware(salt, skewSeconds)
	return &SignCheckMiddleware{inner: cmw.Handle}
}

func (m *SignCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	if m != nil && m.inner != nil {
		return m.inner(next)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
