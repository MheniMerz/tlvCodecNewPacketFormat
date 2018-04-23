package tlv

import (
	"bytes"
	"errors"

	"ndn-router/nfd/tlv/packets"
)

func encodeData(d packets.Data) (Tlv, error) {
	//this returns a slice of Tlvs representing the value of the outer most Tlv
	val, err := encodeSubTlvs(
		d,
		encodeDataName,
		encodeDataMetaInfo,
		encodeDataContent,
		encodeDataSignature,
	)
	if err != nil {
		return Tlv{}, err
	}

	var b bytes.Buffer
	TlvsToBytes(val, &b)
	result := Tlv{
		T: DATA,
		L: uint64(b.Len()),
		V: b.Next(b.Len()),
	}
	return result, nil
}


func encodeDataName(packet interface{}, t []Tlv) ([]Tlv, error) {
	name := packet.(packets.Data).GetName()
	if name.Size() == 0 {
		return nil, errors.New("Encode: -- a packet must have a name --")
	}
	return append(t, encodeName(name)), nil
}


func encodeDataMetaInfo(packet interface{}, t []Tlv) ([]Tlv, error) {
	mi := packet.(packets.Data).GetMetaInfo()
	val := []Tlv{}
	if ct := mi.GetContentType(); ct != packets.Unknown {
		ctToByte := EncodeNonNegativeInteger(uint64(ct))
		x := Tlv{T: CONTENT_TYPE, L: uint64(len(ctToByte)), V: ctToByte}
		val = append(val, x)
	}
	if fp := mi.GetFreshnessPeriod(); fp != -1 {
		fpToByte := EncodeNonNegativeInteger(uint64(fp))
		x := Tlv{T: FRESHNESS_PERIOD, L: uint64(len(fpToByte)), V: fpToByte}
		val = append(val, x)
	}
	if id := mi.GetFinalBlockID(); len(id.ToString()) > 0 {
		idToByte := id.ComponentToBytes()
		x := Tlv{T: CONTENT_TYPE, L: uint64(len(idToByte)), V: idToByte}
		val = append(val, x)
	}

	var b bytes.Buffer
	err := TlvsToBytes(val, &b)
	if err != nil {
		return t, err
	}
	metaInfo := Tlv{
		T: META_INFO,
		L: uint64(b.Len()),
		V: b.Next(b.Len()),
	}
	return append(t, metaInfo), nil
}

func encodeDataContent(packet interface{}, t []Tlv) ([]Tlv, error) {
	ct := packet.(packets.Data).GetContent()
	x := Tlv{
		T: CONTENT,
		L: uint64(len(ct)),
		V: ct,
	}
	return append(t, x), nil
}

func encodeDataSignature(packet interface{}, t []Tlv) ([]Tlv, error) {
	sig := packet.(packets.Data).GetSignature()
	sigInfo := sig.GetsigInfo()
	sigVal := sig.GetsigVal()
	sigInfoTlv := encodeSignatureInfo(sigInfo)
	sigValTlv := encodeSignatureValue(sigVal)
	t = append(t, sigInfoTlv)
	return append(t, sigValTlv), nil
}

func encodeSignatureInfo(sigInfo packets.SignatureInfo) Tlv {
	sigType := sigInfo.GetsigType()
	keyLoc := sigInfo.GetKeyLocator()
	tmp := make([]Tlv, 2)
	tmp[0] = encodeSignatureType(sigType)
	tmp[1] = encodeKeyLocator(keyLoc)
	var b bytes.Buffer
	TlvsToBytes(tmp, &b)
	result := Tlv{
		T: SIGNATURE_INFO,
		L: uint64(b.Len()),
		V: b.Next(b.Len()),
	}
	return result
}

func encodeSignatureType(sigType uint64) Tlv {
	x := EncodeNonNegativeInteger(sigType)
	return Tlv{
		T: SIGNATURE_TYPE,
		L: uint64(len(x)),
		V: x,
	}
}

func encodeKeyLocator(keyLoc packets.KeyLocator) Tlv {
	var b bytes.Buffer
	if keyLoc.HasName {
		x := encodeName(keyLoc.Name)
		err := TlvToBytes(x, &b)
		if err != nil {
			return Tlv{}
		}
		return Tlv{
			T: KEY_LOCATOR,
			L: uint64(b.Len()),
			V: b.Next(b.Len()),
		}
	} else {
		if keyLoc.HasKeyDigest {
			x := encodeKeyDigest(keyLoc.KeyDigest)
			err := TlvToBytes(x, &b)
			if err != nil {
				return Tlv{}
			}
			return Tlv{
				T: KEY_LOCATOR,
				L: uint64(b.Len()),
				V: b.Next(b.Len()),
			}
		}
	}
	return Tlv{}
}

func encodeKeyDigest(kd []byte) Tlv {
	return Tlv{
		T: KEY_DIGEST,
		L: uint64(len(kd)),
		V: kd,
	}
}

func encodeSignatureValue(sigVal []byte) Tlv {
	return Tlv{
		T: SIGNATURE_VALUE,
		L: uint64(len(sigVal)),
		V: sigVal,
	}
}