package main

import (
	"fmt"
	"golang-design-pattern/06_builder/standard"
)

func main() {

	concreteBuilder := standard.NewConcreteBuilder()
	director := standard.NewDirector(concreteBuilder)
	concreteProduct := director.Construct()

	fmt.Println(*concreteProduct)

}
