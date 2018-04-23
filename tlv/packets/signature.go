package packets

// Signature ::= SignatureInfo
//               SignatureValue
type Signature struct {
	sigInfo SignatureInfo
	val     []byte
}

func NewSignature(sigInfo SignatureInfo, val []byte) Signature {
	return Signature{
		sigInfo,
		val,
	}
}

func (s Signature) SetsigInfo(si SignatureInfo) {
	s.sigInfo = si
}

func (s Signature) GetsigInfo() SignatureInfo {
	return s.sigInfo
}

func (s Signature) SetsigVal(val []byte) {
	s.val = val
}

func (s Signature) GetsigVal() []byte {
	return s.val
}
