package bgp

import (
	"strings"
	"testing"
)

func TestExtendedCommunity(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		expect string
	}{
		{
			name:   "ext comm 1",
			input:  []byte{0x00, 0x02, 0x00, 0x64, 0x00, 0x00, 0x00, 0x64},
			expect: "rt=100:100",
		},
		{
			name:   "ext comm 2",
			input:  []byte{0x00, 0x02, 0x00, 0x05, 0x00, 0x00, 0xfd, 0xeb},
			expect: "rt=5:65003",
		},
		{
			name:   "type 8 community",
			input:  []byte{0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expect: "flowspec=redirect_to_ip_next_hop",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ext, err := makeExtCommunity(tt.input)
			if err != nil {
				t.Errorf("with error: %+v", err)
			}
			result := ext.String()
			if strings.Compare(tt.expect, result) != 0 {
				t.Errorf("Result %s does not match the expected community: %s", result, tt.expect)
			}
		})
	}
}
