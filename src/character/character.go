package character

import (
	"state-machine/src/inputs"
	character_state "state-machine/src/state/character"

	//in next snippet
	"state-machine/src/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Character struct {
	state.StateMachine

	texture rl.Texture2D

	position      rl.Vector2
	movementSpeed float32
}

func New(
	position rl.Vector2,
) *Character {
	c := &Character{
		//this is bad, mkay? we'll address this later
		texture: rl.LoadTexture("./assets/Characters/BasicCharacterSpritesheet.png"),

		position:      position,
		movementSpeed: 3,
	}

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

func (c *Character) Update() {
	if c.GetCurrentState() == nil {
		return
	}

	c.GetCurrentState().Update()
}

// but the Rendering is done here
func (c *Character) Render() {
	rl.DrawTextureRec(
		c.texture,
		//sprite size is 48x48 px, we'll draw the first row, first column one from the spritesheet
		rl.NewRectangle(0, 0, 48, 48),
		//character's position is basically a dot in the center of the sprite,
		//so we need to draw it from position shifter by half of sprite both by X and Y
		rl.NewVector2(
			c.position.X-24,
			c.position.Y-24,
		),
		rl.White,
	)
}

func (c *Character) Move(direction rl.Vector2) {
	direction = rl.Vector2Normalize(direction)
	c.position.X += direction.X * c.movementSpeed
	c.position.Y += direction.Y * c.movementSpeed
}
