package payloads

type deviceStatePower struct {
	Level uint16 // 0 = off; 1+ = on
}

func (deviceStatePower) Id() uint16 { return 21 }
