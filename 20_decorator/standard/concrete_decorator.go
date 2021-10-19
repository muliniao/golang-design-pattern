package standard

import "fmt"

/**
装饰者(具体)
*/
type ConcreteDecorator struct {
	baseDecorator	BaseDecorator
}

func NewConcreteDecorator(baseDecorator BaseDecorator) *ConcreteDecorator {
	return &ConcreteDecorator{
		baseDecorator: baseDecorator,
	}
}

func (concreteDecorator *ConcreteDecorator) Operation() {
	concreteDecorator.addFunction()
	concreteDecorator.baseDecorator.Operation()
}

func (concreteDecorator *ConcreteDecorator) addFunction() {
	fmt.Println("为具体[被装饰者角色]增加额外的功能addedFunction()")
}
