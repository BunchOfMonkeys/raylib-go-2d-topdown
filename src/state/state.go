package state

import (
	"raylib-go-2d-topdown/src/inputs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Stateful interface {
	GetStates() map[string]State
	GetState(string) State
	GetCurrentState() State
	EnterState(State)
}

type Movable interface {
	Move(rl.Vector2)
}

type State interface {
	//setup
	Enter()
	//decide if current state needs to be changed
	HandleInputs(*inputs.Inputs) State
	//do the magic
	Update()
	//cleanup
	Exit()
}

type StateMachine struct {
	currentState    State
	availableStates map[string]State
}

func (sm *StateMachine) SetStates(states map[string]State) {
	sm.availableStates = states
}

func (sm *StateMachine) GetStates() map[string]State {
	return sm.availableStates
}

func (sm *StateMachine) GetState(name string) State {
	return sm.availableStates[name]
}

func (sm *StateMachine) GetCurrentState() State {
	return sm.currentState
}

func (sm *StateMachine) EnterState(state State) {
	if sm.currentState != nil {
		sm.currentState.Exit()
	}

	sm.currentState = state
	sm.currentState.Enter()
}
