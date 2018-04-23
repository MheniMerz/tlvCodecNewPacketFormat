package tlv

import (
	"bytes"
	"errors"
	"io"

	"ndn-router/nfd/tlv/name"
	"ndn-router/nfd/tlv/packets"
)

//reads an NDNpacket and writes a stream of bytes to the writer
//provided as a second parameter
func Encode(packet packets.NdnPacket, byteStream io.Writer) error {
	t := Tlv{}
	switch packet.PacketType() {
	case INTEREST:
		t, _ = encodeInterest(packet.(packets.Interest))
	case DATA:
		t, _ = encodeData(packet.(packets.Data))
	default:
		return errors.New("Encode: -- unknown packet type --")
	}

	//write the tlv as bytes
	err := TlvToBytes(t, byteStream)
	return err
}

func encodeSubTlvs(packet interface{}, enc ...encoder) ([]Tlv, error) {
	//create a buffer here then get thte bytes back
	var t []Tlv
	var err error
	for _, e := range enc {
		t, err = e(packet, t)
		if err != nil {
			return t, err
		}
	}
	return t, nil
}

//an encode is any function of this format
type encoder func(packet interface{}, t []Tlv) ([]Tlv, error)

//encoding a name ==> goes back to encoding nae components
func encodeName(n name.Name) Tlv {
	t := []Tlv{}
	for _, comp := range n {
		t = append(t, encodeNameComponent(comp))
	}
	//need to convert []Tlv to []byte which will be the value of the name
	var b bytes.Buffer
	TlvsToBytes(t, &b)
	result := Tlv{
		T: NAME,
		L: uint64(b.Len()),
		V: b.Next(b.Len()),
	}
	return result
}

func encodeNameComponent(comp name.Component) Tlv {
	b := comp.ComponentToBytes()
	t := Tlv{
		T: NAME_COMPONENT,
		L: uint64(len(b)),
		V: b,
	}
	return t
}
