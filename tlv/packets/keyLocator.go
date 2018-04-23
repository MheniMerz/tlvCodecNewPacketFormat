package packets

import (
	"ndn-router/nfd/tlv/name"
)

// KeyLocator ::= KEY-LOCATOR-TYPE TLV-LENGTH (Name | KeyDigest)
// KeyDigest ::= KEY-DIGEST-TYPE TLV-LENGTH BYTE+
type KeyLocator struct {
	Name         name.Name
	HasName      bool
	KeyDigest    []byte
	HasKeyDigest bool
}

func (kl KeyLocator) GetName() name.Name {
	if !kl.HasName {
		return name.NewName()
	}
	return kl.Name
}

func (kl KeyLocator) SetName(n name.Name) {
	kl.HasName = true
	kl.Name = n
}

func (kl KeyLocator) GetKeyDigest() []byte {
	if !kl.HasKeyDigest {
		return nil
	}
	return kl.KeyDigest
}

func (kl KeyLocator) SetKeyDigest(kd []byte) {
	kl.HasKeyDigest = true
	kl.KeyDigest = kd
}

func (kl KeyLocator) IsEmpty() bool {
	return kl.Name == nil && kl.KeyDigest == nil
}
