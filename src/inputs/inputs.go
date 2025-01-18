package inputs

type Inputs struct {
	PressedUp    bool
	PressedDown  bool
	PressedLeft  bool
	PressedRight bool
}

func New() *Inputs {
	return &Inputs{
		PressedUp:    false,
		PressedDown:  false,
		PressedLeft:  false,
		PressedRight: false,
	}
}
