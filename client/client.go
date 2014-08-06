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
	subs       subService
}

func New() *client {
	c := &client{}
	c.init(func() (connection, error) {
		return proto.Connect()
	})

	return c
}

func (c *client) init(connector func() (connection, error)) {
	if conn, err := connector(); err == nil {
		c.connection = conn
		c.connected = true
	}

	messages, errors := c.connection.Listen()

	c.Lights = &lightCollection{client: c}
	c.subs = newSubService(messages)
	c.Messages = c.subs.Sub()
	c.Errors = errors
}

func (c *client) Sub() <-chan proto.Message {
	return c.subs.Sub()
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
