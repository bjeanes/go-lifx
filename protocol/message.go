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
const headerSize = 36

type Message struct {
	*Header
	payloads.Payload
}

// http://golang.org/pkg/encoding/#BinaryUnmarshaler
func (msg *Message) UnmarshalBinary(data []byte) error {
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
func (m *Message) MarshalBinary() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, m.Payload)
	if err != nil {
		return []byte{}, err
	}
	payloadBytes := buf.Bytes()

	raw := m.Header.ToRawHeader()
	raw.Size = headerSize + uint16(len(payloadBytes))
	raw.Type = m.Payload.Id()
	buf.Truncate(0)
	err = binary.Write(buf, binary.LittleEndian, raw)

	if err != nil {
		return []byte{}, err
	}

	return append(buf.Bytes(), payloadBytes...), nil
}

func Decode(b []byte) (Message, error) {
	msg := new(Message)
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

func NewMessageDecoder(datagrams <-chan Datagram) (<-chan Message, <-chan error) {
	msgs := make(chan Message, 1)
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
