package animation

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ErrAnimationNotFound = errors.New("animation not found")

type Animated interface {
	SetAnimation(name string) error
}

type AnimationPlayer struct {
	currentAnimation *Animation
	animations       map[string]*Animation
}

func NewAnimationPlayer(
	currentAnimation *Animation,
	animations map[string]*Animation,
) *AnimationPlayer {
	return &AnimationPlayer{
		currentAnimation: currentAnimation,
		animations:       animations,
	}
}

func (ap *AnimationPlayer) Update(delta float32) {
	if ap.currentAnimation != nil {
		ap.currentAnimation.update(delta)
	}
}

func (ap *AnimationPlayer) Render(position rl.Vector2) {
	if ap.currentAnimation != nil {
		rl.DrawTextureRec(
			*ap.currentAnimation.frames[ap.currentAnimation.currentFrame].texture,
			ap.currentAnimation.frames[ap.currentAnimation.currentFrame].src,
			position,
			rl.White,
		)
	}
}

func (ap *AnimationPlayer) SetAnimation(name string) error {
	animation, ok := ap.animations[name]
	if !ok {
		return ErrAnimationNotFound
	}

	ap.currentAnimation = animation

	return nil
}
