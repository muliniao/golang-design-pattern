package standard

import "fmt"

type ConcreteTemplate struct {
	TemplateFather
}

func (c *ConcreteTemplate) Fun1() {
	fmt.Println("Concrete template Func1()")
}

func (c *ConcreteTemplate) Fun2() {
	fmt.Println("Concrete template Func2()")
}
