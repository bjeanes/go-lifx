package client

import (
	"fmt"
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

func (c *client) Discover() {
	c.SendMessage(payloads.LightGet{})

	ticker := time.NewTicker(time.Second)

	go func() {
		for _ = range ticker.C {
			select {
			case msg := <-c.Messages:
				switch payload := msg.Payload.(type) {
				case *payloads.DeviceStatePanGateway:
					// TODO: record gateway devices
				case *payloads.LightState:
					fmt.Printf("Discovered bulb %s\n", payload.Label)
					c.Lights.Register(payload)
				default:
					fmt.Printf("I heard something, and it was a %T\n", payload)
				}
			}
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
}
