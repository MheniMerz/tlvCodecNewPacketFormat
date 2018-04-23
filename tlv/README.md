TLV encoding/decoding
=====================

## Decoding part
- takes []byte as input, and gives back user defined types Interest or data
- performance could be enhanced by using buffers for the tlvs, instead of passing values back and forth between decode functions.

## Encoding part
- takes as input the NdnPacket and the byte buffer to write on
- the output is written on the buffer, and predefined functions are used to retrieve that output

### To do
- need to complete the packet fields
	- signature for the data packet
	- selectors, linkObject for interest packet
- no concurrency yet ** next model **
	- try concurrency on multiple packets
	- try concurrency on the same packet (previous results were not promissing)
		- in this case (same packet) i will need to use locks on the packet because of simultanous access
