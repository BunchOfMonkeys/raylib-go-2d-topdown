package character

import (
	"raylib-go-2d-topdown/src/animation"
	character_animation "raylib-go-2d-topdown/src/animation/character"
	"raylib-go-2d-topdown/src/inputs"
	"raylib-go-2d-topdown/src/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StatefulAnimatedCharacter interface {
	state.Stateful
	animation.Animated
}

type IdleCharacterState struct {
	character StatefulAnimatedCharacter
	facing    *rl.Vector2
}

func NewIdleCharacterState(
	character StatefulAnimatedCharacter,
) *IdleCharacterState {
	return &IdleCharacterState{
		character: character,
		facing:    &rl.Vector2{X: 0, Y: 1},
	}
}

func (idle *IdleCharacterState) Enter() {
	var idleAnimation string

	if idle.facing.Y != 0 {
		if idle.facing.Y > 0 {
			idleAnimation = character_animation.CharacterIdleDownAnimation
		} else {
			idleAnimation = character_animation.CharacterIdleUpAnimation
		}
	} else {
		if idle.facing.X < 0 {
			idleAnimation = character_animation.CharacterIdleLeftAnimation
		} else {
			idleAnimation = character_animation.CharacterIdleRightAnimation
		}
	}

	idle.character.SetAnimation(idleAnimation)
}

func (idle *IdleCharacterState) HandleInputs(inputs *inputs.Inputs) state.State {
	if inputs.PressedDown != inputs.PressedUp ||
		inputs.PressedLeft != inputs.PressedRight {
		return idle.character.GetState(CharacterStateWalking)
	}

	return nil
}

func (idle *IdleCharacterState) Update() {}

func (idle *IdleCharacterState) Exit() {
	idle.facing.X, idle.facing.Y = 0, 0
}
