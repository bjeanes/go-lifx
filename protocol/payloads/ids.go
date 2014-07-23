package payloads

const (
	DeviceSetSiteID             uint16 = 1
	DeviceGetPanGatewayID              = 2
	DeviceStatePanGatewayID            = 3
	DeviceTimeID                       = 4
	DeviceSetTimeID                    = 5
	DeviceStateTimeID                  = 6
	DeviceGetResetSwitchID             = 7
	DeviceStateResetSwitchID           = 8
	DeviceGetMeshInfoID                = 12
	DeviceStateMeshInfoID              = 13
	DeviceGetMeshFirmwareID            = 14
	DeviceStateMeshFirmwareID          = 15
	DeviceGetWifiInfoID                = 16
	DeviceStateWifiInfoID              = 17
	DeviceGetWifiFirmwareID            = 18
	DeviceStateWifiFirmwareID          = 19
	DeviceGetPowerID                   = 20
	DeviceSetPowerID                   = 21
	DeviceStatePowerID                 = 22
	DeviceGetLabelID                   = 23
	DeviceSetLabelID                   = 24
	DeviceStateLabelID                 = 25
	DeviceGetTagsID                    = 26
	DeviceSetTagsID                    = 27
	DeviceStateTagsID                  = 28
	DeviceGetTagLabelsID               = 29
	DeviceSetTagLabelsID               = 30
	DeviceStateTagLabelsID             = 31
	DeviceGetVersionID                 = 32
	DeviceStateVersionID               = 33
	DeviceGetInfoID                    = 34
	DeviceStateInfoID                  = 35
	DeviceGetMcuRailVoltageID          = 36
	DeviceStateMcuRailVoltageID        = 37
	DeviceRebootID                     = 38
	LightGetID                         = 101
	LightSetID                         = 102
	LightSetWaveformID                 = 103
	LightSetDimAbsoluteID              = 104
	LightSetDimRelativeID              = 105
	LightSetRgbwID                     = 106
	LightStateID                       = 107
	LightGetRailVoltageID              = 108
	LightStateRailVoltageID            = 109
	LightGetTemperatureID              = 110
	LightStateTemperatureID            = 111
	WanConnectPlainID                  = 201
	WanConnectKeyID                    = 202
	WanStateConnectID                  = 203
	WanSubID                           = 204
	WanUnsubID                         = 205
	WanStateSubID                      = 206
	WifiGetID                          = 301
	WifiSetID                          = 302
	WifiStateID                        = 303
	WifiGetAccessPointID               = 304
	WifiSetAccessPointID               = 305
	WifiStateAccessPointID             = 306
	SensorGetAmbientLightID            = 401
	SensorStateAmbientLightID          = 402
	SensorGetDimmerVoltageID           = 403
	SensorStateDimmerVoltageID         = 404
)

