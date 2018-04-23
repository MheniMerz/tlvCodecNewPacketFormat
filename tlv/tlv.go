package tlv

import (
	"encoding/binary"
	"io"
)

type Tlv struct {
	T uint64
	L uint64
	V []byte
}

//tlv reader, reads a slice of bytes and gives back a tlv
//mostly used for the outer-most tlv (interest, data, nack)
func TlvFromBytes(packet []byte) (result Tlv, err error, ts int, ls int) {
	var t, l uint64
	//get type
	t, ts = varDecoding(packet)
	//get length
	l, ls = varDecoding(packet[ts:])
	result = Tlv{t, l, packet[ts+ls : ts+ls+int(l)]}
	err = nil //need to handle byte reading errors
	return
}

//tlv parser reads a stream of bytes and gives back a slice of tlvs (name, nonce, lifetime ...)
//usually takes the value of the outer most tlv
func ParseTlvsFromBytes(packet []byte) (result []Tlv, err error) {
	tmp := Tlv{}
	var ts, ls int
	for i := 0; i < len(packet); i += ts + ls + int(tmp.L) {
		//get the current tlv
		tmp, err, ts, ls = TlvFromBytes(packet[i:])
		if err != nil {
			return
		}
		result = append(result, tmp) // add the tlv to results
	}
	return
}

//reads a tlv and gives back a byte slice,
// need to pay attention to variable length encoding
// type and length may need to be encoded on multiple bytes

func TlvToBytes(t Tlv, byteStream io.Writer) error {
	varEncoding(t.T, byteStream)
	varEncoding(t.L, byteStream)
	if t.V != nil{
		return binary.Write(byteStream, binary.BigEndian, t.V)
	}
	return nil
}

func TlvsToBytes(t []Tlv, byteStream io.Writer) error {
	for _, tt := range t {
		err := TlvToBytes(tt, byteStream)
		if err != nil {
			return err
		}
	}
	return nil
}

//count tlvs.. not needed for now
func countTlv(packet []byte) (count int) {
	var i int // initilized to the zero value
	for len(packet[i:]) != 0 {
		//get type
		_, j := varDecoding(packet[i:])
		i += j
		//get length
		l, j := varDecoding(packet[i:])
		i += j + int(l) // move the index to the next tlv
		count++
	}
	return
}
