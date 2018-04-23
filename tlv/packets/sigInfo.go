package packets

// SignatureInfo ::= SIGNATURE-INFO-TYPE TLV-LENGTH
//                     SignatureType
//                     KeyLocator?
//                     ... (SignatureType-specific TLVs)
type SignatureInfo struct {
	sigType       uint64
	hasKeyLocator bool
	keyLoc        KeyLocator
}

func NewSignatureInfo(sigType uint64, hasKeyLocator bool, keyLocator KeyLocator) SignatureInfo {
	return SignatureInfo{
		sigType,
		hasKeyLocator,
		keyLocator,
	}
}

func (si SignatureInfo) SetsigType(sigType uint64) {
	si.sigType = sigType
}

func (si SignatureInfo) GetsigType() uint64 {
	return si.sigType
}

func (si SignatureInfo) SetKeyLocator(keyLoc KeyLocator) {
	si.keyLoc = keyLoc
	si.hasKeyLocator = true
}

func (si SignatureInfo) GetKeyLocator() KeyLocator {
	if si.hasKeyLocator == false {
		return KeyLocator{}
	}
	return si.keyLoc
}
