package address

import (
	"fmt"
	"regexp"
)

var ethAddressRegex = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)

type Address struct {
	value string
}

func New(value string) (Address, error) {
	if !ethAddressRegex.MatchString(value) {
		return Address{}, fmt.Errorf("invalid ethereum address")
	}
	return Address{value: value}, nil
}

func (a Address) String() string {
	return a.value
}
