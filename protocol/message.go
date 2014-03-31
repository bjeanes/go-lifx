package protocol

import (
  "encoding/binary"
  "fmt"
  "bytes"
)

type (
  bitfield uint16

  rawMessageHeader struct {
    Size uint16

    // 12 bits = protocol version
    // 1 bit   = addressable bool
    // 1 bit   = tagged bool
    // 2 bit   = <reserved>
    Bitfield1 bitfield

    _ uint32 // <reserved>

    Target [8]byte
    Site   [6]byte

    // 1 bit = acknowledge bool
    // 15 bits = <reserved>
    Bitfield2 bitfield

    AtTime uint64
    Type uint16

    _ uint16 // <reserved>
  }

  payload interface{}

  message struct {
    size uint16
    version uint16
    addressable bool
    tagged bool
    target [8]byte
    site   [6]byte
    acknowledge bool
    atTime uint64

    payload
  }
)


func (msg message) String() string {
  return fmt.Sprintf(
    "LIFXMessage{version: %d, type: %T, tagged: %v}",
    msg.version,
    msg.payload,
    msg.tagged,
  )
}


func Decode(b []byte) (message, error) {
  reader := bytes.NewReader(b)
  msgHeader := rawMessageHeader{}
  err := binary.Read(reader, binary.LittleEndian, &msgHeader)

  if err != nil {
    return message{}, err
  }

  msg := message{
    size        : msgHeader.Size,
    version     : 0xfff & uint16(msgHeader.Bitfield1),      // top 12 bits
    addressable : 0x1000 & uint16(msgHeader.Bitfield1) > 0, // next bit
    tagged      : 0x2000 & uint16(msgHeader.Bitfield1) > 0, // next bit
    target      : msgHeader.Target,
    site        : msgHeader.Site,
    acknowledge : 0x1 & uint16(msgHeader.Bitfield2) > 0, // top bit
    atTime      : msgHeader.AtTime,
  }

  return msg, nil
}

func NewMessageDecoder(datagrams <-chan datagram) <-chan message {
  msgs := make(chan message, 1)

  go func() {
    for {
      datagram := <-datagrams
      msg, err := Decode(datagram.Data)

      if err != nil {
        close(msgs)
        return
      }

      msgs <- msg
    }
  }()

  return msgs
}

func debug(str string, vals ...interface{}) {
  fmt.Printf(str + "\n", vals...)
}
