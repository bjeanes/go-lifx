package payloads

type deviceGetPower struct{}

func (deviceGetPower) Id() uint16 { return 20 }
