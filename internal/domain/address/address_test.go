package address

import "testing"

func TestNew(t *testing.T) {
	t.Parallel()

	valid := "0xc94770007dda54cF92009BFF0dE90c06F603a09f"
	if _, err := New(valid); err != nil {
		t.Fatalf("expected valid address, got error: %v", err)
	}

	invalid := "0x123"
	if _, err := New(invalid); err == nil {
		t.Fatal("expected invalid address error")
	}
}
