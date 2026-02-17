package history

import (
	"context"
	"sync"

	"github.com/fabio-lima/go-api/internal/ports"
)

type MemoryRepository struct {
	mu         sync.RWMutex
	maxRecords int
	records    []ports.RequestRecord
}

func NewMemoryRepository(maxRecords int) *MemoryRepository {
	return &MemoryRepository{
		maxRecords: maxRecords,
		records:    make([]ports.RequestRecord, 0, maxRecords),
	}
}

func (r *MemoryRepository) Save(_ context.Context, record ports.RequestRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.records = append([]ports.RequestRecord{record}, r.records...)
	if len(r.records) > r.maxRecords {
		r.records = r.records[:r.maxRecords]
	}
	return nil
}

func (r *MemoryRepository) List(_ context.Context, limit int) ([]ports.RequestRecord, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if limit > len(r.records) {
		limit = len(r.records)
	}
	copied := make([]ports.RequestRecord, limit)
	copy(copied, r.records[:limit])
	return copied, nil
}
