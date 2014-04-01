package protocol

func initPayload(id uint16) (payload) {
  switch id {
  case 2:
    return &plDeviceGetPanGateway{}
  case 3:
    return &plDeviceStatePanGateway{}
  case 107:
    return &plLightState{}
  }

  debug("err: Unknown message payload of type %d\n", id)
  return nil
}

// payloads
type (
  // Payload name                         ID
  plDeviceSetSite struct{              // 1
    Site [6]byte
  }

  plDeviceGetPanGateway struct{}       // 2

  plDeviceStatePanGateway struct{      // 3
    Service uint8
    Port uint32
  }

  plDevicetTime struct{}               // 4

  plDeviceSetTime struct{              // 5
    Time uint64 // nanos since epoch
  }

  plDeviceStateTime struct{            // 6
    Time uint64 // nanos since epoch
  }

  plDeviceGetResetSwitch struct{}      // 7

  plDeviceStateResetSwitch struct{     // 8
    Position uint8
  }

  plDeviceGetMeshInfo struct{}         // 12

  plDeviceStateMeshInfo struct{        // 13
    Signal float32 // milliwatts
    Tx, Rx uint32  // bytes
    McuTemperature int16 // deci-celsius
  }

  plDeviceGetMeshFirmware struct{}     // 14

  plDeviceStateMeshFirmware struct{    // 15
    Build, Install uint64
    Version uint32
  }

  plDeviceGetWifiInfo struct{}         // 16

  plDeviceStateWifiInfo struct{        // 17
    Signal float32 // milliwatts
    Tx, Rx uint32  // bytes
    McuTemperature int16 // deci-celsius
  }

  plDeviceGetWifiFirmware struct{}     // 18

  plDeviceStateWifiFirmware struct{    // 19
    Build, Install uint64
    Version uint32
  }
  plDeviceGetPower struct{}            // 20

  plDeviceSetPower struct{             // 21
    Level uint16 // 0 = off; 1+ = on
  }

  plDeviceStatePower struct{          // 22
    Level uint16 // 0 = off; 1+ = on
  }

  plDeviceGetLabel struct{}            // 23

  plDeviceSetLabel struct{             // 24
    Label [32]byte // string
  }

  plDeviceStateLabel struct{           // 25
    Label [32]byte // string
  }

  plDeviceGetTags struct{}             // 26

  plDeviceSetTags struct{              // 27
    Tags uint64
  }

  plDeviceStateTags struct{            // 28
    Tags uint64
  }

  plDeviceGetTagLabels struct{         // 29
    Tags uint64
  }

  plDeviceSetTagLabels struct{         // 30
    Tags uint64
    Label [32]byte
  }

  plDeviceStateTagLabels struct{       // 31
    Tags uint64
    Label [32]byte
  }

  plDeviceGetVersion struct{}          // 32

  plDeviceStateVersion struct{         // 33
    Vendor, Product, Version uint32
  }

  plDeviceGetInfo struct{}             // 34

  plDeviceStateInfo struct{            // 35
    Time, Uptime, Downtime uint64 // ns
  }

  plDeviceGetMcuRailVoltage struct{}   // 36

  plDeviceStateMcuRailVoltage struct{  // 37
    Voltage uint32
  }

  plDeviceReboot struct{}              // 38

  lightHsbk struct {                   // not a top-level payload
    Hue uint16 // 0-65535 scaled to 0-360°
    Saturation uint16 // 0-65535 scaled to 0-100%
    Brightness uint16 // 0-65535 scaled to 0-100%
    Kelvin uint16 // absolute 2400-10000
  }

  plLightGet struct{}                  // 101

  plLightSet struct{                   // 102
    Stream uint8
    Color lightHsbk
    Duration uint32 // ms
  }

  plLightSetWaveform struct{           // 103
    Stream uint8
    Transient uint8 // 0 false; 1+ true
    Color lightHsbk
    Period uint32 // ms per cycle
    DutyCycle int16
    Waveform uint8
  }

  plLightSetDimAbsolute struct{        // 104
    Brightness int32 // 0 for no change
    Duration uint32 // ms
  }

  plLightSetDimRelative struct{        // 105
    Brightness int32 // 0 for no change
    Duration uint32 // ms
  }

  plLightSetRgbw struct{               // 106
    Color struct {
      Red, Green, Blue, White uint16
    }
  }

  plLightState struct{                 // 107
    Color lightHsbk
    Dim int16
    Power uint16
    Label [32]byte
    Tags uint64
  }

  plLightGetRailVoltage struct{}       // 108

  plLightStateRailVoltage struct{      // 109
    Voltage uint32
  }

  plLightGetTemperature struct{}       // 110

  plLightStateTemperature struct{      // 111
    Temperature int16 // deci-celsius
  }

  plWanConnectPlain struct{            // 201
    User, Pass [32]byte
  }

  plWanConnectKey struct{              // 202
    AuthKey [32]byte
  }

  plWanStateConnect struct{            // 203
    AuthKey [32]byte
  }

  plWanSub struct{                     // 204
    Target [8]byte
    Site   [6]byte
    Device uint8 // 0 device; 1 tag
  }

  plWanUnsub struct{                   // 205
    Target [8]byte
    Site   [6]byte
    Device uint8 // 0 device; 1 tag
  }

  plWanStateSub struct{                // 206
    Target [8]byte
    Site   [6]byte
    Device uint8 // 0 device; 1 tag
  }

  plWifiGet struct{                    // 301
    Iface ifaceType
  }

  plWifiSet struct{                    // 302
    Iface ifaceType
    Active uint8 // 0 false; 1 true
  }

  plWifiState struct{                  // 303
    Iface ifaceType
    Status wifiStatus
    Ipv4 uint32
    Ipv6 [16]byte
  }

  plWifiGetAccessPoint struct{}        // 304

  plWifiSetAccessPoint struct{         // 305
    Iface ifaceType
    Ssid [32]byte
    Password [64]byte
    Security apSecurity
  }

  plWifiStateAccessPoint struct{       // 306
    Iface ifaceType
    Ssid [32]byte
    Password [64]byte
    Security apSecurity
    Channel uint16
  }

  plSensorGetAmbientLight struct{}     // 401

  plSensorStateAmbientLight struct{    // 402
    Lux float32
  }

  plSensorGetDimmerVoltage struct{}    // 403

  plSensorStateDimmerVoltage struct{   // 404
    Voltage uint32
  }
)

