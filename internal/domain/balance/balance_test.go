package balance

import "testing"

func TestWeiHexToETHString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "zero", input: "0x0", want: "0"},
		{name: "one ether", input: "0xde0b6b3a7640000", want: "1"},
		{name: "fraction", input: "0x7be", want: "0.000000000000001982"},
		{name: "invalid", input: "0xZZ", wantErr: true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := WeiHexToETHString(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}
