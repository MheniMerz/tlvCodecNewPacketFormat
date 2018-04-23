package tlv

import (
	"log"
	"errors"

	"ndn-router/nfd/tlv/packets"
)

func decodeNack(t Tlv) (packets.Nack, error) {
	tlvs, _ := ParseTlvsFromBytes(t.V)
	resultNack := packets.Nack{}
	err := decodeTlvs(
		&resultNack, tlvs,
		decodeNackPacket,
		decodeNackReason,
	)

	if err != nil {
		return packets.Nack{}, err
	}
	return resultNack, nil
}


func decodeNackPacket(packet interface{}, tlvs []Tlv) ([]Tlv, error){
	if len(tlvs) < 1 {
		log.Println("decodeNackPacket : --- no tlvs to read ---")
		return nil, errors.New("decodeNackPacket : --- no tlvs to read ---")
	}
	i, _ := decodeInterest(tlvs[0])
	packet.(*packets.Nack).SetPacket(&i)
	return tlvs[1:], nil                     //get rid of the processed tlv		
}

func decodeNackReason(packet interface{}, tlvs []Tlv) ([]Tlv, error){
	if len(tlvs) < 1 {
		log.Println("decodeNackReason : --- no tlvs to read ---")
		return nil, errors.New("decodeNackReason : --- no tlvs to read ---")
	}
	t := tlvs[0]
	reason := DecodeNonNegativeInteger(t.V)
	packet.(*packets.Nack).SetReason(packets.NackReason(reason))

	return tlvs[1:], nil                     //get rid of the processed tlv		
}