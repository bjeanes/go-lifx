package client

import (
	proto "github.com/bjeanes/go-lifx/protocol"
	. "testing"
)

type StubConnection chan proto.Message

func (c StubConnection) Listen() (<-chan proto.Message, <-chan error) {
	return c, make(chan error)
}

func (c StubConnection) WriteMessage(msg proto.Message) error {
	// NO OP
	return nil
}

func NewStubClient() *client {
	c := new(client)
	c.init(func() (connection, error) { return make(StubConnection), nil })
	return c
}

func TestPrimarySubscription(t *T) {
	client := NewStubClient()
	conn := client.connection.(StubConnection)

	msg1 := proto.Message{Header: &proto.Header{Target: [8]byte{0, 0, 0, 0, 0, 0, 0, 1}}}
	msg2 := proto.Message{Header: &proto.Header{Target: [8]byte{0, 0, 0, 0, 0, 0, 0, 2}}}

	go func() {
		conn <- msg1
		conn <- msg2
	}()

	if <-client.Messages != msg1 {
		t.Error("Message did not make it through the primary subscription")
	}

	if <-client.Messages != msg2 {
		t.Error("Message did not make it through the primary subscription")
	}
}

func TestSubsequentSubscribes(t *T) {
	client := NewStubClient()
	conn := client.connection.(StubConnection)

	go func() {
		for {
			// Drain the main messages to prevent deadlock
			<-client.Messages
		}
	}()

	// Given a subscriber listening to all messages
	rcvd := client.Sub()

	// And there are 2 messages
	msg1 := proto.Message{Header: &proto.Header{Target: [8]byte{0, 0, 0, 0, 0, 0, 0, 1}}}
	msg2 := proto.Message{Header: &proto.Header{Target: [8]byte{0, 0, 0, 0, 0, 0, 0, 2}}}

	// When the client receives the messages
	go func() {
		conn <- msg1
		conn <- msg2
	}()

	// Then the subscriber should also receive them
	if msg1 != <-rcvd {
		t.Error("subscriber didn't receive the first message")
	}

	if msg2 != <-rcvd {
		t.Error("subscriber didn't receive the second message")
	}

}
