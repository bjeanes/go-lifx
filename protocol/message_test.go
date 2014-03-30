package protocol

import (
  . "testing"
  "bytes"
)

func TestDecodeMessage(t *T) {
  b := []byte{
      0x39, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x31, 0x6c, 0x69, 0x66, 0x78, 0x31, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x67, 0x00, 0x00, 0x00,
      0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xac, 0x0d, 0xc8, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x80, 0x3f, 0x00, 0x00, 0x00,
  }

  msgs := NewMessageDecoder(bytes.NewReader(b))

  msg := <-msgs

  if msg.size != 57 {
    t.Error("size field incorrect")
  }

  if msg.site != [6]byte{'1','l','i','f','x','1'} {
    t.Error("site field incorrect")
  }

  if msg.atTime != 0 {
    t.Error("atTime field incorrect")
  }

  if msg.msgType != 103 {
    t.Error("type field incorrect")
  }

  if msg.version != 1024 {
    t.Error("protocol version field incorrect")
  }

  if !msg.tagged {
    t.Error("tagged field should be true")
  }

  if !msg.addressable {
    t.Error("adressable field should be true")
  }

  target := [8]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}
  if msg.target != target {
    t.Error("target field incorrect")
  }
}