package protocol

type (

  // Payload name                         ID
  plDeviceSetSite struct{              // 1
    site [6]byte
  }

  plDeviceGetPanGateway struct{}       // 2

  plDeviceStatePanGateway struct{      // 3
    service uint8
    port uint32
  }

  plDevicetTime struct{}               // 4

  plDeviceSetTime struct{              // 5
    time uint64 // nanos since epoch
  }

  plDeviceStateTime struct{            // 6
    time uint64 // nanos since epoch
  }

  plDeviceGetResetSwitch struct{}      // 7

  plDeviceStateResetSwitch struct{     // 8
    position uint8
  }

  plDeviceGetMeshInfo struct{}         // 12

  plDeviceStateMeshInfo struct{        // 13
    signal float32 // milliwatts
    tx, rx uint32  // bytes
    mcu_temperature int16 // deci-celsius
  }

  plDeviceGetMeshFirmware struct{}     // 14

  plDeviceStateMeshFirmware struct{    // 15
    build, install uint64
    version uint32
  }

  plDeviceGetWifiInfo struct{}         // 16

  plDeviceStateWifiInfo struct{        // 17
    signal float32 // milliwatts
    tx, rx uint32  // bytes
    mcu_temperature int16 // deci-celsius
  }

  plDeviceGetWifiFirmware struct{}     // 18

  plDeviceStateWifiFirmware struct{    // 19
    build, install uint64
    version uint32
  }
  plDeviceGetPower struct{}            // 20

  plDeviceSetPower struct{             // 21
    level uint16 // 0 = off; 1+ = on
  }

  plDeviceStatePower struct{          // 22
    level uint16 // 0 = off; 1+ = on
  }

  plDeviceGetLabel struct{}            // 23

  plDeviceSetLabel struct{             // 24
    label [32]byte // string
  }

  plDeviceStateLabel struct{           // 25
    label [32]byte // string
  }

  plDeviceGetTags struct{}             // 26

  plDeviceSetTags struct{              // 27
    tags uint64
  }

  plDeviceStateTags struct{            // 28
    tags uint64
  }

  plDeviceGetTagLabels struct{         // 29
    tags uint64
  }

  plDeviceSetTagLabels struct{         // 30
    tags uint64
    label [32]byte
  }

  plDeviceStateTagLabels struct{       // 31
    tags uint64
    label [32]byte
  }

  plDeviceGetVersion struct{}          // 32

  plDeviceStateVersion struct{         // 33
    vendor, product, version uint32
  }

  plDeviceGetInfo struct{}             // 34

  plDeviceStateInfo struct{            // 35
    time, uptime, downtime uint64 // ns
  }

  plDeviceGetMcuRailVoltage struct{}   // 36

  plDeviceStateMcuRailVoltage struct{  // 37
    voltage uint32
  }

  plDeviceReboot struct{}              // 38

  lightHsbk struct {                   // not a top-level payload
    hue uint16 // 0-65535 scaled to 0-360°
    saturation uint16 // 0-65535 scaled to 0-100%
    brightness uint16 // 0-65535 scaled to 0-100%
    kenvin uint16 // absolute 2400-10000
  }

  plLightGet struct{}                  // 101

  plLightSet struct{                   // 102
    stream uint8
    color lightHsbk
    duration uint32 // ms
  }

  plLightSetWaveform struct{           // 103
    stream uint8
    transient uint8 // 0 false; 1+ true
    color lightHsbk
    period uint32 // ms per cycle
    dutyCycle int16
    waveform uint8
  }

  plLightSetDimAbsolute struct{        // 104
    brightness int32 // 0 for no change
    duration uint32 // ms
  }

  plLightSetDimRelative struct{        // 105
    brightness int32 // 0 for no change
    duration uint32 // ms
  }

  plLightSetRgbw struct{               // 106
    color struct {
      red, green, blue, white uint16
    }
  }

  plLightState struct{                 // 107
    color lightHsbk
    dim int16
    power uint16
    label [32]byte
    tags uint64
  }

  plLightGetRailVoltage struct{}       // 108

  plLightStateRailVoltage struct{      // 109
    voltage uint32
  }

  plLightGetTemperature struct{}       // 110

  plLightStateTemperature struct{      // 111
    temperature int16 // deci-celsius
  }

  plWanConnectPlain struct{            // 201
    user, pass [32]byte
  }

  plWanConnectKey struct{              // 202
    authKey [32]byte
  }

  plWanStateConnect struct{            // 203
    authKey [32]byte
  }

  plWanSub struct{                     // 204
    target [8]byte
    site   [6]byte
    device uint8 // 0 device; 1 tag
  }

  plWanUnsub struct{                   // 205
    target [8]byte
    site   [6]byte
    device uint8 // 0 device; 1 tag
  }

  plWanStateSub struct{                // 206
    target [8]byte
    site   [6]byte
    device uint8 // 0 device; 1 tag
  }

  plWifiGet struct{                    // 301
    iface ifaceType
  }

  plWifiSet struct{                    // 302
    iface ifaceType
    active uint8 // 0 false; 1 true
  }

  plWifiState struct{                  // 303
    iface ifaceType
    status wifiStatus
    ipv4 uint32
    ipv6 [16]byte
  }

  plWifiGetAccessPoint struct{}        // 304

  plWifiSetAccessPoint struct{         // 305
    iface ifaceType
    ssid [32]byte
    password [64]byte
    security apSecurity
  }

  plWifiStateAccessPoint struct{       // 306
    iface ifaceType
    ssid [32]byte
    password [64]byte
    security apSecurity
    channel uint16
  }

  plSensorGetAmbientLight struct{}     // 401

  plSensorStateAmbientLight struct{    // 402
    lux float32
  }

  plSensorGetDimmerVoltage struct{}    // 403

  plSensorStateDimmerVoltage struct{   // 404
    voltage uint32
  }
)

