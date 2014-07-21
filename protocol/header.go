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
)

func (h header) version() uint16 {
	return 0xfff & uint16(h.Bitfield1) // top 12 bits
}
