package protocol

import (
  . "testing"
  "github.com/bjeanes/go-lifx/protocol/payloads"
)

func TestDecodeDeviceGetPanGateway(t *T) {
  // DATA: length=36
  //       000  24 00 00 34 00 00 00 00  00 00 00 00 00 00 00 00  |$..4............|
  //       010  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
  //       020  02 00 00 00                                       |....|

  b := []byte{
    0x24, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x02, 0x00, 0x00, 0x00,
  }

  msg, err := Decode(b)

  if err != nil {
    t.Error("Decode failed with err: " + err.Error())
  }

  if msg.atTime != 0 {
    t.Error("atTime field incorrect")
  }

  if msg.version != 1024 {
    t.Error("protocol version field incorrect")
  }

  if !msg.addressable {
    t.Error("adressable field should be true")
  }
}

func TestDecodeDeviceStatePanGatewayService1(t *T) {
  // DATA: length=41
  //       000  29 00 00 54 00 00 00 00  d0 73 d5 00 f9 14 00 00  |)..T.....s......|
  //       010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
  //       020  03 00 00 00 01 7c dd 00  00                       |.....|...|
  b := []byte{
    0x29, 0x00, 0x00, 0x54, 0x00, 0x00, 0x00, 0x00,
    0xd0, 0x73, 0xd5, 0x00, 0xf9, 0x14, 0x00, 0x00,
    0x4c, 0x49, 0x46, 0x58, 0x56, 0x32, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x03, 0x00, 0x00, 0x00, 0x01, 0x7c, 0xdd, 0x00,
    0x00,
  }

  msg, err := Decode(b)

  if err != nil {
    t.Error("Decode failed with err: " + err.Error())
  }

  if msg.atTime != 0 {
    t.Error("atTime field incorrect")
  }

  if msg.version != 1024 {
    t.Error("protocol version field incorrect")
  }

  if !msg.addressable {
    t.Error("adressable field should be true")
  }

  payload := msg.Payload.(*payloads.DeviceStatePanGateway)

  if payload.Service != 1 {
    t.Error("Service incorrect")
  }

  if payload.Port != 56700 {
    t.Error("Port was not 56700")
  }

}

func TestDecodeDeviceStatePanGatewayService2(t *T) {
  // DATA: length=41
  //       000  29 00 00 54 00 00 00 00  d0 73 d5 00 f9 14 00 00  |)..T.....s......|
  //       010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
  //       020  03 00 00 00 02 00 00 00  00                       |.........|

  b := []byte{
    0x29, 0x00, 0x00, 0x54, 0x00, 0x00, 0x00, 0x00,
    0xd0, 0x73, 0xd5, 0x00, 0xf9, 0x14, 0x00, 0x00,
    0x4c, 0x49, 0x46, 0x58, 0x56, 0x32, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x03, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
    0x00,
  }

  msg, err := Decode(b)

  if err != nil {
    t.Error("Decode failed with err: " + err.Error())
  }

  if msg.atTime != 0 {
    t.Error("atTime field incorrect")
  }

  if msg.version != 1024 {
    t.Error("protocol version field incorrect")
  }

  if !msg.addressable {
    t.Error("adressable field should be true")
  }

  payload := msg.Payload.(*payloads.DeviceStatePanGateway)

  if payload.Service != 2 {
    t.Error("Service incorrect")
  }

  if payload.Port != 0 {
    t.Error("Port incorrect")
  }
}

func TestDecodeDeviceStateTime(t *T) {
  b := []byte{
    0x2c, 0x00, 0x00, 0x54, 0x00, 0x00, 0x00, 0x00, 
    0xd0, 0x73, 0xd5, 0x00, 0xf9, 0x14, 0x00, 0x00,

    0x4c, 0x49, 0x46, 0x58, 0x56, 0x32, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

    0x06, 0x00, 0x00, 0x00, 0xc8, 0x94, 0x8d, 0x8a,
    0x86, 0x2c, 0x83, 0x13,
  }

  msg, err := Decode(b)

  if err != nil {
    t.Error("Decode failed with err: " + err.Error())
  }

  if msg.atTime != 0 {
    t.Error("atTime field incorrect")
  }

  if msg.version != 1024 {
    t.Error("protocol version field incorrect")
  }

  if !msg.addressable {
    t.Error("adressable field should be true")
  }
}