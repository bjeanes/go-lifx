package payloads

import "strings"

type (
	Payload interface {
		Id() uint16
	}

	label [32]byte
)

// FIXME: display actual label (e.g. in lifx-snoop)
func (label label) String() string {
	bytes := [32]byte(label)
	return strings.Trim(string(bytes[:]), "\x00")
}
