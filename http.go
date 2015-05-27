package gtils

import (
	"net/http"
	"os"
)

// EnsureHTTPS wraps a HTTP handler and ensures it was requested over HTTPS.
//
// If DISABLED_ENSURE_HTTPS is in the environment and set to either "1" or "true", EnsureHTTPS will always pass.
func EnsureHTTPS(handler http.HandlerFunc) http.HandlerFunc {
	tmp := os.Getenv("DISABLE_ENSURE_HTTPS")
	disabled := tmp == "1" || tmp == "true"
	return func(w http.ResponseWriter, r *http.Request) {
		if !disabled && (r.URL.Scheme != "https" && r.Header.Get("X-Forwarded-Proto") != "https") {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		handler(w, r)
	}
}
