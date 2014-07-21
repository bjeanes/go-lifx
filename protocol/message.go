package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/bjeanes/go-lifx/protocol/payloads"
	// "strings"
)

const compatibleVersion = 1024

type message struct {
	*Header
	payloads.Payload
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

		if err := binary.Read(reader, binary.LittleEndian, payload); err != nil {
			return message{}, err
		}
	} else {
		return message{}, fmt.Errorf("Unknown message (%d)", msgHeader.Type)
	}

	if err != nil {
		return message{}, err
	}

	msg := message{msgHeader.ToExpandedHeader(), payload}

	return msg, nil
}

type BadDatagram struct {
	Datagram Datagram
	err      error
}

func (e BadDatagram) Error() string {
	return e.err.Error()
}

func NewMessageDecoder(datagrams <-chan Datagram) (<-chan message, <-chan error) {
	msgs := make(chan message, 1)
	errs := make(chan error, 1)

	go func() {
		for datagram := range datagrams {
			msg, err := Decode(datagram.Data)

			if err != nil {
				errs <- &BadDatagram{datagram, err}
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
