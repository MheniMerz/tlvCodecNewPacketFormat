package packets

import (
	"time"

	"ndn-router/nfd/tlv/name"
)

type ContentType int

const (
	Unknown  ContentType = -1
	BLOB     ContentType = 0
	LINK_OBJ ContentType = 1
	PUB_KEY  ContentType = 2
	APP_NACK ContentType = 3
)

// MetaInfo ::= META-INFO-TYPE TLV-LENGTH
//                ContentType?
//                FreshnessPeriod?
//                FinalBlockId?
type MetaInfo struct {
	contentType    ContentType
	hasContentType bool

	freshnessPeriod    time.Duration
	hasFreshnessPeriod bool

	finalBlockID    name.Component
	hasFinalBlockID bool
}

//getters and setters
func (m MetaInfo) GetContentType() ContentType {
	if !m.hasContentType {
		return Unknown
	}
	return m.contentType
}

func (m *MetaInfo) SetContentType(c ContentType) {
	m.hasContentType = true
	m.contentType = c
}

func (m MetaInfo) GetFreshnessPeriod() time.Duration {
	if !m.hasFreshnessPeriod {
		return 0
	}
	return m.freshnessPeriod
}

func (m *MetaInfo) SetFreshnessPeriod(f time.Duration) {
	m.hasFreshnessPeriod = true
	m.freshnessPeriod = f
}

func (m MetaInfo) GetFinalBlockID() name.Component {
	if !m.hasFinalBlockID {
		id := name.Component("")
		return id
	}
	return m.finalBlockID
}

func (m *MetaInfo) SetFinalBlockID(id name.Component) {
	m.hasFinalBlockID = true
	m.finalBlockID = id.Copy()
}
