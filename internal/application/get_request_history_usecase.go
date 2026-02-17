package application

import (
	"context"

	"github.com/fabio-lima/go-api/internal/ports"
)

type GetRequestHistoryUseCase struct {
	historyRepo ports.HistoryRepository
}

type GetRequestHistoryInput struct {
	Limit int
}

type GetRequestHistoryOutput struct {
	Records []ports.RequestRecord
}

func NewGetRequestHistoryUseCase(historyRepo ports.HistoryRepository) *GetRequestHistoryUseCase {
	return &GetRequestHistoryUseCase{historyRepo: historyRepo}
}

func (u *GetRequestHistoryUseCase) Execute(ctx context.Context, input GetRequestHistoryInput) (GetRequestHistoryOutput, error) {
	records, err := u.historyRepo.List(ctx, input.Limit)
	if err != nil {
		return GetRequestHistoryOutput{}, err
	}
	return GetRequestHistoryOutput{Records: records}, nil
}
