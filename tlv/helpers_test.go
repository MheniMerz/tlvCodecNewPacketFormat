package tlv

import (
	"bytes"
	"testing"
)

func TestDecodeNonNegativeInteger(t *testing.T) {
	input := []byte{0xFF, 0xFF, 0xFF, 0xFF}
	result := DecodeNonNegativeInteger(input)
	if result == 0 {
		t.Error("Test Failed 'result' ==", result)
	}
}

type varDecodingTest struct {
	n        []byte
	expected uint64
}

var varDecodingTestTable = []varDecodingTest{
	{[]byte{0xFD, 0x1, 0x1}, 257},
	{[]byte{0xFE, 0x1, 0x1, 0x1, 0x1}, 16843009},
	{[]byte{0xFF, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1}, 72340172838076673},
}

func TestVarDecoding(t *testing.T) {
	for _, tt := range varDecodingTestTable {
		actual, _ := varDecoding(tt.n)
		if actual != tt.expected {
			t.Errorf("varDecoding(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func BenchmarkVarDecoding(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range varDecodingTestTable {
			varDecoding(tt.n)
		}
	}
}

type varEncodingTest struct {
	n        uint64
	expected []byte
}

var varEncodingTestTable = []varEncodingTest{
	{257, []byte{0xFD, 0x1, 0x1}},
	{16843009, []byte{0xFE, 0x1, 0x1, 0x1, 0x1}},
	{72340172838076673, []byte{0xFF, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1}},
}

func TestVarEncoding(t *testing.T) {
	for _, tt := range varEncodingTestTable {
		var buf bytes.Buffer
		var actual []byte
		varEncoding(tt.n, &buf)
		actual = buf.Bytes()
		if !SliceEqual(actual, tt.expected) {
			t.Errorf("varEncoding(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func BenchmarkVarEncoding(b *testing.B) {
	var buf bytes.Buffer
	for n := 0; n < b.N; n++ {
		for _, tt := range varEncodingTestTable {
			varEncoding(tt.n, &buf)
		}
	}
}
