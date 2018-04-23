package tlv

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
)

func DecodeNonNegativeInteger(b []byte) uint64 {
	switch len(b) {
	case 1:
		return uint64(b[0])
	case 2:
		return uint64(binary.BigEndian.Uint16(b))
	case 4:
		return uint64(binary.BigEndian.Uint32(b))
	case 8:
		return binary.BigEndian.Uint64(b)
	default:
		return 0
	}
}

func EncodeNonNegativeInteger(n uint64) []byte {
	b := [8]byte{}
	if n <= 0xFF {
		b[0] = uint8(n)
		return b[:1]
	} else if n <= 0xFFFF {
		binary.BigEndian.PutUint16(b[:], uint16(n))
		return b[:2]
	} else if n <= 0xFFFFFFFF {
		binary.BigEndian.PutUint32(b[:], uint32(n))
		return b[:4]
	} else {
		binary.BigEndian.PutUint64(b[:], uint64(n))
		return b[:8]
	}
}

//variable length decoding
func varDecoding(packet []byte) (val uint64, size int) {
	concat := "" //used to concatenate the bytes that will form the type's value
	if len(packet) > 0{
		switch packet[0] {
		case 0xFD: // 2 bytes
			size = 3
		case 0xFE: // 4 bytes
			size = 5

		case 0xFF: // 8 bytes
			size = 9

		default: //1 byte
			size = 1
			concat += fmt.Sprintf("%02x", packet[0])
		}
		if size != 1 {
			for i := 1; i < size; i++ {
				concat += fmt.Sprintf("%02x", packet[i])
			}
		}
	}
	val, _ = strconv.ParseUint(concat, 16, 64) // conversion to decimal
	return
}

// takes the type, or length as int and transforms it to []byte
// considering the variable length encoding
func varEncoding(num uint64, w io.Writer) error {

	if num < 253 {
		return binary.Write(w, binary.BigEndian, uint8(num))
	}

	var first byte
	var val interface{}
	if num <= 0xFFFF {
		first = 253
		val = uint16(num)
	} else if num <= 0xFFFFFFFF {
		first = 254
		val = uint32(num)
	} else {
		first = 255
		val = num
	}
	err := binary.Write(w, binary.BigEndian, first)
	if err != nil {
		return err
	}

	return binary.Write(w, binary.BigEndian, val)
}

//returns bytes from a buffer
func ReadBytesFromBuffer(buf bytes.Buffer) []byte {
	return buf.Bytes()
}

//creates an io.Reader from a []byte
func CreateIoReader(b []byte) io.Reader {
	return bytes.NewReader(b)
}

//read bytes from io.Reader
func BytesFromIoReader(r io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	//r.Seek(0, 0)
	return buf.Bytes()
}

func SliceEqual(a, b []byte) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
