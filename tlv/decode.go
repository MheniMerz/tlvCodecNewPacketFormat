package tlv

import (
	"errors"
	"log"
	"ndn-router/nfd/tlv/name"
	"ndn-router/nfd/tlv/packets"
)

// delegates the work to the appropriate decoder
func Decode(packet []byte) packets.NdnPacket {
	t, _, _, _ := TlvFromBytes(packet)
	switch t.T {
	case INTEREST:
		resultInterest, _ := decodeInterest(t)
		resultInterest.Setbuffer(packet)
		return resultInterest
	case DATA:
		resultData, _ := decodeData(t)
		resultData.Setbuffer(packet)
		return resultData
	case NACK:
		resultNack, _ := decodeNack(t)
		resultNack.Setbuffer(packet)
		return resultNack
	default:
		log.Println("unknown bytes")
		return nil
	}
}



// a decoder is any function with this prototype
type decoder func(packet interface{}, tlvs []Tlv) ([]Tlv, error)

//+++++++++++++++++++++++++++++++++++++++
//add concurrency here to have concurrency on the same packet
//+++++++++++++++++++++++++++++++++++++++
func decodeTlvs(packet interface{}, tlvs []Tlv, decoders ...decoder) error {
	for _, decoder := range decoders {
		tlvs, _ = decoder(packet, tlvs)
	}
	return nil
}


//get name of either interest or data
func decodeName(t Tlv) (name.Name, error) {
	if t.T != NAME {
		return nil, errors.New("--- Decode Name --- : unexpected type")
	}
	//since the name tlv is multi level we do the same as we did with the outer most tlv
	componentTlvs, err := ParseTlvsFromBytes(t.V) // from []bytes to []Tlv
	if err != nil {
		return nil, err
	}
	components := []name.Component{}
	for _, t := range componentTlvs {
		c, err := decodeNameComponent(t)
		if err != nil {
			return nil, err
		}
		components = append(components, c)
	}
	return name.NewName(components...), nil
}

func decodeNameComponent(t Tlv) (name.Component, error) {
	if t.T != NAME_COMPONENT {
		return name.Component(""), errors.New("--- Decode Name --- : unexpected type")
	}
	//take the value which is bytes and turn it into a component
	c := name.ComponentFromBytes(t.V)
	return c, nil
}