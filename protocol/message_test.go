package protocol

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/bjeanes/go-lifx/protocol/payloads"
	. "testing"
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

	if msg.AtTime != 0 {
		t.Error("atTime field incorrect")
	}

	if msg.Version != 1024 {
		t.Error("protocol version field incorrect")
	}

	if !msg.Addressable {
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

	if msg.AtTime != 0 {
		t.Error("atTime field incorrect")
	}

	if msg.Version != 1024 {
		t.Error("protocol version field incorrect")
	}

	if !msg.Addressable {
		t.Error("adressable field should be true")
	}

	expectedTargetID := [8]byte{0xd0, 0x73, 0xd5, 0x00, 0xf9, 0x14, 0x00, 0x00}
	if msg.Header.Target != expectedTargetID {
		t.Error("Target incorrect")
	}

	expectedSiteID := [6]byte{0x4c, 0x49, 0x46, 0x58, 0x56, 0x32}
	if msg.Header.Site != expectedSiteID {
		t.Error("Site incorrect")
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

	if msg.AtTime != 0 {
		t.Error("atTime field incorrect")
	}

	if msg.Version != 1024 {
		t.Error("protocol version field incorrect")
	}

	if !msg.Addressable {
		t.Error("adressable field should be true")
	}

	expectedTargetID := [8]byte{0xd0, 0x73, 0xd5, 0x00, 0xf9, 0x14, 0x00, 0x00}
	if msg.Header.Target != expectedTargetID {
		t.Error("Target incorrect")
	}

	expectedSiteID := [6]byte{0x4c, 0x49, 0x46, 0x58, 0x56, 0x32}
	if msg.Header.Site != expectedSiteID {
		t.Error("Site incorrect")
	}

	payload := msg.Payload.(*payloads.DeviceStatePanGateway)

	if payload.Service != 2 {
		t.Error("Service incorrect")
	}

	if payload.Port != 0 {
		t.Error("Port incorrect")
	}
}

func TestDecodeLightState(t *T) {
	// DATA: length=88
	//       000  58 00 00 54 00 00 00 00  d0 73 d5 00 f9 14 00 00  |X..T.....s......|
	//       010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
	//       020  6b 00 00 00 00 00 00 00  7a 3e c4 09 00 00 00 00  |k.......z>......|
	//       030  42 65 64 72 6f 6f 6d 00  00 00 00 00 00 00 00 00  |Bedroom.........|
	//       040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
	//       050  01 00 00 00 00 00 00 00                           |........|
	// MSG:  &{version:1024 target:[208 115 213 0 249 20 0 0] site:[76 73 70 88 86 50] atTime:0 addressable:true tagged:false acknowledge:false}
	//       *payloads.LightState &{Color:{Hue:0 Saturation:0 Brightness:15994 Kelvin:2500} Dim:0 Power:0 Label:Bedroom Tags:1}

	b := []byte{
		0x58, 0x00, 0x00, 0x54, 0x00, 0x00, 0x00, 0x00,
		0xd0, 0x73, 0xd5, 0x00, 0xf9, 0x14, 0x00, 0x00,

		0x4c, 0x49, 0x46, 0x58, 0x56, 0x32, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x6b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x7a, 0x3e, 0xc4, 0x09, 0x00, 0x00, 0x00, 0x00,

		0x42, 0x65, 0x64, 0x72, 0x6f, 0x6f, 0x6d, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	msg, err := Decode(b)

	if err != nil {
		t.Error("Decode failed with err: " + err.Error())
	}

	if msg.AtTime != 0 {
		t.Error("atTime field incorrect")
	}

	if msg.Version != 1024 {
		t.Error("protocol version field incorrect")
	}

	if !msg.Addressable {
		t.Error("adressable field should be true")
	}

	expectedTargetID := [8]byte{0xd0, 0x73, 0xd5, 0x00, 0xf9, 0x14, 0x00, 0x00}
	if msg.Header.Target != expectedTargetID {
		t.Error("Target incorrect")
	}

	expectedSiteID := [6]byte{0x4c, 0x49, 0x46, 0x58, 0x56, 0x32}
	if msg.Header.Site != expectedSiteID {
		t.Error("Site incorrect")
	}

	payload := msg.Payload.(*payloads.LightState)
	// *payloads.LightState &{Color:{Hue:0 Saturation:0 Brightness:15994 Kelvin:2500} Dim:0 Power:0 Label:Bedroom Tags:1}

	if payload.Color.Hue != 0 {
		t.Error("Hue incorrect")
	}

	if payload.Color.Saturation != 0 {
		t.Error("Saturation incorrect")
	}

	if payload.Color.Brightness != 15994 {
		t.Error("Brightness incorrect")
	}

	if payload.Color.Kelvin != 2500 {
		t.Error("Brightness incorrect")
	}

	if payload.Dim != 0 {
		t.Error("Brightness incorrect")
	}

	if payload.Power != 0 {
		t.Error("Power incorrect")
	}

	if payload.Label.String() != "Bedroom" {
		t.Error("Label incorrect")
	}

	if payload.Tags != 1 {
		t.Error("Tags incorrect")
	}

}

func TestDeviceSetPowerOn(t *T) {
	// DATA: length=38
	//       000  26 00 00 34 00 00 00 00  00 00 00 00 00 00 00 00  |&..4............|
	//       010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
	//       020  15 00 00 00 01 00                                 |......|
	// MSG:  &{Version:1024 Target:[0 0 0 0 0 0 0 0] Site:[76 73 70 88 86 50] AtTime:0 Addressable:true Tagged:true Acknowledge:false}
	//       *payloads.DeviceSetPower &{Level:1}

	b := []byte{
		0x26, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00,
		0xd0, 0x73, 0xd5, 0x00, 0xf9, 0x14, 0x00, 0x00,

		0x4c, 0x49, 0x46, 0x58, 0x56, 0x32, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x15, 0x00, 0x00, 0x00, 0x01, 0x00,
	}

	msg, err := Decode(b)

	if err != nil {
		t.Error("Decode failed with err: " + err.Error())
	}

	payload := msg.Payload.(*payloads.DeviceSetPower)

	if payload.Level != 1 {
		t.Error("Level incorrect")
	}
}

func TestDeviceSetPowerOff(t *T) {
	// DATA: length=38
	//       000  26 00 00 54 00 00 00 00  d0 73 d5 00 f9 14 00 00  |&..T.....s......|
	//       010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
	//       020  15 00 00 00 00 00                                 |......|
	// MSG:  &{version:1024 target:[208 115 213 0 249 20 0 0] site:[76 73 70 88 86 50] atTime:0 addressable:true tagged:false acknowledge:false}
	//       *payloads.DeviceSetPower &{Level:0}

	b := []byte{
		0x26, 0x00, 0x00, 0x54, 0x00, 0x00, 0x00, 0x00,
		0xd0, 0x73, 0xd5, 0x00, 0xf9, 0x14, 0x00, 0x00,

		0x4c, 0x49, 0x46, 0x58, 0x56, 0x32, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x15, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	msg, err := Decode(b)

	if err != nil {
		t.Error("Decode failed with err: " + err.Error())
	}

	payload := msg.Payload.(*payloads.DeviceSetPower)

	if payload.Level != 0 {
		t.Error("Level incorrect")
	}
}

func TestDecodeDeviceStatePower(t *T) {
	// DATA: length=38
	//       000  26 00 00 54 00 00 00 00  d0 73 d5 00 f9 14 00 00  |&..T.....s......|
	//       010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
	//       020  16 00 00 00 00 00                                 |......|
	// MSG:  &{version:1024 target:[208 115 213 0 249 20 0 0] site:[76 73 70 88 86 50] atTime:0 addressable:true tagged:false acknowledge:false}
	//       *payloads.DeviceStatePower &{Level:65535}

	b := []byte{
		0x26, 0x00, 0x00, 0x54, 0x00, 0x00, 0x00, 0x00,
		0xd0, 0x73, 0xd5, 0x00, 0xf9, 0x14, 0x00, 0x00,

		0x4c, 0x49, 0x46, 0x58, 0x56, 0x32, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		0x16, 0x00, 0x00, 0x00, 0xff, 0xff,
	}

	msg, err := Decode(b)

	if err != nil {
		t.Error("Decode failed with err: " + err.Error())
	}

	if msg.AtTime != 0 {
		t.Error("atTime field incorrect")
	}

	if msg.Version != 1024 {
		t.Error("protocol version field incorrect")
	}

	if !msg.Addressable {
		t.Error("adressable field should be true")
	}

	expectedTargetID := [8]byte{0xd0, 0x73, 0xd5, 0x00, 0xf9, 0x14, 0x00, 0x00}
	if msg.Header.Target != expectedTargetID {
		t.Error("Target incorrect")
	}

	expectedSiteID := [6]byte{0x4c, 0x49, 0x46, 0x58, 0x56, 0x32}
	if msg.Header.Site != expectedSiteID {
		t.Error("Site incorrect")
	}

	payload := msg.Payload.(*payloads.DeviceStatePower)

	if payload.Level != 65535 {
		t.Error("Power level incorrect")
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

	if msg.AtTime != 0 {
		t.Error("atTime field incorrect")
	}

	if msg.Version != 1024 {
		t.Error("protocol version field incorrect")
	}

	if !msg.Addressable {
		t.Error("adressable field should be true")
	}
}

func TestMarshalBinary(t *T) {
	msg := Message{}
	target := [8]byte{0xd0, 0x73, 0xd5, 0x00, 0x49, 0x14, 0x00, 0x00}
	site := [6]byte{0x4c, 0x49, 0x46, 0x58, 0x56, 0x32}
	header := Header{
		Version: 1024,

		Target:      target,
		Site:        site,
		AtTime:      0,
		Addressable: false,
		Tagged:      false,
		Acknowledge: false,
	}
	msg.Header = &header
	msg.Payload = payloads.DeviceGetPanGateway{}

	data, err := msg.MarshalBinary()
	if err != nil {
		t.Error(err)
	}

	// t.Log("\n" + hex.Dump(data))
	// It should be this:
	// DATA: length=36
	// 00000000  24 00 00 04 00 00 00 00  d0 73 d5 00 49 14 00 00  |$........s..I...|
	// 00000010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
	// 00000020  02 00 00 00                                       |....|

	if !bytes.Equal(data[0:2], []byte{0x24, 0x00}) {
		t.Error("Size incorrect")
	}

	if !bytes.Equal(data[3:4], []byte{0x04}) {
		t.Error("Incorrect bitfield combo of version, addressable and tagged.")
		t.Error(fmt.Sprintf("Expected 0x04, but got 0x%s", hex.EncodeToString(data[3:4])))
	}

	if !bytes.Equal(data[8:16], target[0:8]) {
		t.Error("Target incorrect")
	}

	if !bytes.Equal(data[16:22], site[0:6]) {
		t.Error("Site incorrect")
	}

	if !bytes.Equal(data[32:34], []byte{0x02, 0x00}) {
		t.Error("Message type incorrect")
	}
}
