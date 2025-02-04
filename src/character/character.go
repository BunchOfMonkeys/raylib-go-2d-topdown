package character

import (
	"raylib-go-2d-topdown/src/animation"
	character_animation "raylib-go-2d-topdown/src/animation/character"
	"raylib-go-2d-topdown/src/inputs"
	character_state "raylib-go-2d-topdown/src/state/character"

	"raylib-go-2d-topdown/src/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	characterWidth      = 48
	characterHeight     = 48
	characterHalfWidth  = characterWidth / 2
	characterHalfHeight = characterHeight / 2
)

type Character struct {
	state.StateMachine

	texture rl.Texture2D

	position      rl.Vector2
	movementSpeed float32

	animationPlayer *animation.AnimationPlayer
}

func New(
	position rl.Vector2,
) *Character {
	c := &Character{
		texture: rl.LoadTexture("./assets/Characters/BasicCharacterSpritesheet.png"),

		position:      position,
		movementSpeed: 3,
	}

	c.animationPlayer = character_animation.NewCharacterAnimationPlayer(c)

	idleState := character_state.NewIdleCharacterState(c)
	walkingState := character_state.NewWalkingCharacterState(c)

	c.SetStates(map[string]state.State{
		character_state.CharacterStateIdle:    idleState,
		character_state.CharacterStateWalking: walkingState,
	})

	c.EnterState(idleState)

	return c
}

func (c *Character) HandleInputs(inputs *inputs.Inputs) {
	//not the changes from c.currentState to c.GetCurrentState(),
	//since c.currentState is private now
	if c.GetCurrentState() == nil {
		return
	}

	nextState := c.GetCurrentState().HandleInputs(inputs)
	if nextState != nil {
		c.EnterState(nextState)
	}
}

func (c *Character) Update(delta float32) {
	if c.GetCurrentState() == nil {
		return
	}

	c.GetCurrentState().Update()
	c.animationPlayer.Update(delta)
}

// now the animation player takes care of rendering
func (c *Character) Render() {
	c.animationPlayer.Render(
		rl.NewVector2(
			c.position.X-characterHalfWidth,
			c.position.Y-characterHalfHeight,
		),
	)
}

func (c *Character) Move(direction rl.Vector2) {
	direction = rl.Vector2Normalize(direction)
	c.position.X += direction.X * c.movementSpeed
	c.position.Y += direction.Y * c.movementSpeed
}

// implement Animated interface
func (c *Character) SetAnimation(name string) error {
	return c.animationPlayer.SetAnimation(name)
}

// implement Textured interface
func (c *Character) GetTexture() *rl.Texture2D {
	return &c.texture
}

func (c *Character) GetTextureDimensions() rl.Vector2 {
	return rl.NewVector2(characterWidth, characterHeight)
}
