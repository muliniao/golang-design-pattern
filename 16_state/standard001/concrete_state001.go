package standard001

import "fmt"

/**
ConcreteState:具体状态，实现所有抽象接口State的方法
*/
type ConcreteState001 struct {
	StateMachine *StateMachine
}

func NewConcreteState001(machine *StateMachine) *ConcreteState001 {
	return &ConcreteState001{
		StateMachine: machine,
	}
}

func (concreteState001 *ConcreteState001) Event001() {
	fmt.Println("concreteState001-Event001: do nothing")
}

func (concreteState001 *ConcreteState001) Event002() {
	fmt.Println("concreteState001-Event002")
	/**
	切换状态到ConcreteState002
	*/
	concreteState001.StateMachine.SetState(new(ConcreteState002))
}

func (concreteState001 *ConcreteState001) Event003() {
	fmt.Println("concreteState001-Event003")
	/**
	切换状态到ConcreteState003
	*/
	concreteState001.StateMachine.SetState(new(ConcreteState003))
}


func (concreteState001 *ConcreteState001) GetCurrentState() string {
	return "ConcreteState001"
}
