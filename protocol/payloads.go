package protocol

import type_registry "github.com/bjeanes/go-lifx/type_registry"

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

	payload interface{}

	lightHsbk struct {
		Hue        uint16 // 0-65535 scaled to 0-360°
		Saturation uint16 // 0-65535 scaled to 0-100%
		Brightness uint16 // 0-65535 scaled to 0-100%
		Kelvin     uint16 // absolute 2400-10000
	}
)

func (h header) version() uint16 {
	return 0xfff & uint16(h.Bitfield1) // top 12 bits
}

var payloads type_registry.TypeRegistry

func init() {
	payloads = type_registry.New()

	type deviceSetSite struct {
		Site [6]byte
	}
	payloads.Register(1, (*deviceSetSite)(nil))

	type deviceGetPanGateway struct{}
	payloads.Register(2, (*deviceGetPanGateway)(nil))

	type deviceStatePanGateway struct {
		Service uint8  // 1 for UDP; 2 for TCP
		Port    uint32 // 0 for disabled, else it's the service port
	}
	payloads.Register(3, (*deviceStatePanGateway)(nil))

	type deviceTime struct{}
	payloads.Register(4, (*deviceTime)(nil))

	type deviceSetTime struct {
		Time uint64 // nanos since epoch
	}
	payloads.Register(5, (*deviceSetTime)(nil))

	type deviceStateTime struct {
		Time uint64 // nanos since epoch
	}
	payloads.Register(6, (*deviceStateTime)(nil))

	type deviceGetResetSwitch struct{}
	payloads.Register(7, (*deviceGetResetSwitch)(nil))

	type deviceStateResetSwitch struct {
		Position uint8
	}
	payloads.Register(8, (*deviceStateResetSwitch)(nil))

	type deviceGetMeshInfo struct{}
	payloads.Register(12, (*deviceGetMeshInfo)(nil))

	type deviceStateMeshInfo struct {
		Signal         float32 // milliwatts
		Tx, Rx         uint32  // bytes
		McuTemperature int16   // deci-celsius
	}
	payloads.Register(13, (*deviceStateMeshInfo)(nil))

	type deviceGetMeshFirmware struct{}
	payloads.Register(14, (*deviceGetMeshFirmware)(nil))

	type deviceStateMeshFirmware struct {
		Build, Install uint64
		Version        uint32
	}
	payloads.Register(15, (*deviceStateMeshFirmware)(nil))

	type deviceGetWifiInfo struct{}
	payloads.Register(16, (*deviceGetWifiInfo)(nil))

	type deviceStateWifiInfo struct {
		Signal         float32 // milliwatts
		Tx, Rx         uint32  // bytes
		McuTemperature int16   // deci-celsius
	}
	payloads.Register(17, (*deviceStateWifiInfo)(nil))

	type deviceGetWifiFirmware struct{}
	payloads.Register(18, (*deviceGetWifiFirmware)(nil))

	type deviceStateWifiFirmware struct {
		Build, Install uint64
		Version        uint32
	}
	payloads.Register(19, (*deviceStateWifiFirmware)(nil))
	type deviceGetPower struct{}
	payloads.Register(20, (*deviceGetPower)(nil))

	type deviceSetPower struct {
		Level uint16 // 0 = off; 1+ = on
	}
	payloads.Register(21, (*deviceSetPower)(nil))

	type deviceStatePower struct {
		Level uint16 // 0 = off; 1+ = on
	}
	payloads.Register(22, (*deviceStatePower)(nil))

	type deviceGetLabel struct{}
	payloads.Register(23, (*deviceGetLabel)(nil))

	type deviceSetLabel struct {
		Label [32]byte // string
	}
	payloads.Register(24, (*deviceSetLabel)(nil))

	type deviceStateLabel struct {
		Label [32]byte // string
	}
	payloads.Register(25, (*deviceStateLabel)(nil))

	type deviceGetTags struct{}
	payloads.Register(26, (*deviceGetTags)(nil))

	type deviceSetTags struct {
		Tags uint64
	}
	payloads.Register(27, (*deviceSetTags)(nil))

	type deviceStateTags struct {
		Tags uint64
	}
	payloads.Register(28, (*deviceStateTags)(nil))

	type deviceGetTagLabels struct {
		Tags uint64
	}
	payloads.Register(29, (*deviceGetTagLabels)(nil))

	type deviceSetTagLabels struct {
		Tags  uint64
		Label [32]byte
	}
	payloads.Register(30, (*deviceSetTagLabels)(nil))

	type deviceStateTagLabels struct {
		Tags  uint64
		Label [32]byte
	}
	payloads.Register(31, (*deviceStateTagLabels)(nil))

	type deviceGetVersion struct{}
	payloads.Register(32, (*deviceGetVersion)(nil))

	type deviceStateVersion struct {
		Vendor, Product, Version uint32
	}
	payloads.Register(33, (*deviceStateVersion)(nil))

	type deviceGetInfo struct{}
	payloads.Register(34, (*deviceGetInfo)(nil))

	type deviceStateInfo struct {
		Time, Uptime, Downtime uint64 // ns
	}
	payloads.Register(35, (*deviceStateInfo)(nil))

	type deviceGetMcuRailVoltage struct{}
	payloads.Register(36, (*deviceGetMcuRailVoltage)(nil))

	type deviceStateMcuRailVoltage struct {
		Voltage uint32
	}
	payloads.Register(37, (*deviceStateMcuRailVoltage)(nil))

	type deviceReboot struct{}
	payloads.Register(38, (*deviceReboot)(nil))

	type lightGet struct{}
	payloads.Register(101, (*lightGet)(nil))

	type lightSet struct {
		Stream   uint8
		Color    lightHsbk
		Duration uint32 // ms
	}
	payloads.Register(102, (*lightSet)(nil))

	type lightSetWaveform struct {
		Stream    uint8
		Transient uint8 // 0 false; 1+ true
		Color     lightHsbk
		Period    uint32 // ms per cycle
		DutyCycle int16
		Waveform  uint8
	}
	payloads.Register(103, (*lightSetWaveform)(nil))

	type lightSetDimAbsolute struct {
		Brightness int32  // 0 for no change
		Duration   uint32 // ms
	}
	payloads.Register(104, (*lightSetDimAbsolute)(nil))

	type lightSetDimRelative struct {
		Brightness int32  // 0 for no change
		Duration   uint32 // ms
	}
	payloads.Register(105, (*lightSetDimRelative)(nil))

	type lightSetRgbw struct {
		Color struct {
			Red, Green, Blue, White uint16
		}
	}

	type lightState struct {
		Color lightHsbk
		Dim   int16
		Power uint16
		Label [32]byte
		Tags  uint64
	}
	payloads.Register(107, (*lightState)(nil))

	type lightGetRailVoltage struct{}
	payloads.Register(108, (*lightGetRailVoltage)(nil))

	type lightStateRailVoltage struct {
		Voltage uint32
	}
	payloads.Register(109, (*lightStateRailVoltage)(nil))

	type lightGetTemperature struct{}
	payloads.Register(110, (*lightGetTemperature)(nil))

	type lightStateTemperature struct {
		Temperature int16 // deci-celsius
	}
	payloads.Register(111, (*lightStateTemperature)(nil))

	type wanConnectPlain struct {
		User, Pass [32]byte
	}
	payloads.Register(201, (*wanConnectPlain)(nil))

	type wanConnectKey struct {
		AuthKey [32]byte
	}
	payloads.Register(202, (*wanConnectKey)(nil))

	type wanStateConnect struct {
		AuthKey [32]byte
	}
	payloads.Register(203, (*wanStateConnect)(nil))

	type wanSub struct {
		Target [8]byte
		Site   [6]byte
		Device uint8 // 0 device; 1 tag
	}
	payloads.Register(204, (*wanSub)(nil))

	type wanUnsub struct {
		Target [8]byte
		Site   [6]byte
		Device uint8 // 0 device; 1 tag
	}
	payloads.Register(205, (*wanUnsub)(nil))

	type wanStateSub struct {
		Target [8]byte
		Site   [6]byte
		Device uint8 // 0 device; 1 tag
	}
	payloads.Register(206, (*wanStateSub)(nil))

	type wifiGet struct {
		Iface ifaceType
	}
	payloads.Register(301, (*wifiGet)(nil))

	type wifiSet struct {
		Iface  ifaceType
		Active uint8 // 0 false; 1 true
	}
	payloads.Register(302, (*wifiSet)(nil))

	type wifiState struct {
		Iface  ifaceType
		Status wifiStatus
		Ipv4   uint32
		Ipv6   [16]byte
	}
	payloads.Register(303, (*wifiState)(nil))

	type wifiGetAccessPoint struct{}
	payloads.Register(304, (*wifiGetAccessPoint)(nil))

	type wifiSetAccessPoint struct {
		Iface    ifaceType
		Ssid     [32]byte
		Password [64]byte
		Security apSecurity
	}
	payloads.Register(305, (*wifiSetAccessPoint)(nil))

	type wifiStateAccessPoint struct {
		Iface    ifaceType
		Ssid     [32]byte
		Password [64]byte
		Security apSecurity
		Channel  uint16
	}
	payloads.Register(306, (*wifiStateAccessPoint)(nil))

	type sensorGetAmbientLight struct{}
	payloads.Register(401, (*sensorGetAmbientLight)(nil))

	type sensorStateAmbientLight struct {
		Lux float32
	}
	payloads.Register(402, (*sensorStateAmbientLight)(nil))

	type sensorGetDimmerVoltage struct{}
	payloads.Register(403, (*sensorGetDimmerVoltage)(nil))

	type sensorStateDimmerVoltage struct {
		Voltage uint32
	}
	payloads.Register(404, (*sensorStateDimmerVoltage)(nil))
}
