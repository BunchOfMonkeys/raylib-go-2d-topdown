package character

import (
	"raylib-go-2d-topdown/src/animation"
	character_animation "raylib-go-2d-topdown/src/animation/character"
	"raylib-go-2d-topdown/src/inputs"
	"raylib-go-2d-topdown/src/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StatefulAnimatedAndMovable interface {
	state.Stateful
	animation.Animated
	state.Movable
}

type WalkingCharacterState struct {
	character StatefulAnimatedAndMovable
	direction *rl.Vector2
}

func NewWalkingCharacterState(character StatefulAnimatedAndMovable) *WalkingCharacterState {
	return &WalkingCharacterState{
		character: character,
		direction: &rl.Vector2{X: 0, Y: 0},
	}
}

func (walking *WalkingCharacterState) Enter() {}

func (walking *WalkingCharacterState) HandleInputs(inputs *inputs.Inputs) state.State {
	var X, Y float32 = 0, 0
	if inputs.PressedRight {
		X++
	}
	if inputs.PressedLeft {
		X--
	}
	if inputs.PressedUp {
		Y--
	}
	if inputs.PressedDown {
		Y++
	}

	if X == 0 && Y == 0 {
		state := walking.character.GetStates()[CharacterStateIdle]
		switch idleState := state.(type) {
		case *IdleCharacterState:
			//set idle facing to last movement direction
			idleState.facing.X = walking.direction.X
			idleState.facing.Y = walking.direction.Y
		}

		return state
	}

	if X != walking.direction.X || Y != walking.direction.Y {
		//update animation
		var walkingAnimation string

		if Y != 0 {
			if Y > 0 {
				walkingAnimation = character_animation.CharacterWalkingDownAnimation
			} else {
				walkingAnimation = character_animation.CharacterWalkingUpAnimation
			}
		} else {
			if X > 0 {
				walkingAnimation = character_animation.CharacterWalkingRightAnimation
			} else {
				walkingAnimation = character_animation.CharacterWalkingLeftAnimation
			}
		}

		walking.character.SetAnimation(walkingAnimation)
	}

	walking.direction.X = X
	walking.direction.Y = Y

	return nil
}

func (walking *WalkingCharacterState) Update() {
	walking.character.Move(*walking.direction)
}

func (walking *WalkingCharacterState) Exit() {
	walking.direction.X = 0
	walking.direction.Y = 0
}
