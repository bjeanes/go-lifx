package client

import (
	proto "github.com/bjeanes/go-lifx/protocol"
	"time"
)

type connection interface {
	Listen() (<-chan proto.Message, <-chan error)
	WriteMessage(proto.Message) error
}

type client struct {
	connected  bool
	connection connection
	Messages   <-chan proto.Message
	Errors     <-chan error
	Lights     *lightCollection
}

func New() *client {
	c := &client{}
	if conn, err := proto.Connect(); err == nil {
		c.connection = conn
		c.connected = true
	}

	messages, errors := c.connection.Listen()

	c.Messages = messages
	c.Errors = errors
	c.Lights = &lightCollection{client: c}

	return c
}

func (c *client) SendMessage(payload proto.Payload) (data []byte, error error) {
	msg := proto.Message{}
	msg.Payload = payload

	c.connection.WriteMessage(msg)
	return data, nil
}

func (c *client) Discover() <-chan *light {
	ch := make(chan *light)

	go func() {
		timeout := time.After(5 * time.Second)

		c.SendMessage(proto.LightGet{})

		for {
			select {
			case <-timeout:
				close(ch)
			case msg := <-c.Messages:
				switch payload := msg.Payload.(type) {
				case *proto.DeviceStatePanGateway:
					// TODO: record gateway devices
				case *proto.LightState:
					ch <- c.Lights.Register(payload)
				default:
					// nada
				}
			}
		}
	}()

	return ch
}
