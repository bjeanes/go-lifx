package client

import (
	proto "github.com/bjeanes/go-lifx/protocol"
	payloads "github.com/bjeanes/go-lifx/protocol/payloads"
	"time"
)

type client struct {
	connected  bool
	connection *proto.Connection
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

func (c *client) SendMessage(payload payloads.Payload) (data []byte, error error) {
	msg := proto.Message{}
	msg.Payload = payload

	c.connection.WriteMessage(msg)
	return data, nil
}

func (c *client) Discover() <-chan *light {
	ch := make(chan *light)

	go func() {
		timeout := time.After(5 * time.Second)
		tick := time.Tick(400 * time.Millisecond)

		for {
			select {
			case <-timeout:
				close(ch)
			case <-tick:
				c.SendMessage(payloads.LightGet{})
			case msg := <-c.Messages:
				switch payload := msg.Payload.(type) {
				case *payloads.DeviceStatePanGateway:
					// TODO: record gateway devices
				case *payloads.LightState:
					// If we find a bulb, let's extend the timeout another second:
					timeout = time.After(1 * time.Second)
					ch <- c.Lights.Register(payload)
				default:
					// nada
				}
			}
		}
	}()

	return ch
}
