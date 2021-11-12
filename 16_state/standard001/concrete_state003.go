package standard001

import "fmt"

/**
ConcreteState:具体状态，实现所有抽象接口State的方法
*/
type ConcreteState003 struct {
	StateMachine *StateMachine
}

func NewConcreteState003(machine *StateMachine) *ConcreteState003 {
	return &ConcreteState003{
		StateMachine: machine,
	}
}

func (concreteState003 *ConcreteState003) Event001() {
	fmt.Println("concreteState003-Event001")
	/**
	切换状态到ConcreteState001
	*/
	concreteState003.StateMachine.SetState(new(ConcreteState001))
}

func (concreteState003 *ConcreteState003) Event002() {
	fmt.Println("concreteState003-Event002")
	/**
	切换状态到ConcreteState002
	*/
	concreteState003.StateMachine.SetState(new(ConcreteState002))
}

func (concreteState003 *ConcreteState003) Event003() {
	fmt.Println("concreteState003-Event003: do nothing")
}

func (concreteState003 *ConcreteState003) GetCurrentState() string {
	return "ConcreteState003"
}
