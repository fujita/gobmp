package bgp

import (
	"reflect"
	"testing"
)

func TestUnmarshaBaseAttributes(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		expect *BaseAttributes
	}{
		{
			name:  "panic 1",
			input: []byte{0x40, 0x01, 0x01, 0x00, 0x40, 0x02, 0x20, 0x02, 0x06, 0x00, 0x00, 0x88, 0x38, 0x00, 0x00, 0x9a, 0x6d, 0x00, 0x00, 0x19, 0x35, 0x00, 0x00, 0x0a, 0x7f, 0x00, 0x00, 0x65, 0x20, 0x00, 0x00, 0x53, 0x4e, 0x01, 0x01, 0x00, 0x00, 0x12, 0xc9, 0x40, 0x03, 0x04, 0xc2, 0x1c, 0x62, 0x25, 0x80, 0x04, 0x04, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x07, 0x08, 0x00, 0x00, 0x65, 0x20, 0xc0, 0x78, 0x51, 0x88, 0xc0, 0x08, 0x18, 0x00, 0x00, 0x9a, 0x6d, 0x19, 0x35, 0x00, 0x56, 0x19, 0x35, 0x0b, 0xb8, 0x19, 0x35, 0x0c, 0x1c, 0x19, 0x35, 0x0c, 0x1e, 0x9a, 0x6d, 0xc2, 0x02, 0xc0, 0x20, 0x30, 0x00, 0x00, 0x88, 0x38, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0xd3, 0x00, 0x00, 0x88, 0x38, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x88, 0x38, 0x00, 0x00, 0x00, 0x64, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x88, 0x38, 0x00, 0x00, 0x00, 0x7a, 0x00, 0x00, 0x00, 0x01},
			expect: &BaseAttributes{
				BaseAttrHash:    "354de7a76afc292b187aa3ea32aa76b9",
				Origin:          "igp",
				ASPath:          []uint32{34872, 39533, 6453, 2687, 25888, 21326, 4809},
				ASPathCount:     7,
				Nexthop:         "194.28.98.37",
				Aggregator:      []byte{0, 0, 101, 32, 192, 120, 81, 136},
				CommunityList:   "0:39533, 6453:86, 6453:3000, 6453:3100, 6453:3102, 39533:49666",
				LgCommunityList: "34872:10:211, 34872:11:1, 34872:100:49, 34872:122:1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalBGPBaseAttributes(tt.input)
			if err != nil {
				t.Fatalf("expected to succeed but failed with error: %+v", err)
			}
			if err == nil {
				if !reflect.DeepEqual(got, tt.expect) {
					t.Errorf("Expected extCommunity %+v does not match to actual extCommunity %+v", tt.expect, got)
				}
			}
		})
	}
}

func TestUnmarshalASPath(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		asPath []uint32
	}{
		{
			name:   "panic #1",
			input:  []byte{0x02, 0x08, 0x00, 0x00, 0x24, 0x58, 0x00, 0x00, 0x92, 0x5c, 0x00, 0x00, 0xf1, 0x88, 0x00, 0x04, 0x03, 0xb8, 0x00, 0x00, 0x6e, 0xd0, 0x00, 0x04, 0x03, 0xb8, 0x00, 0x00, 0x6e, 0xd0, 0x00, 0x04, 0x03, 0xb8},
			asPath: []uint32{9304, 37468, 61832, 263096, 28368, 263096, 28368, 263096},
		},
		{
			name:   "panic #2",
			input:  []byte{0x02, 0x48, 0x00, 0x00, 0xce, 0x89, 0x00, 0x00, 0x32, 0x9c, 0x00, 0x00, 0xf0, 0x1c, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0x14},
			asPath: []uint32{52873, 12956, 61468, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 269844},
		},
		{
			name:   "panic #3",
			input:  []byte{0x02, 0xa2, 0x00, 0x00, 0xbe, 0xb5, 0x00, 0x03, 0x21, 0x38, 0x00, 0x00, 0xc5, 0xc5, 0x00, 0x00, 0x00, 0xae, 0x00, 0x00, 0x6c, 0x66, 0x00, 0x00, 0xf0, 0x1c, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0xfa, 0x00, 0x04, 0x1e, 0x14},
			asPath: []uint32{48821, 205112, 50629, 174, 27750, 61468, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 270074, 269844},
		},
		{
			name:   "panic #4",
			input:  []byte{0x02, 0x06, 0x00, 0x00, 0x88, 0x38, 0x00, 0x00, 0x9a, 0x6d, 0x00, 0x00, 0x19, 0x35, 0x00, 0x00, 0x0a, 0x7f, 0x00, 0x00, 0x65, 0x20, 0x00, 0x00, 0x53, 0x4e, 0x01, 0x01, 0x00, 0x00, 0x12, 0xc9},
			asPath: []uint32{34872, 39533, 6453, 2687, 25888, 21326, 4809},
		},
		{
			name:   "1 AS4 segment",
			input:  []byte{0x02, 0x01, 0x00, 0x00, 0x88, 0x38},
			asPath: []uint32{34872},
		},
		{
			name:   "1 AS2 segment",
			input:  []byte{0x02, 0x01, 0x88, 0x38},
			asPath: []uint32{34872},
		},
		{
			name:   "2 AS4 segments",
			input:  []byte{0x02, 0x01, 0x00, 0x00, 0x88, 0x38, 0x01, 0x01, 0x00, 0x00, 0x88, 0x38},
			asPath: []uint32{34872, 34872},
		},
		{
			name:   "2 AS2 segments",
			input:  []byte{0x02, 0x01, 0x88, 0x38, 0x01, 0x01, 0x88, 0x38},
			asPath: []uint32{34872, 34872},
		},
	}
	for _, tt := range tests {
		r := unmarshalAttrASPath(tt.input)
		if !reflect.DeepEqual(tt.asPath, r) {
			t.Fatalf("expected %+v and result %+v as path do not match", tt.asPath, r)
		}
	}
}
