package main

import (
	"fmt"
	"golang-design-pattern/16_state/standard"
)

func main() {

	context := standard.NewContext(new(standard.ConcreteState001))
	fmt.Println(context.GetCurrentState())

	// concreteState001 ---> concreteState002
	context.Operation002(context)
	fmt.Println(context.GetCurrentState())

	// concreteState002 ---> concreteState001
	context.Operation001(context)
	fmt.Println(context.GetCurrentState())

}
