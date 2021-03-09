package main

import (
	"fmt"
	"golang-design-pattern/06_builder/sample"
)

func main() {

	commonHouse := sample.NewCommonHouse()
	director := sample.NewDirector(commonHouse)
	director.Construct()
	fmt.Println(director.Construct())

	highBuildingHouse := sample.NewHighBuildingHouse()
	director = sample.NewDirector(highBuildingHouse)
	director.Construct()
	fmt.Println(director.Construct())

}
