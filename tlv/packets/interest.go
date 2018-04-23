package packets

import (
	"math/rand"
	//"ndn-router/ndn/tlv"
	"encoding/binary"
	"ndn-router/nfd/tlv/name"
	"time"
)

// Interest ::= INTEREST-TYPE TLV-LENGTH
// 		Name
// 		CanBePrefix?
// 		MustBeFresh?
// 		ForwardingHint?
// 		Nonce?
// 		InterestLifetime?
// 		HopLimit?
// 		Parameters?

type Interest struct {
	name        name.Name
	canBePrefix	bool
	mustBeFresh bool
	//ForwardingHint 
	nonce       [4]byte
	lifetime    time.Duration
	hasHopLimit	bool
	hopLimit	uint8
	parameters	[]byte
	buffer      []byte
}


func NewInterest(name name.Name) *Interest {
	i := Interest{
		name: name,
		lifetime: 4 * time.Second,
	}
	i.GenerateNonce()
	return &i
}

//implementing the NdnPacket interface
func (i Interest) PacketType() uint64 {
	return 5
}

// getters and setters
func (i Interest) GetName() name.Name {
	if i.name == nil {
		i.name = name.Name{}
	}
	return i.name
}

func (i *Interest) SetName(n name.Name) {
	i.name = n
}

func (i Interest) GetCanBePrefix() bool{
	return i.canBePrefix
}

func (i *Interest) SetCanBePrefix(cbp bool){
	i.canBePrefix = cbp
}

func (i Interest) GetMustBeFresh() bool{
	return i.mustBeFresh
}

func (i *Interest) SetMustBeFresh(mbf bool){
	i.mustBeFresh = mbf
}

func (i Interest) GetNonce() [4]byte {
	return i.nonce
}

func (i *Interest) SetNonce(n [4]byte) {
	i.nonce = n
}

func (i *Interest) GenerateNonce() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	randNonce := r.Int31()
	i.nonce = nonceToBytes(randNonce)
}

func nonceToBytes(n int32) [4]byte {
	b := [4]byte{}
	binary.BigEndian.PutUint32(b[:], uint32(n))
	return b
}

func (i Interest) GetInterestLifetime() time.Duration {
	return i.lifetime
}

func (i *Interest) SetInterestLifetime(x time.Duration) {
	i.lifetime = x
}

func (i Interest) HasHopLimit() bool{
	return i.hasHopLimit
}

func (i Interest) GetHopLimit() uint8{
		return i.hopLimit
}

func (i *Interest) SetHopLimit(hp uint8){
	i.hopLimit = hp
	i.hasHopLimit = true
}

func (i Interest) GetParameters() []byte {
	return i.parameters
}

func (i *Interest) SetParameters(p []byte) {
	i.parameters = p
}

func (i Interest) GetBuffer() []byte {
	return i.buffer
}

func (i *Interest) Setbuffer(b []byte) {
	i.buffer = b
}
