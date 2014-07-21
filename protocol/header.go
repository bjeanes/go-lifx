package protocol

type (
	bitfield uint16

	header struct {
		Size uint16

		// 12 bits = protocol version
		// 1 bit   = addressable bool
		// 1 bit   = tagged bool
		// 2 bit   = <reserved>
		Bitfield1 bitfield

		_      uint32 // <reserved>
		Target [8]byte
		Site   [6]byte

		// 1 bit = acknowledge bool
		// 15 bits = <reserved>
		Bitfield2 bitfield

		AtTime uint64
		Type   uint16

		_ uint16 // <reserved>
	}

	Header struct {
		version     uint16
		target      [8]byte
		site        [6]byte
		atTime      uint64
		addressable bool
		tagged      bool
		acknowledge bool
	}
)

func (h *header) version() uint16 { return 0xfff & uint16(h.Bitfield1) }

func (raw *header) ToExpandedHeader() *Header {
	h := new(Header)
	h.atTime = raw.AtTime
	h.target = raw.Target
	h.site = raw.Site
	h.version = raw.version()                        // top 12 bits
	h.addressable = 0x1000&uint16(raw.Bitfield1) > 0 // next bit
	h.tagged = 0x2000&uint16(raw.Bitfield1) > 0      // next bit
	h.acknowledge = 0x1&uint16(raw.Bitfield2) > 0    // top bit
	return h
}

func (h *Header) ToRawHeader() header {
	// TODO
	return header{}
}
