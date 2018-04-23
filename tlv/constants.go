package tlv

const(
	//packet types
	INTEREST					= 0x05
	DATA						= 0x06
	NACK						= 0xdd					
	//common fields
	NAME						= 0x07
	NAME_COMPONENT				= 0x08
	IMPLICIT_DIGEST				= 0x01
	//interest packet
	CAN_BE_PREFIX				= 0x21
	MUST_BE_FRESH				= 0x12
	FORWARDING_HINT				= 0x1e
	NONCE						= 0x0a
	INTEREST_LIFE_TIME			= 0x0c
	HOP_LIMIT					= 0x22
	PARAMETERS					= 0x23
	//Data packet
	META_INFO					= 0x14
	CONTENT						= 0x15
	SIGNATURE_INFO				= 0x16
	SIGNATURE_VALUE				= 0x17
	//Data META_INFO
	CONTENT_TYPE				= 0x18
	FRESHNESS_PERIOD			= 0x19
	FINAL_BLOCK_ID				= 0x1a
	//Data signature
	SIGNATURE_TYPE				= 0x1b
	KEY_LOCATOR					= 0x1c
	KEY_DIGEST					= 0x1d
	//Link Object
	DELEGEATION					= 0x1f
	PREFERENCE					= 0x1e
)
