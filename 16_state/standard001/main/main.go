package main

import (
	"fmt"

	"golang-design-pattern/16_state/standard001"
)

func main() {

	stateMachine := standard001.NewStateMachine(new(standard001.ConcreteState001))
	fmt.Println(stateMachine.GetCurrentState())

	// concreteState001 ---> concreteState002
	state001 := standard001.NewConcreteState001(stateMachine)
	state001.Event002()
	fmt.Println(stateMachine.GetCurrentState())


	// concreteState002 ---> concreteState003
	state002 := standard001.NewConcreteState002(stateMachine)
	state002.Event003()
	fmt.Println(stateMachine.GetCurrentState())

}
