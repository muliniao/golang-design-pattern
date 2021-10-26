package main

import (
	"fmt"

	"golang-design-pattern/12_Iterator/standard"
)

func main() {

	collection := standard.NewConcreteCollection()
	collection.List = append(collection.List, "aaa", "bbb", "ccc", "ddd", "eee")

	iterator := collection.CreateIterator()

	fmt.Println(iterator.CurrentItem())
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}

}
