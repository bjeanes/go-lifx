package client

import (
	"github.com/bjeanes/go-lifx/protocol/payloads"
)

type light struct {
	*client
}

func (l light) TurnOff() {
	l.client.SendMessage(payloads.DeviceSetPower{Level: 0})
}

func (l light) TurnOn() {
	l.client.SendMessage(payloads.DeviceSetPower{Level: 1})
}
