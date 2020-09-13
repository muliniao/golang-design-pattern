package standard

import "fmt"

/**
	ConcreteState:具体状态，实现所有抽象接口State的方法
*/
type ConcreteState002 struct {

}

func (concreteState002 *ConcreteState002) Operation001(context *Context) {
	fmt.Println("concreteState002-operation001")
	/**
		切换状态到ConcreteState002
	*/
	context.SetState(new(ConcreteState001))
}


func (concreteState002 *ConcreteState002) Operation002(context *Context) {
	fmt.Println("concreteState002-operation001")
}

func (concreteState002 *ConcreteState002) GetCurrentState() string {
	return "ConcreteState002"
}
