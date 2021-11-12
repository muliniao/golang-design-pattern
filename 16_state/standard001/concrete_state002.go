package standard001

import "fmt"

/**
ConcreteState:具体状态，实现所有抽象接口State的方法
*/
type ConcreteState002 struct {
	StateMachine *StateMachine
}

func NewConcreteState002(machine *StateMachine) *ConcreteState002 {
	return &ConcreteState002{
		StateMachine: machine,
	}
}

func (concreteState002 *ConcreteState002) Event001() {
	fmt.Println("concreteState002-Event001")
	/**
	切换状态到ConcreteState002
	*/
	concreteState002.StateMachine.SetState(new(ConcreteState001))
}

func (concreteState002 *ConcreteState002) Event002() {
	fmt.Println("concreteState002-Event002: do nothing")
}

func (concreteState002 *ConcreteState002) Event003() {
	fmt.Println("concreteState002-Event003")
	/**
	切换状态到ConcreteState003
	*/
	concreteState002.StateMachine.SetState(new(ConcreteState003))
}


func (concreteState002 *ConcreteState002) GetCurrentState() string {
	return "ConcreteState002"
}
