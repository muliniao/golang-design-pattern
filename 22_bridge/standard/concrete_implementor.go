package standard

import "fmt"

/**
Concrete Implementor:具体实现化
*/
type ConcreteImplementor struct {
}

func NewConcreteImplementor() *ConcreteImplementor {
	return &ConcreteImplementor{}
}

func (concreteImplementor *ConcreteImplementor) OperationImp1() {
	fmt.Println("具体实现化[Concrete Implementor]角色被调用")
}
