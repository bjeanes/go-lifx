package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

const compatibleVersion = 1024

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
	msgHeader := header{}
	err := binary.Read(reader, binary.LittleEndian, &msgHeader)

	if msgHeader.Size != uint16(len(b)) {
		// TODO: figure out if this is actually a problem or just ignoreable padding/noise at the end of the datagram?
		return message{}, errors.New(fmt.Sprintf("Incorrect message size (data: %d, header: %d)", len(b), msgHeader.Size))
	}

	if v := msgHeader.version(); v != compatibleVersion {
		return message{}, errors.New(fmt.Sprintf("Unknown message version (%d)", v))
	}

	payload := payloads.New(msgHeader.Type)
	if payload != nil {
		if reader.Len() != binary.Size(payload) {
			return message{}, errors.New(fmt.Sprintf("Unexpected payload size for %T", payload))
		}

		binary.Read(reader, binary.LittleEndian, payload)
	} else {
		payload = &struct {
			UnrecognizedMessage uint16
			rawBytes            []byte
		}{msgHeader.Type, b}
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

func (e BadDatagram) Error() string {
	return fmt.Sprintf("Error (%s) decoding datagram: %+v", e.err.Error(), e.Datagram)
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
