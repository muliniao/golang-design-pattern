package standard

import "fmt"

/**
	ConcreteState:具体状态，实现所有抽象接口State的方法
 */
type ConcreteState001 struct {
}

func (concreteState001 *ConcreteState001) Operation001(context *Context) {
	fmt.Println("concreteState001-operation001")
}

func (concreteState001 *ConcreteState001) Operation002(context *Context) {
	fmt.Println("concreteState001-operation002")
	/**
		切换状态到ConcreteState002
	 */
	context.SetState(new(ConcreteState002))
}

func (concreteState001 *ConcreteState001) GetCurrentState() string {
	return "ConcreteState001"
}
