package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/bjeanes/go-lifx/protocol/payloads"
	"strings"
)

const compatibleVersion = 1024

type message struct {
	target  [8]byte
	site    [6]byte
	atTime  uint64
	payload payloads.Payload
}

func (msg message) String() string {
	t := "<unknown>"
	if msg.payload != nil {
		t = strings.Replace(fmt.Sprintf("%T", msg.payload), "*protocol.", "", 1)
	}

	return strings.Replace(fmt.Sprintf(
		"LIFXMessage(%s)%+v",
		t,
		msg.payload,
	), "&{", "{", -1)
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

	payload := payloads.ForId(msgHeader.Type)
	if payload != nil {
		if reader.Len() != binary.Size(payload) {
			return message{}, errors.New(fmt.Sprintf("Unexpected payload size for %T", payload))
		}

		binary.Read(reader, binary.LittleEndian, payload)
	} else {
		return message{}, errors.New("Unknown message")
	}

	if err != nil {
		return message{}, err
	}

	msg := message{
		atTime:  msgHeader.AtTime,
		target:  msgHeader.Target,
		site:    msgHeader.Site,
		payload: payload,
		// addressable: 0x1000&uint16(msgHeader.Bitfield1) > 0, // next bit
		// tagged:      0x2000&uint16(msgHeader.Bitfield1) > 0, // next bit
		// acknowledge: 0x1&uint16(msgHeader.Bitfield2) > 0, // top bit
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
