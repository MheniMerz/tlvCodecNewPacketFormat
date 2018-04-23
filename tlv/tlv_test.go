package tlv

import (
	"bytes"
	"testing"
)

type tlvToBytesTest struct {
	in       Tlv
	expected []byte
}

var tlvToBytesTestTable = []tlvToBytesTest{
	{Tlv{T: 1, L: 1, V: []byte{0x1}}, []byte{0x1, 0x1, 0x1}},
	{Tlv{T: 255, L: 1, V: []byte{0x1}}, []byte{0xFD, 0x00, 0xFF, 0x1, 0x1}},
	{Tlv{T: 1111, L: 2, V: []byte{0x1, 0x1}}, []byte{0xFD, 0x04, 0x57, 0x2, 0x1, 0x1}},
}

func TestTlvToBytes(t *testing.T) {
	var buf bytes.Buffer
	var actual []byte
	for _, tt := range tlvToBytesTestTable {
		//empty the buffer to avoid conflicts
		buf.Reset()
		TlvToBytes(tt.in, &buf)
		actual = buf.Bytes()
		if !SliceEqual(actual, tt.expected) {
			t.Errorf("TlvToBytes(%v): expected %d, actual %d", tt.in, tt.expected, actual)
		}
	}
}

type parseTlvsFromBytesBytesTest struct {
	in       []byte
	expected []Tlv
}

var parseTlvsFromBytesBytesTestTable = []parseTlvsFromBytesBytesTest{
	{
		[]byte{0x1, 0x1, 0x1, 0x1, 0x1, 0x1},
		[]Tlv{
			{T: 1, L: 1, V: []byte{0x1}},
			{T: 1, L: 1, V: []byte{0x1}},
		},
	},
	{
		[]byte{0xFD, 0x00, 0xFF, 0x1, 0x1, 0x1, 0x1, 0x1},
		[]Tlv{
			{T: 255, L: 1, V: []byte{0x1}},
			{T: 1, L: 1, V: []byte{0x1}},
		},
	},

	// {Tlv{T: 255, L: 1, V: []byte{0x1}}, []byte{0xFD, 0x00, 0xFF, 0x1, 0x1}},
	// {Tlv{T: 1111, L: 2, V: []byte{0x1, 0x1}}, []byte{0xFD, 0x04, 0x57, 0x2, 0x1, 0x1}},
}

func TestParseTlvsFromBytes(t *testing.T) {
	for _, tt := range parseTlvsFromBytesBytesTestTable {
		actual, _ := ParseTlvsFromBytes(tt.in)
		for i, j := range actual {
			if j.T != tt.expected[i].T || j.L != tt.expected[i].L || !SliceEqual(j.V, tt.expected[i].V) {
				t.Errorf("TlvToBytes(%v): expected %d, actual %d", tt.in, tt.expected, actual)
			}
		}
	}
}
