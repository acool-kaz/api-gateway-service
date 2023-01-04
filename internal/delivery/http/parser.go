package http

import (
	"encoding/json"
	"net/http"

	parser_pb "github.com/acool-kaz/api-gateway-service/pkg/parser/pb"
)

func (h *Handler) parserHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := h.parserClient.Client.Parser(r.Context(), &parser_pb.ParserRequest{})
	if err != nil {
		h.errorHandler(w, http.StatusBadGateway, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}
}
