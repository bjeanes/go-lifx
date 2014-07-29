package payloads

import "strings"

type Payload interface {
	Id() uint16
}

type site [6]byte

func (site site) String() string {
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

type LightHsbk struct {
	Hue        uint16 // 0-65535 scaled to 0-360°
	Saturation uint16 // 0-65535 scaled to 0-100%
	Brightness uint16 // 0-65535 scaled to 0-100%
	Kelvin     uint16 // absolute 2400-10000
}

type DeviceSetSite struct {
	Site site
}

type DeviceGetPanGateway struct{}

type DeviceStatePanGateway struct {
	Service uint8  // 1 for UDP; 2 for TCP
	Port    uint32 // 0 for disabled, else it's the service port
}

type DeviceTime struct{}

type DeviceSetTime struct {
	Time uint64 // nanos since epoch
}

type DeviceStateTime struct {
	Time uint64 // nanos since epoch
}

type DeviceGetResetSwitch struct{}

type DeviceStateResetSwitch struct {
	Position uint8
}

type DeviceGetMeshInfo struct{}

type DeviceStateMeshInfo struct {
	Signal         float32 // milliwatts
	Tx, Rx         uint32  // bytes
	McuTemperature int16   // deci-celsius
}

type DeviceGetMeshFirmware struct{}

type DeviceStateMeshFirmware struct {
	Build, Install uint64
	Version        uint32
}

type DeviceGetWifiInfo struct{}

type DeviceStateWifiInfo struct {
	Signal         float32 // milliwatts
	Tx, Rx         uint32  // bytes
	McuTemperature int16   // deci-celsius
}

type DeviceGetWifiFirmware struct{}

type DeviceStateWifiFirmware struct {
	Build, Install uint64
	Version        uint32
}

type DeviceGetPower struct{}

type DeviceSetPower struct {
	Level uint16 // 0 = off; 1+ = on
}

type DeviceStatePower struct {
	Level uint16 // 0 = off; 1+ = on
}

type DeviceGetLabel struct{}

type DeviceSetLabel struct {
	Label label // string
}

type DeviceStateLabel struct {
	Label label // string
}

type DeviceGetTags struct{}

type DeviceSetTags struct {
	Tags uint64
}

type DeviceStateTags struct {
	Tags uint64
}

type DeviceGetTagLabels struct {
	Tags uint64
}

type DeviceSetTagLabels struct {
	Tags  uint64
	Label label
}

type DeviceStateTagLabels struct {
	Tags  uint64
	Label label
}

type DeviceGetVersion struct{}

type DeviceStateVersion struct {
	Vendor, Product, Version uint32
}

type DeviceGetInfo struct{}

type DeviceStateInfo struct {
	Time, Uptime, Downtime uint64 // ns
}

type DeviceGetMcuRailVoltage struct{}

type DeviceStateMcuRailVoltage struct {
	Voltage uint32
}

type DeviceReboot struct{}

type LightGet struct{}

type LightSet struct {
	Stream   uint8
	Color    LightHsbk
	Duration uint32 // ms
}

type LightSetWaveform struct {
	Stream    uint8
	Transient uint8 // 0 false; 1+ true
	Color     LightHsbk
	Period    uint32 // ms per cycle
	DutyCycle int16
	Waveform  uint8
}

type LightSetDimAbsolute struct {
	Brightness int32  // 0 for no change
	Duration   uint32 // ms
}

type LightSetDimRelative struct {
	Brightness int32  // 0 for no change
	Duration   uint32 // ms
}

type LightSetRgbw struct {
	Color struct {
		Red, Green, Blue, White uint16
	}
}

type LightState struct {
	Color LightHsbk
	Dim   int16
	Power uint16
	Label label
	Tags  uint64
}

type LightGetRailVoltage struct{}

type LightStateRailVoltage struct {
	Voltage uint32
}

type LightGetTemperature struct{}

type LightStateTemperature struct {
	Temperature int16 // deci-celsius
}

type WanConnectPlain struct {
	User, Pass [32]byte
}

type WanConnectKey struct {
	AuthKey [32]byte
}

type WanStateConnect struct {
	AuthKey [32]byte
}

type WanSub struct {
	Target [8]byte
	Site   site
	Device uint8 // 0 device; 1 tag
}

type WanUnsub struct {
	Target [8]byte
	Site   site
	Device uint8 // 0 device; 1 tag
}

type WanStateSub struct {
	Target [8]byte
	Site   site
	Device uint8 // 0 device; 1 tag
}

type WifiGet struct {
	Iface ifaceType
}

type WifiSet struct {
	Iface  ifaceType
	Active uint8 // 0 false; 1 true
}

type WifiState struct {
	Iface  ifaceType
	Status wifiStatus
	Ipv4   uint32
	Ipv6   [16]byte
}

type WifiGetAccessPoint struct{}

type WifiSetAccessPoint struct {
	Iface    ifaceType
	Ssid     [32]byte
	Password [64]byte
	Security apSecurity
}

type WifiStateAccessPoint struct {
	Iface    ifaceType
	Ssid     [32]byte
	Password [64]byte
	Security apSecurity
	Channel  uint16
}

type SensorGetAmbientLight struct{}

type SensorStateAmbientLight struct {
	Lux float32
}

type SensorGetDimmerVoltage struct{}

type SensorStateDimmerVoltage struct {
	Voltage uint32
}
