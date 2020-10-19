package main

import (
	"fmt"
	"golang-design-pattern/23_visitor/standard"
)

func main()  {

	os := standard.NewObjectStructure()

	os.Attach(new(standard.ConcreteElementA))
	os.Attach(new(standard.ConcreteElementB))

	var visitor001 standard.Visitor = new(standard.ConcreteVisitorA)
	os.Display(visitor001)

	fmt.Println("-------------------------")

	var visitor002 standard.Visitor = new(standard.ConcreteVisitorB)
	os.Display(visitor002)
}
