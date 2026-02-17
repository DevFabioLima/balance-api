package application

import (
	"context"
	"errors"
	"testing"

	"github.com/fabio-lima/go-api/internal/ports"
)

type mockEthereumClient struct {
	result string
	err    error
}

func (m *mockEthereumClient) GetBalance(_ context.Context, _ string, _ string) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.result, nil
}

type mockHistoryRepository struct {
	records []ports.RequestRecord
}

func (m *mockHistoryRepository) Save(_ context.Context, record ports.RequestRecord) error {
	m.records = append(m.records, record)
	return nil
}

func (m *mockHistoryRepository) List(_ context.Context, _ int) ([]ports.RequestRecord, error) {
	return m.records, nil
}

func TestGetBalanceUseCase_Execute(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		repo := &mockHistoryRepository{}
		uc := NewGetBalanceUseCase(
			&mockEthereumClient{result: "0xde0b6b3a7640000"},
			repo,
			"latest",
		)

		out, err := uc.Execute(context.Background(), GetBalanceInput{
			Address: "0xc94770007dda54cF92009BFF0dE90c06F603a09f",
		})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if out.BalanceETH != "1" {
			t.Fatalf("expected balance 1, got %s", out.BalanceETH)
		}
		if len(repo.records) != 1 || repo.records[0].Status != "success" {
			t.Fatalf("expected one successful history record, got %+v", repo.records)
		}
	})

	t.Run("invalid address", func(t *testing.T) {
		t.Parallel()
		uc := NewGetBalanceUseCase(
			&mockEthereumClient{result: "0x0"},
			&mockHistoryRepository{},
			"latest",
		)
		_, err := uc.Execute(context.Background(), GetBalanceInput{Address: "invalid"})
		if !errors.Is(err, ErrInvalidAddress) {
			t.Fatalf("expected ErrInvalidAddress, got %v", err)
		}
	})

	t.Run("upstream error", func(t *testing.T) {
		t.Parallel()
		uc := NewGetBalanceUseCase(
			&mockEthereumClient{err: errors.New("boom")},
			&mockHistoryRepository{},
			"latest",
		)
		_, err := uc.Execute(context.Background(), GetBalanceInput{
			Address: "0xc94770007dda54cF92009BFF0dE90c06F603a09f",
		})
		if !errors.Is(err, ErrUpstreamFailure) {
			t.Fatalf("expected ErrUpstreamFailure, got %v", err)
		}
	})
}
