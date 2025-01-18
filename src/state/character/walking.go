package character

import (
	"state-machine/src/inputs"
	"state-machine/src/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StatefulAndMovable interface {
	state.Stateful
	state.Movable
}

type WalkingCharacterState struct {
	character StatefulAndMovable
	direction *rl.Vector2
}

func NewWalkingCharacterState(character StatefulAndMovable) *WalkingCharacterState {
	return &WalkingCharacterState{
		character: character,
		direction: &rl.Vector2{X: 0, Y: 0},
	}
}

func (walking *WalkingCharacterState) Enter() {}

func (walking *WalkingCharacterState) HandleInputs(inputs *inputs.Inputs) state.State {
	if inputs.PressedDown == inputs.PressedUp &&
		inputs.PressedLeft == inputs.PressedRight {
		return walking.character.GetState(CharacterStateIdle)
	}

	var updatedX, updatedY float32 = 0, 0
	
	if inputs.PressedDown {
		updatedY = 1
	} else if inputs.PressedUp {
		updatedY = -1
	}

	if inputs.PressedLeft {
		updatedX = -1
	} else if inputs.PressedRight {
		updatedX = 1
	}

	walking.direction.X = updatedX
	walking.direction.Y = updatedY

	return nil
}

func (walking *WalkingCharacterState) Update() {
	walking.character.Move(*walking.direction)
}

func (walking *WalkingCharacterState) Exit() {
	walking.direction.X = 0
	walking.direction.Y = 0
}
