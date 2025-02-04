package character

import (
	"raylib-go-2d-topdown/src/animation"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CharacterIdleUpAnimation       = "player-idle-up-animation"
	CharacterIdleDownAnimation     = "player-idle-down-animation"
	CharacterIdleLeftAnimation     = "player-idle-left-animation"
	CharacterIdleRightAnimation    = "player-idle-right-animation"
	CharacterWalkingUpAnimation    = "player-walking-up-animation"
	CharacterWalkingDownAnimation  = "player-walking-down-animation"
	CharacterWalkingLeftAnimation  = "player-walking-left-animation"
	CharacterWalkingRightAnimation = "player-walking-right-animation"
)

type Textured interface {
	GetTexture() *rl.Texture2D
	GetTextureDimensions() rl.Vector2
}

func NewCharacterAnimationPlayer(player Textured) *animation.AnimationPlayer {
	vectorUp := rl.NewVector2(0, -1)
	vectorDown := rl.NewVector2(0, 1)
	vectorLeft := rl.NewVector2(-1, 0)
	vectorRight := rl.NewVector2(1, 0)

	animations := map[string]*animation.Animation{
		CharacterIdleUpAnimation:    createIdleAnimation(player, vectorUp),
		CharacterIdleDownAnimation:  createIdleAnimation(player, vectorDown),
		CharacterIdleLeftAnimation:  createIdleAnimation(player, vectorLeft),
		CharacterIdleRightAnimation: createIdleAnimation(player, vectorRight),

		CharacterWalkingUpAnimation:    createWalkingAnimation(player, vectorUp),
		CharacterWalkingDownAnimation:  createWalkingAnimation(player, vectorDown),
		CharacterWalkingLeftAnimation:  createWalkingAnimation(player, vectorLeft),
		CharacterWalkingRightAnimation: createWalkingAnimation(player, vectorRight),
	}

	return animation.NewAnimationPlayer(
		animations[CharacterIdleDownAnimation],
		animations,
	)
}

func createIdleAnimation(
	character Textured,
	facing rl.Vector2,
) *animation.Animation {
	var row float32 = 0
	if facing.Y != 0 {
		if facing.Y < 0 {
			row = 1
		}
	} else {
		if facing.X < 0 {
			row = 2
		} else {
			row = 3
		}
	}

	spriteWidth := character.GetTextureDimensions().X
	spriteHeight := character.GetTextureDimensions().Y

	frames := make([]animation.Frame, 2)
	for i := 0; i < 2; i++ {
		frames[i] = animation.NewFrame(
			character.GetTexture(),
			rl.Rectangle{
				X:      spriteWidth * float32(i),
				Y:      spriteHeight * float32(row),
				Width:  spriteWidth,
				Height: spriteHeight,
			},
			0.25,
		)
	}

	return animation.NewAnimation(frames)
}

func createWalkingAnimation(
	character Textured,
	facing rl.Vector2,
) *animation.Animation {
	var row float32 = 0
	if facing.Y != 0 {
		if facing.Y < 0 {
			row = 1
		}
	} else {
		if facing.X < 0 {
			row = 2
		} else {
			row = 3
		}
	}

	spriteWidth := character.GetTextureDimensions().X
	spriteHeight := character.GetTextureDimensions().Y

	frames := make([]animation.Frame, 2)
	for i := 0; i < 2; i++ {
		frames[i] = animation.NewFrame(
			character.GetTexture(),
			rl.Rectangle{
				X:      spriteWidth * float32(2+i),
				Y:      spriteHeight * float32(row),
				Width:  spriteWidth,
				Height: spriteHeight,
			},
			0.25,
		)
	}

	return animation.NewAnimation(frames)
}
