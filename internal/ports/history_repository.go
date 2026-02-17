package ports

import (
	"context"
	"time"
)

type RequestRecord struct {
	Address     string    `json:"address"`
	BalanceETH  string    `json:"balance"`
	BlockTag    string    `json:"block"`
	RequestedAt time.Time `json:"requestedAt"`
	Status      string    `json:"status"`
	Error       string    `json:"error,omitempty"`
}

type HistoryRepository interface {
	Save(ctx context.Context, record RequestRecord) error
	List(ctx context.Context, limit int) ([]RequestRecord, error)
}
