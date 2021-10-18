package main

import (
	"fmt"

	"golang-design-pattern/06_builder/sample002"
)

func main () {
	resourceBuilder := sample002.NewResourceBuilder()
	director := sample002.NewDirector(resourceBuilder)
	fmt.Println(director.Construct())
}
