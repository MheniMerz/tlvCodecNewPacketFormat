package packets

import (
	"ndn-router/nfd/tlv/name"
)

// Data ::= DATA-TLV TLV-LENGTH
//            Name
//            MetaInfo
//            Content
//            Signature

type Data struct {
	name      name.Name
	MetaInfo  MetaInfo
	content   []byte
	signature Signature
	buffer    []byte
}

//implementing the NdnPacket interface
func (d Data) PacketType() uint64 {
	return 6
}

func (d Data) GetName() name.Name {
	if d.name == nil {
		d.name = name.Name{}
	}
	return d.name
}

func (d *Data) SetName(x name.Name) {
	d.name = x
}

func (d *Data) Setbuffer(b []byte) {
	d.buffer = b
}

func (d Data) GetBuffer() []byte {
	return d.buffer
}

func (d Data) GetContent() []byte {
	if d.content == nil {
		d.content = []byte{}
	}
	return d.content
}

func (d *Data) SetContent(x []byte) {
	d.content = x
}

func (d Data) GetMetaInfo() MetaInfo {
	return d.MetaInfo
}

func (d Data) SetMetaInfo(m MetaInfo) {
	d.MetaInfo = m
}

func (d Data) GetSignature() Signature {
	return d.signature
}

func (d *Data) SetSignature(s Signature) {
	d.signature = s
}
