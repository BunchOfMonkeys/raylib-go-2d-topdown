package main

import (
	"fmt"
	"state-machine/src/character"
	"state-machine/src/inputs"

	character_state "state-machine/src/state/character"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	bgColor           = rl.SkyBlue
	playerInputs      = inputs.New()
	characterInstance *character.Character
)

func main() {
	//setup once
	setup()

	//check if we need to exit (pressing ESC by default)
	for !rl.WindowShouldClose() {
		//get input, update and render stuff
		getInputs()
		update()
		render()
	}

	//finally cleanup and quit
	cleanup()
}

func setup() {
	rl.InitWindow(
		1000, 500,
		"State Machine Demo",
	)
	rl.SetTargetFPS(60)

	//make a new instance, keep all the setup inside the New method
	characterInstance = character.New(rl.NewVector2(500, 250))
}

func getInputs() {
	playerInputs.PressedUp = rl.IsKeyDown(rl.KeyW)
	playerInputs.PressedDown = rl.IsKeyDown(rl.KeyS)
	playerInputs.PressedLeft = rl.IsKeyDown(rl.KeyA)
	playerInputs.PressedRight = rl.IsKeyDown(rl.KeyD)

	characterInstance.HandleInputs(playerInputs)
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bgColor)

	characterInstance.Render()

	up, down, left, right := " ", " ", " ", " "
	if playerInputs.PressedUp {
		up = "U"
	}
	if playerInputs.PressedDown {
		down = "D"
	}
	if playerInputs.PressedLeft {
		left = "L"
	}
	if playerInputs.PressedRight {
		right = "R"
	}

	rl.DrawText(
		fmt.Sprintf("inputs: %s %s %s %s", up, down, left, right),
		10, 10,
		32,
		rl.Black,
	)
	state := ""
	switch characterInstance.GetCurrentState().(type) {
	case *character_state.IdleCharacterState:
		state = "idle"
	case *character_state.WalkingCharacterState:
		state = "walking"
	default:
		state = "unknown"
	}
	rl.DrawText(
		fmt.Sprintf("state: %s", state),
		10, 50,
		32,
		rl.Black,
	)

	rl.EndMode2D()
	rl.EndDrawing()
}

func update() {
	characterInstance.Update()
}

func cleanup() {
	rl.CloseWindow()
}
