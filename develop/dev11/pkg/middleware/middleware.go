package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Print(
			"Method: ", r.Method,
			"; Path: ", r.URL.EscapedPath(),
			"; Duration: ", time.Since(start), "\n")
	})
}
