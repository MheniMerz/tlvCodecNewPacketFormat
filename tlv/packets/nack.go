package packets

type NackReason uint64

const (
	nackReasonCongestion	NackReason	= 50
	nackReasonDuplicate					= 100
	nackReasonNoRoute					= 150
)

type Nack struct{
	packet *Interest
	reason NackReason
	buffer      []byte	
}

//implementing the NdnPacket interface
func (n Nack) PacketType() uint64 {
	return 0xdd
}

func (n *Nack) SetPacket(i *Interest){
	n.packet = i
}

func (n Nack) GetPacket() *Interest{
	return n.packet
}

func (n *Nack) SetReason(r NackReason){
	n.reason = r
}

func (n Nack)GetReason() string{
	switch n.reason {
	case nackReasonCongestion :
		return "Congestion"
	case nackReasonDuplicate :
		return "Duplicate"
	case nackReasonNoRoute :
		return "No Route"
	}
	return "Unknown"
}

func (n *Nack) Setbuffer(b []byte) {
	n.buffer = b
}

func (n Nack) GetBuffer() []byte {
	return n.buffer
}