func ForId(id uint16) Payload {
	switch id {
	case DeviceSetSiteID:
		return new(DeviceSetSite)
	case DeviceGetPanGatewayID:
		return new(DeviceGetPanGateway)
	case DeviceStatePanGatewayID:
		return new(DeviceStatePanGateway)
	case DeviceTimeID:
		return new(DeviceTime)
	case DeviceSetTimeID:
		return new(DeviceSetTime)
	case DeviceStateTimeID:
		return new(DeviceStateTime)
	case DeviceGetResetSwitchID:
		return new(DeviceGetResetSwitch)
	case DeviceStateResetSwitchID:
		return new(DeviceStateResetSwitch)
	case DeviceGetMeshInfoID:
		return new(DeviceGetMeshInfo)
	case DeviceStateMeshInfoID:
		return new(DeviceStateMeshInfo)
	case DeviceGetMeshFirmwareID:
		return new(DeviceGetMeshFirmware)
	case DeviceStateMeshFirmwareID:
		return new(DeviceStateMeshFirmware)
	case DeviceGetWifiInfoID:
		return new(DeviceGetWifiInfo)
	case DeviceStateWifiInfoID:
		return new(DeviceStateWifiInfo)
	case DeviceGetWifiFirmwareID:
		return new(DeviceGetWifiFirmware)
	case DeviceStateWifiFirmwareID:
		return new(DeviceStateWifiFirmware)
	case DeviceGetPowerID:
		return new(DeviceGetPower)
	case DeviceSetPowerID:
		return new(DeviceSetPower)
	case DeviceStatePowerID:
		return new(DeviceStatePower)
	case DeviceGetLabelID:
		return new(DeviceGetLabel)
	case DeviceSetLabelID:
		return new(DeviceSetLabel)
	case DeviceStateLabelID:
		return new(DeviceStateLabel)
	case DeviceGetTagsID:
		return new(DeviceGetTags)
	case DeviceSetTagsID:
		return new(DeviceSetTags)
	case DeviceStateTagsID:
		return new(DeviceStateTags)
	case DeviceGetTagLabelsID:
		return new(DeviceGetTagLabels)
	case DeviceSetTagLabelsID:
		return new(DeviceSetTagLabels)
	case DeviceStateTagLabelsID:
		return new(DeviceStateTagLabels)
	case DeviceGetVersionID:
		return new(DeviceGetVersion)
	case DeviceStateVersionID:
		return new(DeviceStateVersion)
	case DeviceGetInfoID:
		return new(DeviceGetInfo)
	case DeviceStateInfoID:
		return new(DeviceStateInfo)
	case DeviceGetMcuRailVoltageID:
		return new(DeviceGetMcuRailVoltage)
	case DeviceStateMcuRailVoltageID:
		return new(DeviceStateMcuRailVoltage)
	case DeviceRebootID:
		return new(DeviceReboot)
	case LightGetID:
		return new(LightGet)
	case LightSetID:
		return new(LightSet)
	case LightSetWaveformID:
		return new(LightSetWaveform)
	case LightSetDimAbsoluteID:
		return new(LightSetDimAbsolute)
	case LightSetDimRelativeID:
		return new(LightSetDimRelative)
	case LightSetRgbwID:
		return new(LightSetRgbw)
	case LightStateID:
		return new(LightState)
	case LightGetRailVoltageID:
		return new(LightGetRailVoltage)
	case LightStateRailVoltageID:
		return new(LightStateRailVoltage)
	case LightGetTemperatureID:
		return new(LightGetTemperature)
	case LightStateTemperatureID:
		return new(LightStateTemperature)
	case WanConnectPlainID:
		return new(WanConnectPlain)
	case WanConnectKeyID:
		return new(WanConnectKey)
	case WanStateConnectID:
		return new(WanStateConnect)
	case WanSubID:
		return new(WanSub)
	case WanUnsubID:
		return new(WanUnsub)
	case WanStateSubID:
		return new(WanStateSub)
	case WifiGetID:
		return new(WifiGet)
	case WifiSetID:
		return new(WifiSet)
	case WifiStateID:
		return new(WifiState)
	case WifiGetAccessPointID:
		return new(WifiGetAccessPoint)
	case WifiSetAccessPointID:
		return new(WifiSetAccessPoint)
	case WifiStateAccessPointID:
		return new(WifiStateAccessPoint)
	case SensorGetAmbientLightID:
		return new(SensorGetAmbientLight)
	case SensorStateAmbientLightID:
		return new(SensorStateAmbientLight)
	case SensorGetDimmerVoltageID:
		return new(SensorGetDimmerVoltage)
	case SensorStateDimmerVoltageID:
		return new(SensorStateDimmerVoltage)
	default:
		return nil
	}
}

