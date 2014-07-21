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

func ForId(id uint16) Payload {
	switch id {
	case DeviceSetSite:
		return new(deviceSetSite)
	case DeviceGetPanGateway:
		return new(deviceGetPanGateway)
	case DeviceStatePanGateway:
		return new(deviceStatePanGateway)
	case DeviceTime:
		return new(deviceTime)
	case DeviceSetTime:
		return new(deviceSetTime)
	case DeviceStateTime:
		return new(deviceStateTime)
	case DeviceGetResetSwitch:
		return new(deviceGetResetSwitch)
	case DeviceStateResetSwitch:
		return new(deviceStateResetSwitch)
	case DeviceGetMeshInfo:
		return new(deviceGetMeshInfo)
	case DeviceStateMeshInfo:
		return new(deviceStateMeshInfo)
	case DeviceGetMeshFirmware:
		return new(deviceGetMeshFirmware)
	case DeviceStateMeshFirmware:
		return new(deviceStateMeshFirmware)
	case DeviceGetWifiInfo:
		return new(deviceGetWifiInfo)
	case DeviceStateWifiInfo:
		return new(deviceStateWifiInfo)
	case DeviceGetWifiFirmware:
		return new(deviceGetWifiFirmware)
	case DeviceStateWifiFirmware:
		return new(deviceStateWifiFirmware)
	case DeviceGetPower:
		return new(deviceGetPower)
	case DeviceSetPower:
		return new(deviceSetPower)
	case DeviceStatePower:
		return new(deviceStatePower)
	case DeviceGetLabel:
		return new(deviceGetLabel)
	case DeviceSetLabel:
		return new(deviceSetLabel)
	case DeviceStateLabel:
		return new(deviceStateLabel)
	case DeviceGetTags:
		return new(deviceGetTags)
	case DeviceSetTags:
		return new(deviceSetTags)
	case DeviceStateTags:
		return new(deviceStateTags)
	case DeviceGetTagLabels:
		return new(deviceGetTagLabels)
	case DeviceSetTagLabels:
		return new(deviceSetTagLabels)
	case DeviceStateTagLabels:
		return new(deviceStateTagLabels)
	case DeviceGetVersion:
		return new(deviceGetVersion)
	case DeviceStateVersion:
		return new(deviceStateVersion)
	case DeviceGetInfo:
		return new(deviceGetInfo)
	case DeviceStateInfo:
		return new(deviceStateInfo)
	case DeviceGetMcuRailVoltage:
		return new(deviceGetMcuRailVoltage)
	case DeviceStateMcuRailVoltage:
		return new(deviceStateMcuRailVoltage)
	case DeviceReboot:
		return new(deviceReboot)
	case LightGet:
		return new(lightGet)
	case LightSet:
		return new(lightSet)
	case LightSetWaveform:
		return new(lightSetWaveform)
	case LightSetDimAbsolute:
		return new(lightSetDimAbsolute)
	case LightSetDimRelative:
		return new(lightSetDimRelative)
	case LightSetRgbw:
		return new(lightSetRgbw)
	case LightState:
		return new(lightState)
	case LightGetRailVoltage:
		return new(lightGetRailVoltage)
	case LightStateRailVoltage:
		return new(lightStateRailVoltage)
	case LightGetTemperature:
		return new(lightGetTemperature)
	case LightStateTemperature:
		return new(lightStateTemperature)
	case WanConnectPlain:
		return new(wanConnectPlain)
	case WanConnectKey:
		return new(wanConnectKey)
	case WanStateConnect:
		return new(wanStateConnect)
	case WanSub:
		return new(wanSub)
	case WanUnsub:
		return new(wanUnsub)
	case WanStateSub:
		return new(wanStateSub)
	case WifiGet:
		return new(wifiGet)
	case WifiSet:
		return new(wifiSet)
	case WifiState:
		return new(wifiState)
	case WifiGetAccessPoint:
		return new(wifiGetAccessPoint)
	case WifiSetAccessPoint:
		return new(wifiSetAccessPoint)
	case WifiStateAccessPoint:
		return new(wifiStateAccessPoint)
	case SensorGetAmbientLight:
		return new(sensorGetAmbientLight)
	case SensorStateAmbientLight:
		return new(sensorStateAmbientLight)
	case SensorGetDimmerVoltage:
		return new(sensorGetDimmerVoltage)
	case SensorStateDimmerVoltage:
		return new(sensorStateDimmerVoltage)
	default:
		return nil
	}
}

