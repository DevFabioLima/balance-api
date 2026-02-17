package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/fabio-lima/go-api/internal/application"
	"github.com/fabio-lima/go-api/internal/config"
	historyinfra "github.com/fabio-lima/go-api/internal/infrastructure/history"
	"github.com/fabio-lima/go-api/internal/infrastructure/infura"
	httpiface "github.com/fabio-lima/go-api/internal/interfaces/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	httpClient := &http.Client{Timeout: cfg.RequestTimeout}
	ethClient := infura.NewClient(cfg.InfuraURL, httpClient)
	historyRepo := historyinfra.NewMemoryRepository(cfg.HistoryMaxRecords)

	getBalanceUseCase := application.NewGetBalanceUseCase(
		ethClient,
		historyRepo,
		"latest",
	)
	getHistoryUseCase := application.NewGetRequestHistoryUseCase(historyRepo)

	handler := httpiface.NewHandler(getBalanceUseCase, getHistoryUseCase, cfg.HistoryListLimit)
	router := httpiface.NewRouter(handler)

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("api listening on :%s", cfg.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}

	_ = server.Shutdown(context.Background())
}
