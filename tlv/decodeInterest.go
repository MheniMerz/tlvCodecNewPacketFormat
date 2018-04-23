package tlv

import (
	"errors"
	"ndn-router/nfd/tlv/packets"
	"time"
	"log"
)

//take the value od the interest TLV and call the different decode functiond for each sub tlv
//and return ant interest
func decodeInterest(t Tlv) (packets.Interest, error) {
	tlvs, _ := ParseTlvsFromBytes(t.V)
	resultInterest := packets.Interest{}
	err := decodeTlvs(
		&resultInterest, tlvs,
		decodeInterestName,
		decodeInterestCanBePrefix,
		decodeInterestMustBeFresh,
		//decodeInterestForwardingHint,
		decodeInterestNonce,
		decodeInterestLifeTime,
		decodeInterestHopLimit,
		decodeInterestParameters,
	)
	if err != nil {
		return packets.Interest{}, err
	}
	return resultInterest, nil
}

//takes the name tlv and calls the decode name to get the name back and sets the result's name
func decodeInterestName(packet interface{}, tlvs []Tlv) ([]Tlv, error) {
	if len(tlvs) < 1 {
		return nil, errors.New("DecodeInterestName : --- no tlvs to read ---")
	}
	//decodeName is common to both interest and data
	name, err := decodeName(tlvs[0]) //the name tlv is the first one tlvs[0]
	if err != nil {
		return tlvs, err
	}
	packet.(*packets.Interest).SetName(name) //the result
	return tlvs[1:], nil                     //get rid of the processed tlv
}

func decodeInterestCanBePrefix(packet interface{}, tlvs []Tlv) ([]Tlv, error){
	if len(tlvs) < 1 {
		packet.(*packets.Interest).SetCanBePrefix(false)			
		log.Println("decodeInterestCanBePrefix : --- no tlvs to read ---")
		return tlvs, nil
	}
	t := tlvs[0]
	if t.T != CAN_BE_PREFIX{
		packet.(*packets.Interest).SetCanBePrefix(false)
		log.Println("decodeInterestCanBePrefix : --- unexpected type ---")		
		return tlvs, nil
	}
	packet.(*packets.Interest).SetCanBePrefix(true)		
	return tlvs[1:], nil                     //get rid of the processed tlv
}

func decodeInterestMustBeFresh(packet interface{}, tlvs []Tlv) ([]Tlv, error){
	if len(tlvs) < 1 {
		packet.(*packets.Interest).SetMustBeFresh(false)
		log.Println("decodeInterestMustBeFresh : --- no tlvs to read ---")
		return tlvs, nil
	}
	t := tlvs[0]
	if t.T != MUST_BE_FRESH{
		packet.(*packets.Interest).SetMustBeFresh(false)
		log.Println("decodeInterestMustBeFresh : --- unexpected type ---")		
		return tlvs, nil		
	}
	packet.(*packets.Interest).SetMustBeFresh(true)
	return tlvs[1:], nil                     //get rid of the processed tlv	
}

func decodeInterestNonce(packet interface{}, tlvs []Tlv) ([]Tlv, error) {
	if len(tlvs) < 1 {
		log.Println("decodeInterestNonce : --- no tlvs to read ---")
		return tlvs, nil
	}
	t := tlvs[0]
	if t.T != NONCE {
		log.Println("decodeInterestNonce : --- unexpected type ---")		
		return tlvs, nil
	}
	nonce := [4]byte{}
	for i := 0; i < len(nonce); i++ {
		nonce[i] = t.V[i]
	}
	packet.(*packets.Interest).SetNonce(nonce)
	return tlvs[1:], nil
}

func decodeInterestLifeTime(packet interface{}, tlvs []Tlv) ([]Tlv, error) {
	if len(tlvs) < 1 {
		return nil, errors.New("DecodeInterestLifeTime : --- no tlvs to read ---")
	}
	t := tlvs[0]
	if t.T != INTEREST_LIFE_TIME {
		log.Println("decodeInterestLifeTime : --- unexpected type ---")		
		return tlvs, errors.New("--- DecodeInterestLifeTime --- : unexpected type")
	}
	lifeTime := DecodeNonNegativeInteger(t.V)
	packet.(*packets.Interest).SetInterestLifetime(time.Duration(lifeTime))
	return tlvs[1:], nil
}

func decodeInterestHopLimit(packet interface{}, tlvs []Tlv) ([]Tlv, error) {
	if len(tlvs) < 1 {
		log.Println("deecodeInterestHopLimit : --- no tlvs to read ---")
		return tlvs, nil
	}
	t := tlvs[0]
	if t.T != HOP_LIMIT {
		log.Println("deecodeInterestHopLimit : --- unexpected type ---")		
		return tlvs, nil
	}
	hl := uint8(DecodeNonNegativeInteger(t.V))
	packet.(*packets.Interest).SetHopLimit(hl)	
	return tlvs[1:], nil
}

func decodeInterestParameters(packet interface{}, tlvs []Tlv) ([]Tlv, error){
	if len(tlvs) < 1 {
		log.Println("decodeInterestParameters : --- no tlvs to read ---")
		return tlvs, nil
	}
	t := tlvs[0]
	if t.T != PARAMETERS {
		log.Println("decodeInterestParameters : --- unexpected type ---")		
		return tlvs, nil
	}
	//the content is []bytes, it is value of the current tlv
	packet.(*packets.Interest).SetParameters(t.V)	
	return tlvs[1:], nil
}