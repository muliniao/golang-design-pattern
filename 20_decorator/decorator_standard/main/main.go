package main

import (
	"fmt"
	"golang-design-pattern/20_decorator/decorator_standard"
)

func main() {

	var component001 decorator_standard.Component = decorator_standard.NewConcreteComponent()
	component001.Operation()

	fmt.Println("-----------------------------------------------------")

	var component002 decorator_standard.Component = decorator_standard.NewConcreteDecorator(decorator_standard.NewDecorator(component001))
	component002.Operation()
}
