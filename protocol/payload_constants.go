package protocol

type (
  apSecurity uint8
  ifaceType uint8
  wifiStatus uint8
)

const(
  _            ifaceType = iota
  ifaceSoftAP
  ifaceStation
)

const (
  apSecurityUnknown apSecurity = iota
  apSecurityOpen
  apSecurityWEP_PSK
  apSecurityWPA_TKIP_PSK
  apSecurityWPA_AES_PSK
  apSecurityWPA2_AES_PSK
  apSecurityWPA2_TKIP_PSK
  apSecurityWPA2_MIXED_PSK
)

const (
  statusConnecting wifiStatus = iota
  statusConnected
  statusFailed
  statusOff
)

