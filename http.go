package gtils

import "net/http"

func EnsureHTTPS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Scheme != "https" && r.Header.Get("X-Forwarded-Proto") != "https" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		handler(w, r)
	}
}
