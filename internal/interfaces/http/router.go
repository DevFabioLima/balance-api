package http

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handler.Health)
	mux.HandleFunc("GET /address/balance/{address}", handler.GetBalance)
	mux.HandleFunc("GET /requests/history", handler.GetRequestHistory)
	return mux
}
