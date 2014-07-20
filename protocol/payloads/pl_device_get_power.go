package payloads

type deviceGetPower struct{}

func (deviceGetPower) Id() uint8 { return 20 }
