package client

import (
	proto "github.com/bjeanes/go-lifx/protocol"
)

type lightCollection struct {
	*client
	lights []light
}

func (lc *lightCollection) Register(state *proto.LightState) *light {
	light := light{client: lc.client}
	light.UpdateFromState(state)
	lc.lights = append(lc.lights, light)
	return &light
}

func (lc *lightCollection) Count() int {
	return len(lc.lights)
}

func (lc *lightCollection) All() []light {
	return lc.lights
}

type light struct {
	*client
	state *proto.LightState
}

func (l *light) Label() string {
	return l.state.Label.String()
}

func (l *light) UpdateFromState(state *proto.LightState) {
	l.state = state
}

func (l light) TurnOff() {
	l.client.SendMessage(proto.DeviceSetPower{Level: 0})
}

func (l light) TurnOn() {
	l.client.SendMessage(proto.DeviceSetPower{Level: 1})
}
