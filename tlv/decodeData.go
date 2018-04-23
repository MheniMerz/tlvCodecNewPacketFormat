package tlv

import (
	"errors"
	"ndn-router/nfd/tlv/name"
	"ndn-router/nfd/tlv/packets"
	"time"
)

func decodeData(t Tlv) (packets.Data, error) {
	tlvs, _ := ParseTlvsFromBytes(t.V)
	resultData := packets.Data{}
	err := decodeTlvs(
		&resultData, tlvs,
		decodeDataName,
		decodeDataMetaInfo,
		decodeDataContent,
		decodeDataSignature,
	)
	if err != nil {
		return packets.Data{}, err
	}
	return resultData, nil
}

func decodeDataName(packet interface{}, tlvs []Tlv) ([]Tlv, error) {
	if len(tlvs) < 1 {
		return nil, errors.New("DecodeDataName : --- no tlvs to read ---")
	}
	//decodeName is common to both interest and data
	name, err := decodeName(tlvs[0]) //the name tlv is the first one tlvs[0]
	if err != nil {
		return tlvs, err
	}
	packet.(*packets.Data).SetName(name) //the result
	return tlvs[1:], nil                 //get rid of the processed tlv
}

func decodeDataMetaInfo(packet interface{}, tlvs []Tlv) ([]Tlv, error) {
	if len(tlvs) < 1 {
		return nil, errors.New("DecodeDataMetaInfo : --- no tlvs to read ---")
	}
	t := tlvs[0]
	if t.T != META_INFO {
		return tlvs, errors.New("--- DecodeDataMetaInfo --- : unexpected type")
	}
	metaFields, _ := ParseTlvsFromBytes(t.V) // from []bytes to []Tlv
	for _, field := range metaFields {
		err := decodeMetaField(field, packet.(*packets.Data))
		if err != nil {
			return nil, err
		}
	}
	return tlvs[1:], nil
}

func decodeMetaField(field Tlv, packet *packets.Data) error {
	switch field.T {
	case CONTENT_TYPE:
		x := DecodeNonNegativeInteger(field.V)
		packet.MetaInfo.SetContentType(packets.ContentType(x))
	case FRESHNESS_PERIOD:
		x := DecodeNonNegativeInteger(field.V)
		packet.MetaInfo.SetFreshnessPeriod(time.Duration(x))

	case FINAL_BLOCK_ID:
		t, _, _, _ := TlvFromBytes(field.V)
		x, err := decodeNameComponent(t)
		if err != nil {
			return err
		}
		packet.MetaInfo.SetFinalBlockID(x)
	}
	return nil
}

func decodeDataContent(packet interface{}, tlvs []Tlv) ([]Tlv, error) {
	if len(tlvs) < 1 {
		return nil, errors.New("DecodeDataContent : --- no tlvs to read ---")
	}
	t := tlvs[0]
	if t.T != CONTENT {
		return tlvs, errors.New("--- DecodeDataContent --- : unexpected type")
	}
	//the content is []bytes, it is value of the current tlv
	packet.(*packets.Data).SetContent(t.V)
	return tlvs[1:], nil
}

func decodeDataSignature(packet interface{}, tlvs []Tlv) ([]Tlv, error) {
	if len(tlvs) < 2 {
		return nil, errors.New("DecodeDataContent : --- no tlvs to read ---")
	}
	info := tlvs[0]
	val := tlvs[1]
	if info.T != SIGNATURE_INFO {
		return tlvs, errors.New("--- DecodeDataSignature ..SigInfo.. --- : unexpected type")
	}
	if val.T != SIGNATURE_VALUE {
		return tlvs, errors.New("--- DecodeDataSignature ..SigVal.. --- : unexpected type")
	}
	valBytes, _ := decodeSignatureValue(val)
	sigInfo, _ := decodeSignatureInfo(info)
	sig := packets.NewSignature(sigInfo, valBytes)
	packet.(*packets.Data).SetSignature(sig)
	return tlvs[2:], nil
}

func decodeSignatureValue(t Tlv) ([]byte, error) {
	if t.T != SIGNATURE_VALUE {
		return nil, errors.New("DecodeSignatureValue : --- unexpected type ---")
	}
	return t.V, nil
}

func decodeSignatureInfo(t Tlv) (packets.SignatureInfo, error) {
	tlvs, _ := ParseTlvsFromBytes(t.V)
	sigTypeTlv := tlvs[0]
	keyLocatorTlv := Tlv{}
	//checking if the keyLocator is there
	if len(tlvs) > 1 {
		keyLocatorTlv = tlvs[1]
	}
	sigType, _ := decodeSignatureType(sigTypeTlv)
	keyLocator, hasKeyLoc, _ := decodeKeyLocator(keyLocatorTlv)
	result := packets.NewSignatureInfo(sigType, hasKeyLoc, keyLocator)
	return result, nil
}

func decodeSignatureType(t Tlv) (uint64, error) {
	if t.T != SIGNATURE_TYPE {
		return 0, errors.New("DecodeSignatureValue : --- unexpected type ---")
	}
	sigType := DecodeNonNegativeInteger(t.V)
	return sigType, nil
}

func decodeKeyLocator(t Tlv) (packets.KeyLocator, bool, error) {
	if t.T != KEY_LOCATOR {
		return packets.KeyLocator{}, false, errors.New("DecodeSignatureValue : --- unexpected type ---")
	}
	keyLocValueTlv, _, _, _ := TlvFromBytes(t.V)
	result := packets.KeyLocator{}
	switch keyLocValueTlv.T {
	case NAME:
		nameRef, _ := decodeName(keyLocValueTlv)
		//fmt.Printf("+++++ %v +++++", nameRef)
		//fmt.Printf("+++++ %v +++++", keyLocValueTlv)
		result = packets.KeyLocator{
			nameRef,
			true,
			nil,
			false,
		}
	case KEY_DIGEST:
		keyDigest, _ := decodeKeyDigest(keyLocValueTlv)
		result = packets.KeyLocator{
			name.NewName(),
			false,
			keyDigest,
			true,
		}

	}
	return result, true, nil
}

func decodeKeyDigest(t Tlv) ([]byte, error) {
	if t.T != KEY_DIGEST {
		return nil, errors.New("DecodeSignatureValue : --- unexpected type ---")
	}
	return t.V, nil
}