package bgp

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestUnmarshalMPReachNLRI(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		expect *MPReachNLRI
		srv6   bool
	}{
		{
			name:  "issue_173",
			input: []byte{0x00, 0x02, 0x01, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0x0A, 0x98, 0xB7, 0x0B, 0x00, 0x10, 0x20, 0x01},
			expect: &MPReachNLRI{
				AddressFamilyID:      2,
				SubAddressFamilyID:   1,
				NextHopAddressLength: 16,
				NextHopAddress:       []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0x0A, 0x98, 0xB7, 0x0B},
				NLRI:                 []byte{0x10, 0x20, 0x01},
			},
			srv6: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := UnmarshalMPReachNLRI(tt.input, tt.srv6)
			if err != nil {
				t.Fatalf("failed to unmarshal MP Reach NLRI with error: %+v", err)
			}
			if !reflect.DeepEqual(tt.expect, actual) {
				t.Logf("differences: %+v", deep.Equal(tt.expect, actual))
				t.Fatal("the expected object does not match the actual")
			}
		})
	}
}
