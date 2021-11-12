package standard001

import "fmt"

/**
StateMachine
*/
type StateMachine struct {
	currentState IState
}

func NewStateMachine(state IState) *StateMachine {
	return &StateMachine{
		currentState: state,
	}
}

func (stateMachine *StateMachine) Event001() {
	stateMachine.currentState.Event001()
}

func (stateMachine *StateMachine) Event002() {
	stateMachine.currentState.Event002()
}

func (stateMachine *StateMachine) Event003() {
	stateMachine.currentState.Event003()
}

func (stateMachine *StateMachine) SetState(currentState IState) {
	stateMachine.currentState = currentState
}

func (stateMachine *StateMachine) GetCurrentState() string {
	fmt.Println("当前状态:" + stateMachine.currentState.GetCurrentState())
	return stateMachine.currentState.GetCurrentState()
}