func (DeviceSetSite) Id() uint16             { return DeviceSetSiteID }
func (DeviceGetPanGateway) Id() uint16       { return DeviceGetPanGatewayID }
func (DeviceStatePanGateway) Id() uint16     { return DeviceStatePanGatewayID }
func (DeviceTime) Id() uint16                { return DeviceTimeID }
func (DeviceSetTime) Id() uint16             { return DeviceSetTimeID }
func (DeviceStateTime) Id() uint16           { return DeviceStateTimeID }
func (DeviceGetResetSwitch) Id() uint16      { return DeviceGetResetSwitchID }
func (DeviceStateResetSwitch) Id() uint16    { return DeviceStateResetSwitchID }
func (DeviceGetMeshInfo) Id() uint16         { return DeviceGetMeshInfoID }
func (DeviceStateMeshInfo) Id() uint16       { return DeviceStateMeshInfoID }
func (DeviceGetMeshFirmware) Id() uint16     { return DeviceGetMeshFirmwareID }
func (DeviceStateMeshFirmware) Id() uint16   { return DeviceStateMeshFirmwareID }
func (DeviceGetWifiInfo) Id() uint16         { return DeviceGetWifiInfoID }
func (DeviceStateWifiInfo) Id() uint16       { return DeviceStateWifiInfoID }
func (DeviceGetWifiFirmware) Id() uint16     { return DeviceGetWifiFirmwareID }
func (DeviceStateWifiFirmware) Id() uint16   { return DeviceStateWifiFirmwareID }
func (DeviceGetPower) Id() uint16            { return DeviceGetPowerID }
func (DeviceSetPower) Id() uint16            { return DeviceSetPowerID }
func (DeviceStatePower) Id() uint16          { return DeviceStatePowerID }
func (DeviceGetLabel) Id() uint16            { return DeviceGetLabelID }
func (DeviceSetLabel) Id() uint16            { return DeviceSetLabelID }
func (DeviceStateLabel) Id() uint16          { return DeviceStateLabelID }
func (DeviceGetTags) Id() uint16             { return DeviceGetTagsID }
func (DeviceSetTags) Id() uint16             { return DeviceSetTagsID }
func (DeviceStateTags) Id() uint16           { return DeviceStateTagsID }
func (DeviceGetTagLabels) Id() uint16        { return DeviceGetTagLabelsID }
func (DeviceSetTagLabels) Id() uint16        { return DeviceSetTagLabelsID }
func (DeviceStateTagLabels) Id() uint16      { return DeviceStateTagLabelsID }
func (DeviceGetVersion) Id() uint16          { return DeviceGetVersionID }
func (DeviceStateVersion) Id() uint16        { return DeviceStateVersionID }
func (DeviceGetInfo) Id() uint16             { return DeviceGetInfoID }
func (DeviceStateInfo) Id() uint16           { return DeviceStateInfoID }
func (DeviceGetMcuRailVoltage) Id() uint16   { return DeviceGetMcuRailVoltageID }
func (DeviceStateMcuRailVoltage) Id() uint16 { return DeviceStateMcuRailVoltageID }
func (DeviceReboot) Id() uint16              { return DeviceRebootID }
func (LightGet) Id() uint16                  { return LightGetID }
func (LightSet) Id() uint16                  { return LightSetID }
func (LightSetWaveform) Id() uint16          { return LightSetWaveformID }
func (LightSetDimAbsolute) Id() uint16       { return LightSetDimAbsoluteID }
func (LightSetDimRelative) Id() uint16       { return LightSetDimRelativeID }
func (LightSetRgbw) Id() uint16              { return LightSetRgbwID }
func (LightState) Id() uint16                { return LightStateID }
func (LightGetRailVoltage) Id() uint16       { return LightGetRailVoltageID }
func (LightStateRailVoltage) Id() uint16     { return LightStateRailVoltageID }
func (LightGetTemperature) Id() uint16       { return LightGetTemperatureID }
func (LightStateTemperature) Id() uint16     { return LightStateTemperatureID }
func (WanConnectPlain) Id() uint16           { return WanConnectPlainID }
func (WanConnectKey) Id() uint16             { return WanConnectKeyID }
func (WanStateConnect) Id() uint16           { return WanStateConnectID }
func (WanSub) Id() uint16                    { return WanSubID }
func (WanUnsub) Id() uint16                  { return WanUnsubID }
func (WanStateSub) Id() uint16               { return WanStateSubID }
func (WifiGet) Id() uint16                   { return WifiGetID }
func (WifiSet) Id() uint16                   { return WifiSetID }
func (WifiState) Id() uint16                 { return WifiStateID }
func (WifiGetAccessPoint) Id() uint16        { return WifiGetAccessPointID }
func (WifiSetAccessPoint) Id() uint16        { return WifiSetAccessPointID }
func (WifiStateAccessPoint) Id() uint16      { return WifiStateAccessPointID }
func (SensorGetAmbientLight) Id() uint16     { return SensorGetAmbientLightID }
func (SensorStateAmbientLight) Id() uint16   { return SensorStateAmbientLightID }
func (SensorGetDimmerVoltage) Id() uint16    { return SensorGetDimmerVoltageID }
func (SensorStateDimmerVoltage) Id() uint16  { return SensorStateDimmerVoltageID }
