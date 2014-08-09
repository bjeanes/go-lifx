package protocol

import (
	"fmt"
	"math"
	"strings"
)

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

type Degrees uint16

func (d Degrees) String() string {
	return fmt.Sprintf("%.2f°", 360*float32(d)/math.MaxUint16)
}

type Percent uint16

func (p Percent) String() string {
	return fmt.Sprintf("%.2f%%", 100*float32(p)/math.MaxUint16)
}

type Kelvin uint16

func (k Kelvin) String() string {
	return fmt.Sprintf("%d°", uint16(k))
}

type Hsbk struct {
	Hue        Degrees // 0-65535 scaled to 0-360°
	Saturation Percent // 0-65535 scaled to 0-100%
	Brightness Percent // 0-65535 scaled to 0-100%
	Kelvin     Kelvin  // absolute 2400-10000
}
