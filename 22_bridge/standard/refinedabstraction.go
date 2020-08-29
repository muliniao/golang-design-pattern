package standard

import "fmt"

/**
	Refined	Abstraction:扩展抽象化角色
 */
type RefinedAbstraction struct {
	implementor Implementor
}

func NewRefinedAbstraction(implementor Implementor) *RefinedAbstraction {
	return &RefinedAbstraction{
		implementor: implementor,
	}
}

func (refinedAbstraction *RefinedAbstraction) Operation() {
	fmt.Println("具体实现化[Refined Abstraction]角色被调用")
	refinedAbstraction.implementor.OperationImp1()
}
