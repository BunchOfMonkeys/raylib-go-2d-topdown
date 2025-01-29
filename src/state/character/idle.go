package character

import (
	"raylib-go-2d-topdown/src/inputs"
	"raylib-go-2d-topdown/src/state"
)

type IdleCharacterState struct {
	character state.Stateful
}

func NewIdleCharacterState(
	character state.Stateful,
) *IdleCharacterState {
	return &IdleCharacterState{
		character: character,
	}
}

func (idle *IdleCharacterState) Enter() {}

func (idle *IdleCharacterState) HandleInputs(inputs *inputs.Inputs) state.State {
	if inputs.PressedDown != inputs.PressedUp ||
		inputs.PressedLeft != inputs.PressedRight {
		return idle.character.GetState(CharacterStateWalking)
	}

	return nil
}

func (idle *IdleCharacterState) Update() {}

func (idle *IdleCharacterState) Exit() {}
