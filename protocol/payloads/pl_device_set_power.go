package payloads

type deviceSetPower struct {
	Level uint16 // 0 = off; 1+ = on
}

func (deviceSetPower) Id() uint8 { return 21 }