func (deviceSetSite) Id() uint16             { return DeviceSetSite }
func (deviceGetPanGateway) Id() uint16       { return DeviceGetPanGateway }
func (deviceStatePanGateway) Id() uint16     { return DeviceStatePanGateway }
func (deviceTime) Id() uint16                { return DeviceTime }
func (deviceSetTime) Id() uint16             { return DeviceSetTime }
func (deviceStateTime) Id() uint16           { return DeviceStateTime }
func (deviceGetResetSwitch) Id() uint16      { return DeviceGetResetSwitch }
func (deviceStateResetSwitch) Id() uint16    { return DeviceStateResetSwitch }
func (deviceGetMeshInfo) Id() uint16         { return DeviceGetMeshInfo }
func (deviceStateMeshInfo) Id() uint16       { return DeviceStateMeshInfo }
func (deviceGetMeshFirmware) Id() uint16     { return DeviceGetMeshFirmware }
func (deviceStateMeshFirmware) Id() uint16   { return DeviceStateMeshFirmware }
func (deviceGetWifiInfo) Id() uint16         { return DeviceGetWifiInfo }
func (deviceStateWifiInfo) Id() uint16       { return DeviceStateWifiInfo }
func (deviceGetWifiFirmware) Id() uint16     { return DeviceGetWifiFirmware }
func (deviceStateWifiFirmware) Id() uint16   { return DeviceStateWifiFirmware }
func (deviceGetPower) Id() uint16            { return DeviceGetPower }
func (deviceSetPower) Id() uint16            { return DeviceSetPower }
func (deviceStatePower) Id() uint16          { return DeviceStatePower }
func (deviceGetLabel) Id() uint16            { return DeviceGetLabel }
func (deviceSetLabel) Id() uint16            { return DeviceSetLabel }
func (deviceStateLabel) Id() uint16          { return DeviceStateLabel }
func (deviceGetTags) Id() uint16             { return DeviceGetTags }
func (deviceSetTags) Id() uint16             { return DeviceSetTags }
func (deviceStateTags) Id() uint16           { return DeviceStateTags }
func (deviceGetTagLabels) Id() uint16        { return DeviceGetTagLabels }
func (deviceSetTagLabels) Id() uint16        { return DeviceSetTagLabels }
func (deviceStateTagLabels) Id() uint16      { return DeviceStateTagLabels }
func (deviceGetVersion) Id() uint16          { return DeviceGetVersion }
func (deviceStateVersion) Id() uint16        { return DeviceStateVersion }
func (deviceGetInfo) Id() uint16             { return DeviceGetInfo }
func (deviceStateInfo) Id() uint16           { return DeviceStateInfo }
func (deviceGetMcuRailVoltage) Id() uint16   { return DeviceGetMcuRailVoltage }
func (deviceStateMcuRailVoltage) Id() uint16 { return DeviceStateMcuRailVoltage }
func (deviceReboot) Id() uint16              { return DeviceReboot }
func (lightGet) Id() uint16                  { return LightGet }
func (lightSet) Id() uint16                  { return LightSet }
func (lightSetWaveform) Id() uint16          { return LightSetWaveform }
func (lightSetDimAbsolute) Id() uint16       { return LightSetDimAbsolute }
func (lightSetDimRelative) Id() uint16       { return LightSetDimRelative }
func (lightSetRgbw) Id() uint16              { return LightSetRgbw }
func (lightState) Id() uint16                { return LightState }
func (lightGetRailVoltage) Id() uint16       { return LightGetRailVoltage }
func (lightStateRailVoltage) Id() uint16     { return LightStateRailVoltage }
func (lightGetTemperature) Id() uint16       { return LightGetTemperature }
func (lightStateTemperature) Id() uint16     { return LightStateTemperature }
func (wanConnectPlain) Id() uint16           { return WanConnectPlain }
func (wanConnectKey) Id() uint16             { return WanConnectKey }
func (wanStateConnect) Id() uint16           { return WanStateConnect }
func (wanSub) Id() uint16                    { return WanSub }
func (wanUnsub) Id() uint16                  { return WanUnsub }
func (wanStateSub) Id() uint16               { return WanStateSub }
func (wifiGet) Id() uint16                   { return WifiGet }
func (wifiSet) Id() uint16                   { return WifiSet }
func (wifiState) Id() uint16                 { return WifiState }
func (wifiGetAccessPoint) Id() uint16        { return WifiGetAccessPoint }
func (wifiSetAccessPoint) Id() uint16        { return WifiSetAccessPoint }
func (wifiStateAccessPoint) Id() uint16      { return WifiStateAccessPoint }
func (sensorGetAmbientLight) Id() uint16     { return SensorGetAmbientLight }
func (sensorStateAmbientLight) Id() uint16   { return SensorStateAmbientLight }
func (sensorGetDimmerVoltage) Id() uint16    { return SensorGetDimmerVoltage }
func (sensorStateDimmerVoltage) Id() uint16  { return SensorStateDimmerVoltage }
