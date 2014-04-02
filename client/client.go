package client

import (
	proto "github.com/bjeanes/go-lifx/protocol"
)

type client struct {
	connected  bool
	connection *proto.Connection
}

func New() *client {
	c := &client{}
	go func() {
		if conn, err := proto.Connect(); err == nil {
			c.connection = conn
			c.connected = true
		}
	}()

	return c
}
