package client

import (
	"github.com/bjeanes/go-lifx/protocol/payloads"
)

type lightCollection struct {
	*client
	lights []light
}

func (lc *lightCollection) Register(state *payloads.LightState) {
	light := light{client: lc.client}
	light.UpdateFromState(state)
	lc.lights = append(lc.lights, light)
}

func (lc *lightCollection) Count() int {
	return len(lc.lights)
}

func (lc *lightCollection) All() []light {
	return lc.lights
}

type light struct {
	*client
	state *payloads.LightState
}

func (l *light) Label() string {
	return l.state.Label.String()
}

func (l *light) UpdateFromState(state *payloads.LightState) {
	l.state = state
}

func (l light) TurnOff() {
	l.client.SendMessage(payloads.DeviceSetPower{Level: 0})
}

func (l light) TurnOn() {
	l.client.SendMessage(payloads.DeviceSetPower{Level: 1})
}
