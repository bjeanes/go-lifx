package payloads

import "strings"

type Payload interface {
	Id() uint16
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

type lightHsbk struct {
	Hue        uint16 // 0-65535 scaled to 0-360°
	Saturation uint16 // 0-65535 scaled to 0-100%
	Brightness uint16 // 0-65535 scaled to 0-100%
	Kelvin     uint16 // absolute 2400-10000
}

type deviceSetSite struct {
	Site [6]byte
}

type deviceGetPanGateway struct{}

type deviceStatePanGateway struct {
	Service uint8  // 1 for UDP; 2 for TCP
	Port    uint32 // 0 for disabled, else it's the service port
}

type deviceTime struct{}

type deviceSetTime struct {
	Time uint64 // nanos since epoch
}

type deviceStateTime struct {
	Time uint64 // nanos since epoch
}

type deviceGetResetSwitch struct{}

type deviceStateResetSwitch struct {
	Position uint8
}

type deviceGetMeshInfo struct{}

type deviceStateMeshInfo struct {
	Signal         float32 // milliwatts
	Tx, Rx         uint32  // bytes
	McuTemperature int16   // deci-celsius
}

type deviceGetMeshFirmware struct{}

type deviceStateMeshFirmware struct {
	Build, Install uint64
	Version        uint32
}

type deviceGetWifiInfo struct{}

type deviceStateWifiInfo struct {
	Signal         float32 // milliwatts
	Tx, Rx         uint32  // bytes
	McuTemperature int16   // deci-celsius
}

type deviceGetWifiFirmware struct{}

type deviceStateWifiFirmware struct {
	Build, Install uint64
	Version        uint32
}

type deviceGetPower struct{}

type deviceSetPower struct {
	Level uint16 // 0 = off; 1+ = on
}

type deviceStatePower struct {
	Level uint16 // 0 = off; 1+ = on
}

type deviceGetLabel struct{}

type deviceSetLabel struct {
	Label [32]byte // string
}

type deviceStateLabel struct {
	Label [32]byte // string
}

type deviceGetTags struct{}

type deviceSetTags struct {
	Tags uint64
}

type deviceStateTags struct {
	Tags uint64
}

type deviceGetTagLabels struct {
	Tags uint64
}

type deviceSetTagLabels struct {
	Tags  uint64
	Label [32]byte
}

type deviceStateTagLabels struct {
	Tags  uint64
	Label [32]byte
}

type deviceGetVersion struct{}

type deviceStateVersion struct {
	Vendor, Product, Version uint32
}

type deviceGetInfo struct{}

type deviceStateInfo struct {
	Time, Uptime, Downtime uint64 // ns
}

type deviceGetMcuRailVoltage struct{}

type deviceStateMcuRailVoltage struct {
	Voltage uint32
}

type deviceReboot struct{}

type lightGet struct{}

type lightSet struct {
	Stream   uint8
	Color    lightHsbk
	Duration uint32 // ms
}

type lightSetWaveform struct {
	Stream    uint8
	Transient uint8 // 0 false; 1+ true
	Color     lightHsbk
	Period    uint32 // ms per cycle
	DutyCycle int16
	Waveform  uint8
}

type lightSetDimAbsolute struct {
	Brightness int32  // 0 for no change
	Duration   uint32 // ms
}

type lightSetDimRelative struct {
	Brightness int32  // 0 for no change
	Duration   uint32 // ms
}

type lightSetRgbw struct {
	Color struct {
		Red, Green, Blue, White uint16
	}
}

type lightState struct {
	Color lightHsbk
	Dim   int16
	Power uint16
	Label label
	Tags  uint64
}

type lightGetRailVoltage struct{}

type lightStateRailVoltage struct {
	Voltage uint32
}

type lightGetTemperature struct{}

type lightStateTemperature struct {
	Temperature int16 // deci-celsius
}

type wanConnectPlain struct {
	User, Pass [32]byte
}

type wanConnectKey struct {
	AuthKey [32]byte
}

type wanStateConnect struct {
	AuthKey [32]byte
}

type wanSub struct {
	Target [8]byte
	Site   [6]byte
	Device uint8 // 0 device; 1 tag
}

type wanUnsub struct {
	Target [8]byte
	Site   [6]byte
	Device uint8 // 0 device; 1 tag
}

type wanStateSub struct {
	Target [8]byte
	Site   [6]byte
	Device uint8 // 0 device; 1 tag
}

type wifiGet struct {
	Iface ifaceType
}

type wifiSet struct {
	Iface  ifaceType
	Active uint8 // 0 false; 1 true
}

type wifiState struct {
	Iface  ifaceType
	Status wifiStatus
	Ipv4   uint32
	Ipv6   [16]byte
}

type wifiGetAccessPoint struct{}

type wifiSetAccessPoint struct {
	Iface    ifaceType
	Ssid     [32]byte
	Password [64]byte
	Security apSecurity
}

type wifiStateAccessPoint struct {
	Iface    ifaceType
	Ssid     [32]byte
	Password [64]byte
	Security apSecurity
	Channel  uint16
}

type sensorGetAmbientLight struct{}

type sensorStateAmbientLight struct {
	Lux float32
}

type sensorGetDimmerVoltage struct{}

type sensorStateDimmerVoltage struct {
	Voltage uint32
}
