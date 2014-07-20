package payloads

const (
	DeviceSetSite             uint16 = 1
	DeviceGetPanGateway              = 2
	DeviceStatePanGateway            = 3
	DeviceTime                       = 4
	DeviceSetTime                    = 5
	DeviceStateTime                  = 6
	DeviceGetResetSwitch             = 7
	DeviceStateResetSwitch           = 8
	DeviceGetMeshInfo                = 12
	DeviceStateMeshInfo              = 13
	DeviceGetMeshFirmware            = 14
	DeviceStateMeshFirmware          = 15
	DeviceGetWifiInfo                = 16
	DeviceStateWifiInfo              = 17
	DeviceGetWifiFirmware            = 18
	DeviceStateWifiFirmware          = 19
	DeviceGetPower                   = 20
	DeviceSetPower                   = 21
	DeviceStatePower                 = 22
	DeviceGetLabel                   = 23
	DeviceSetLabel                   = 24
	DeviceStateLabel                 = 25
	DeviceGetTags                    = 26
	DeviceSetTags                    = 27
	DeviceStateTags                  = 28
	DeviceGetTagLabels               = 29
	DeviceSetTagLabels               = 30
	DeviceStateTagLabels             = 31
	DeviceGetVersion                 = 32
	DeviceStateVersion               = 33
	DeviceGetInfo                    = 34
	DeviceStateInfo                  = 35
	DeviceGetMcuRailVoltage          = 36
	DeviceStateMcuRailVoltage        = 37
	DeviceReboot                     = 38
	LightGet                         = 101
	LightSet                         = 102
	LightSetWaveform                 = 103
	LightSetDimAbsolute              = 104
	LightSetDimRelative              = 105
	LightSetRgbw                     = 106
	LightState                       = 107
	LightGetRailVoltage              = 108
	LightStateRailVoltage            = 109
	LightGetTemperature              = 110
	LightStateTemperature            = 111
	WanConnectPlain                  = 201
	WanConnectKey                    = 202
	WanStateConnect                  = 203
	WanSub                           = 204
	WanUnsub                         = 205
	WanStateSub                      = 206
	WifiGet                          = 301
	WifiSet                          = 302
	WifiState                        = 303
	WifiGetAccessPoint               = 304
	WifiSetAccessPoint               = 305
	WifiStateAccessPoint             = 306
	SensorGetAmbientLight            = 401
	SensorStateAmbientLight          = 402
	SensorGetDimmerVoltage           = 403
	SensorStateDimmerVoltage         = 404
)

type Payload interface {
	Id() uint16
}

func ForId(id uint16) Payload {
	switch id {
	case DeviceGetPower:
		return new(deviceGetPower)
	case DeviceSetPower:
		return new(deviceSetPower)
	case DeviceStatePower:
		return new(deviceStatePower)
	case LightState:
		return new(lightState)
	default:
		return nil
	}
}
