package decorator_standard

import "fmt"

/**
装饰者(具体)
*/
type ConcreteDecorator struct {
	decorator *Decorator
}

func NewConcreteDecorator(decorator *Decorator) *ConcreteDecorator {
	return &ConcreteDecorator{decorator: decorator}
}

func (concreteDecorator *ConcreteDecorator) Operation() {
	concreteDecorator.decorator.Operation()
	concreteDecorator.addFunction()
}

func (concreteDecorator *ConcreteDecorator) addFunction() {
	fmt.Println("为具体[被装饰者角色]增加额外的功能addedFunction()")
}
