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

// http://golang.org/pkg/encoding/#BinaryUnmarshaler
func (msg *message) UnmarshalBinary(data []byte) error {
	reader := bytes.NewReader(data)
	msgHeader := header{}
	err := binary.Read(reader, binary.LittleEndian, &msgHeader)

	if msgHeader.Size != uint16(len(data)) {
		// TODO: figure out if this is actually a problem or just ignoreable padding/noise at the end of the datagram?
		return errors.New(fmt.Sprintf("Incorrect message size (data: %d, header: %d)", len(data), msgHeader.Size))
	}

	if v := msgHeader.version(); v != compatibleVersion {
		return errors.New(fmt.Sprintf("Unknown message version (%d)", v))
	}

	payload := payloads.ForId(msgHeader.Type)
	if payload != nil {
		if reader.Len() != binary.Size(payload) {
			return errors.New(fmt.Sprintf("Unexpected payload size for %T", payload))
		}

		if err := binary.Read(reader, binary.LittleEndian, payload); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Unknown message (%d)", msgHeader.Type)
	}

	if err != nil {
		return err
	}

	msg.Header = msgHeader.ToExpandedHeader()
	msg.Payload = payload

	return nil
}

// http://golang.org/pkg/encoding/#BinaryMarshaler
func (m *message) MarshalBinary() (data []byte, err error) {
	// TODO
	return
}

func Decode(b []byte) (message, error) {
	msg := new(message)
	err := msg.UnmarshalBinary(b)
	return *msg, err
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
