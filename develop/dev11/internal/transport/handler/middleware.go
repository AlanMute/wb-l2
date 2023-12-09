package handler

import (
	"log"
	"net/http"
)

func (h *Handler) MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		log.Printf("%s %s %s Body: %s", r.Method, r.RemoteAddr, r.URL.Path, r.Form)
	}
}
