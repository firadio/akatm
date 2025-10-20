package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sort"
	"strings"
	"time"
)

type SignCheckMiddleware struct {
	PublicKey   string
	TimeWindowS int64 // seconds
}

func NewSignCheckMiddleware(publicKey string, timeWindowSeconds int64) *SignCheckMiddleware {
	return &SignCheckMiddleware{PublicKey: publicKey, TimeWindowS: timeWindowSeconds}
}

func (m *SignCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ts := r.Header.Get("X-Timestamp")
		sign := r.Header.Get("X-Sign")
		if ts == "" || sign == "" {
			http.Error(w, "missing sign headers", http.StatusUnauthorized)
			return
		}

		// timestamp check
		ms, err := time.ParseDuration(ts + "ms")
		if err != nil {
			http.Error(w, "invalid timestamp", http.StatusUnauthorized)
			return
		}
		if time.Since(time.Unix(0, 0).Add(ms)) > time.Duration(m.TimeWindowS)*time.Second ||
			time.Until(time.Unix(0, 0).Add(ms)) > time.Duration(m.TimeWindowS)*time.Second {
			http.Error(w, "timestamp expired", http.StatusUnauthorized)
			return
		}

		// collect params (query only for simplicity; body signing can be added)
		params := make([]string, 0)
		for k, v := range r.URL.Query() {
			if len(v) > 0 {
				params = append(params, k+"="+strings.TrimSpace(v[0]))
			}
		}
		sort.Strings(params)
		canonical := strings.Join(params, "&") + m.PublicKey
		h := sha256.Sum256([]byte(canonical))
		calc := strings.ToUpper(hex.EncodeToString(h[:]))
		if calc != sign {
			http.Error(w, "invalid sign", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
