package http

import (
	"log"
	"net/http"
)

func (h *Handler) errorHandler(w http.ResponseWriter, code int, msg string) {
	log.Println(msg)
	http.Error(w, http.StatusText(code), code)
}
