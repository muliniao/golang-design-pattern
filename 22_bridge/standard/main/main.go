package main

import (
	"golang-design-pattern/22_bridge/standard"
)

func main() {

	var abstraction standard.Abstraction = standard.NewRefinedAbstraction(standard.NewConcreteImplementor())
	abstraction.Operation()
}
