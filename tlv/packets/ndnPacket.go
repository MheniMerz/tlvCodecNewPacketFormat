package packets

// Ndn packet is the generic packet -- could be interest, data ot nack --
//any type that implements the PacketType() methode will be an NdnPacket
//this is used for the return value of the decode function, because it could return
//either an interest, data, or nack ==> which means we need a more generic type that
// can be any of the three at once
type NdnPacket interface {
	PacketType() uint64
}
