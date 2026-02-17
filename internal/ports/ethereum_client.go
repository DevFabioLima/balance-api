package ports

import "context"

type EthereumClient interface {
	GetBalance(ctx context.Context, address string, blockTag string) (string, error)
}
