package payloads

type lightState struct {
	Color lightHsbk
	Dim   int16
	Power uint16
	Label label
	Tags  uint64
}

func (lightState) Id() uint16 { return 107 }
