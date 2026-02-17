package balance

import (
	"fmt"
	"math/big"
	"strings"
)

var weiPerETH = big.NewInt(1_000_000_000_000_000_000)

func WeiHexToETHString(hexWei string) (string, error) {
	normalized := strings.TrimPrefix(hexWei, "0x")
	if normalized == "" {
		normalized = "0"
	}

	wei := new(big.Int)
	if _, ok := wei.SetString(normalized, 16); !ok {
		return "", fmt.Errorf("invalid wei hex value")
	}

	rat := new(big.Rat).SetFrac(wei, weiPerETH)
	eth := rat.FloatString(18)
	return strings.TrimRight(strings.TrimRight(eth, "0"), "."), nil
}
