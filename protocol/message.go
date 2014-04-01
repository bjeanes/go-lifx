package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
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
		Type   uint16

		_ uint16 // <reserved>
	}

	payload interface{}

	message struct {
		size        uint16
		version     uint16
		addressable bool
		tagged      bool
		target      [8]byte
		site        [6]byte
		acknowledge bool
		atTime      uint64

		payload
	}
)

func (msg message) String() string {
	t := "<unknown>"
	if msg.payload != nil {
		t = fmt.Sprintf("%T", msg.payload)
	}

	return fmt.Sprintf(
		"LIFXMessage{version: %d, tagged: %v, type: %s, payload: %+v}",
		msg.version,
		msg.tagged,
		t,
		msg.payload,
	)
}

func Decode(b []byte) (message, error) {
	reader := bytes.NewReader(b)
	msgHeader := rawMessageHeader{}
	err := binary.Read(reader, binary.LittleEndian, &msgHeader)

	if msgHeader.Size != uint16(len(b)) {
		return message{}, errors.New(fmt.Sprintf("Incorrect message size (data: %d, header: %d)", len(b), msgHeader.Size))
	}

	payload := initPayload(msgHeader.Type)
	if payload != nil {
		binary.Read(reader, binary.LittleEndian, payload)
	} else {
		payload = &struct{ rawBytes []byte }{b}
	}

	if err != nil {
		return message{}, err
	}

	msg := message{
		size:        msgHeader.Size,
		version:     0xfff & uint16(msgHeader.Bitfield1),    // top 12 bits
		addressable: 0x1000&uint16(msgHeader.Bitfield1) > 0, // next bit
		tagged:      0x2000&uint16(msgHeader.Bitfield1) > 0, // next bit
		target:      msgHeader.Target,
		site:        msgHeader.Site,
		acknowledge: 0x1&uint16(msgHeader.Bitfield2) > 0, // top bit
		atTime:      msgHeader.AtTime,
		payload:     payload,
	}

	return msg, nil
}

type BadDatagram struct {
	Datagram datagram
	err      error
}

func (err BadDatagram) Error() string {
	return err.err.Error()
}

type errChan chan error

func (ch errChan) send(err error) {
	select {
	case ch <- err:
		// Error sent
	default:
		// Drop error if the channel is blocked.
	}
}

func NewMessageDecoder(datagrams <-chan datagram) (<-chan message, errChan) {
	msgs := make(chan message, 1)
	errs := make(errChan, 1)

	go func() {
		for datagram := range datagrams {
			msg, err := Decode(datagram.Data)

			if err != nil {
				errs.send(&BadDatagram{datagram, err})
				continue
			}

			msgs <- msg
		}

		close(msgs)
	}()

	return msgs, errs
}

func debug(str string, vals ...interface{}) {
	fmt.Printf(str+"\n", vals...)
}
