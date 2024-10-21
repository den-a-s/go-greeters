package middleware

import (
	"log"
	"net/http"
)

type LogHandler struct{
	h http.Handler
}

func NewLogMux(h http.Handler) http.Handler {
	return &LogHandler{h: h}
}

func (h *LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("got somerequest")
	h.h.ServeHTTP(w, r)
	log.Printf("end request")
}