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
	Lights     []light
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

	return c
}

func (c *client) SendMessage(payload payloads.Payload) (data []byte, error error) {
	msg := proto.Message{}
	msg.Payload = payload

	c.connection.WriteMessage(msg)
	return data, nil
}

func (c *client) Discover() {
	c.SendMessage(payloads.DeviceGetPanGateway{})

	ticker := time.NewTicker(time.Second)

	go func() {
		for _ = range ticker.C {
			select {
			case msg := <-c.Messages:
				switch msg.Payload.(type) {
				case *payloads.DeviceStatePanGateway:
					c.Lights = append(c.Lights, light{client: *c})
				default:
					fmt.Println(fmt.Sprintf("I heard something, and it was a %T", msg.Payload))
				}
			}
		}
	}()

	time.Sleep(time.Second * 5)
	ticker.Stop()

	fmt.Println(fmt.Sprintf("Found %d lights", len(c.Lights)))

}
