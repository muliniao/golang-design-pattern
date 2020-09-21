package main

import (
	"fmt"
	"golang-design-pattern/12_Iterator/standard"
)

func main() {

	var aggregate standard.Aggregate = standard.NewConcreteAggregate()

	aggregate.Add("AAAAAA")
	aggregate.Add("BBBBBB")
	aggregate.Add("CCCCCC")

	fmt.Println("聚合内容有:")

	var itetator standard.Iterator = aggregate.GetIterator()

	fmt.Printf("print the first element: [%v]\n", itetator.First())

	fmt.Println("print all elements one by one")
	for itetator.HasNext() {
		fmt.Println(itetator.Next())
	}
}
