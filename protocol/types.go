package protocol

import "strings"

type Site [6]byte

func (site Site) String() string {
	bytes := [6]byte(site)
	return strings.Trim(string(bytes[:]), "\x00")
}

type label [32]byte

func (label label) String() string {
	bytes := [32]byte(label)
	return strings.Trim(string(bytes[:]), "\x00")
}

type (
	apSecurity uint8
	ifaceType  uint8
	wifiStatus uint8
)

type Hsbk struct {
	Hue        uint16 // 0-65535 scaled to 0-360°
	Saturation uint16 // 0-65535 scaled to 0-100%
	Brightness uint16 // 0-65535 scaled to 0-100%
	Kelvin     uint16 // absolute 2400-10000
}
