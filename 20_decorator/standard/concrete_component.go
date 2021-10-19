package standard

import "fmt"

/**
被装饰者(具体)
*/
type ConcreteComponent struct {

}

func NewConcreteComponent() *ConcreteComponent {
	return &ConcreteComponent{}
}

func (concreteComponent *ConcreteComponent) Operation() {
	fmt.Println("调用[被装饰者角色]的方法operation()")
}
