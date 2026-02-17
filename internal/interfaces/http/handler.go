package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fabio-lima/go-api/internal/application"
)

type Handler struct {
	getBalanceUseCase   *application.GetBalanceUseCase
	getHistoryUseCase   *application.GetRequestHistoryUseCase
	defaultHistoryLimit int
}

func NewHandler(
	getBalanceUseCase *application.GetBalanceUseCase,
	getHistoryUseCase *application.GetRequestHistoryUseCase,
	defaultHistoryLimit int,
) *Handler {
	return &Handler{
		getBalanceUseCase:   getBalanceUseCase,
		getHistoryUseCase:   getHistoryUseCase,
		defaultHistoryLimit: defaultHistoryLimit,
	}
}

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	address := r.PathValue("address")
	blockTag := r.URL.Query().Get("block")
	if blockTag == "" {
		blockTag = "latest"
	}

	result, err := h.getBalanceUseCase.Execute(r.Context(), application.GetBalanceInput{
		Address:  address,
		BlockTag: blockTag,
	})
	if err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, application.ErrInvalidAddress):
			status = http.StatusBadRequest
		case errors.Is(err, application.ErrUpstreamFailure):
			status = http.StatusBadGateway
		}
		writeJSON(w, status, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"balance": result.BalanceETH})
}

func (h *Handler) GetRequestHistory(w http.ResponseWriter, r *http.Request) {
	result, err := h.getHistoryUseCase.Execute(r.Context(), application.GetRequestHistoryInput{
		Limit: h.defaultHistoryLimit,
	})
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to list request history"})
		return
	}
	writeJSON(w, http.StatusOK, result.Records)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
