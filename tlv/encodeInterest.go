package tlv

import (
	"bytes"
	"errors"

	"ndn-router/nfd/tlv/packets"
)

func encodeInterest(i packets.Interest) (Tlv, error) {
	//this returns a slice of Tlvs representing the value of the outer most Tlv
	val, err := encodeSubTlvs(
		i,
		encodeInterestName,
		encodeInterestCanBePrefix,
		encodeInterestMusBeFresh,
		encodeInterestNonce,
		encodeInterestLifeTime,
		encodeInterestHopLimit,
		encodeInterestParameters,
	)

	if err != nil {
		return Tlv{}, err
	}

	var b bytes.Buffer
	TlvsToBytes(val, &b)
	result := Tlv{
		T: INTEREST,
		L: uint64(b.Len()),
		V: b.Next(b.Len()),
	}
	return result, nil
}

func encodeInterestName(packet interface{}, t []Tlv) ([]Tlv, error) {
	name := packet.(packets.Interest).GetName()
	if name.Size() == 0 {
		return nil, errors.New("Encode: -- a packet must have a name --")
	}
	return append(t, encodeName(name)), nil
}

func encodeInterestCanBePrefix(packet interface{}, t []Tlv) ([]Tlv, error){
	cbp := packet.(packets.Interest).GetCanBePrefix()
	if !cbp {
		return t, nil
	}
	canBePrefix := Tlv{
		T: CAN_BE_PREFIX,
		L: uint64(0),
	}
	return append(t, canBePrefix), nil
}

func encodeInterestMusBeFresh(packet interface{}, t []Tlv) ([]Tlv, error){
	mbf := packet.(packets.Interest).GetMustBeFresh()
	if !mbf {
		return t, nil
	}
	mustBeFresh := Tlv{
		T: MUST_BE_FRESH,
		L: uint64(0),
	}
	return append(t, mustBeFresh), nil
}

func encodeInterestNonce(packet interface{}, t []Tlv) ([]Tlv, error) {
	n := packet.(packets.Interest).GetNonce()
	nonce := Tlv{
		T: NONCE,
		L: uint64(len(n)),
		V: n[:],
	}
	return append(t, nonce), nil
}

func encodeInterestLifeTime(packet interface{}, t []Tlv) ([]Tlv, error) {
	lt := packet.(packets.Interest).GetInterestLifetime()
	lifeTime := Tlv{
		T: INTEREST_LIFE_TIME,
	}
	if int64(lt) == -1 {
		return t, nil
		// lifeTime.L = 2                  // 4000 Millisecond is 0x0FA0 ==> 2 bytes
		// lifeTime.V = []byte{0x0F, 0xA0} // 4 Seconds
	} else {
		//need to convert lt to []byte
		b := EncodeNonNegativeInteger(uint64(lt))
		lifeTime.L = uint64(len(b))
		lifeTime.V = b
	}
	return append(t, lifeTime), nil
}

func encodeInterestHopLimit(packet interface{}, t []Tlv) ([]Tlv, error) {
	hl := packet.(packets.Interest).GetHopLimit()
	if !packet.(packets.Interest).HasHopLimit(){
		return t, nil
	}
	hoplimit := Tlv{
		T: HOP_LIMIT,
		L: uint64(1),
		V: EncodeNonNegativeInteger(uint64(hl)),
	}
	return append(t, hoplimit), nil
}

func encodeInterestParameters(packet interface{}, t []Tlv) ([]Tlv, error){
	p := packet.(packets.Interest).GetParameters()
	
	if p ==nil{
		return t, nil
	}	

	parameters := Tlv{
		T: PARAMETERS,
		L: uint64(len(p)),
		V: p,
	}
	return append(t, parameters), nil
}