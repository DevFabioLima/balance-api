package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/fabio-lima/go-api/internal/domain/address"
	"github.com/fabio-lima/go-api/internal/domain/balance"
	"github.com/fabio-lima/go-api/internal/ports"
)

var (
	ErrInvalidAddress  = errors.New("invalid address")
	ErrUpstreamFailure = errors.New("upstream failure")
)

type GetBalanceUseCase struct {
	ethClient       ports.EthereumClient
	historyRepo     ports.HistoryRepository
	defaultBlockTag string
}

type GetBalanceInput struct {
	Address  string
	BlockTag string
}

type GetBalanceOutput struct {
	BalanceETH string
}

func NewGetBalanceUseCase(
	ethClient ports.EthereumClient,
	historyRepo ports.HistoryRepository,
	defaultBlockTag string,
) *GetBalanceUseCase {
	return &GetBalanceUseCase{
		ethClient:       ethClient,
		historyRepo:     historyRepo,
		defaultBlockTag: defaultBlockTag,
	}
}

func (u *GetBalanceUseCase) Execute(ctx context.Context, input GetBalanceInput) (GetBalanceOutput, error) {
	blockTag := input.BlockTag
	if blockTag == "" {
		blockTag = u.defaultBlockTag
	}

	validAddress, err := address.New(input.Address)
	if err != nil {
		u.saveHistory(ctx, ports.RequestRecord{
			Address:     input.Address,
			BlockTag:    blockTag,
			RequestedAt: time.Now().UTC(),
			Status:      "error",
			Error:       ErrInvalidAddress.Error(),
		})
		return GetBalanceOutput{}, ErrInvalidAddress
	}

	hexWei, err := u.ethClient.GetBalance(ctx, validAddress.String(), blockTag)
	if err != nil {
		u.saveHistory(ctx, ports.RequestRecord{
			Address:     validAddress.String(),
			BlockTag:    blockTag,
			RequestedAt: time.Now().UTC(),
			Status:      "error",
			Error:       err.Error(),
		})
		return GetBalanceOutput{}, fmt.Errorf("%w: %v", ErrUpstreamFailure, err)
	}

	balanceETH, err := balance.WeiHexToETHString(hexWei)
	if err != nil {
		u.saveHistory(ctx, ports.RequestRecord{
			Address:     validAddress.String(),
			BlockTag:    blockTag,
			RequestedAt: time.Now().UTC(),
			Status:      "error",
			Error:       err.Error(),
		})
		return GetBalanceOutput{}, fmt.Errorf("%w: %v", ErrUpstreamFailure, err)
	}

	u.saveHistory(ctx, ports.RequestRecord{
		Address:     validAddress.String(),
		BalanceETH:  balanceETH,
		BlockTag:    blockTag,
		RequestedAt: time.Now().UTC(),
		Status:      "success",
	})

	return GetBalanceOutput{BalanceETH: balanceETH}, nil
}

func (u *GetBalanceUseCase) saveHistory(ctx context.Context, record ports.RequestRecord) {
	if u.historyRepo == nil {
		return
	}
	_ = u.historyRepo.Save(ctx, record)
}
