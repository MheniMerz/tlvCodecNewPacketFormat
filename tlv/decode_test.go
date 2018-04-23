package tlv

import (
	"ndn-router/nfd/tlv/name"
	"ndn-router/nfd/tlv/packets"
	"testing"
)

type decodeTest struct {
	input    []byte
	expected *packets.Interest
}

var n = name.NewName("A")

var decodeTestTable = []decodeTest{
	//{[]byte{0x5, 0x0}, packets.NewInterest(nil)},
	{[]byte{0x5, 0x1, 0x7, 0x3, 0x8, 0x1, 0x65}, packets.NewInterest(n)},
	//{[]byte{0xFF, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x65}, 72340172838076673},
}

func TestDecode(t *testing.T) {
	for _, tt := range decodeTestTable {
		x := Decode(tt.input)
		if x != tt.expected {
			t.Errorf("TestDecode : expected %v, actual %v", tt.expected, x)
		}
	}
}

func TestInterestDecodeName(t *testing.T) {

}

func TestDecodeNameComponent(t *testing.T) {

}
