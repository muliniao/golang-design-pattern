package main

import (
	"fmt"

	"golang-design-pattern/06_builder/sample001"
)

func main() {

	commonHouse := sample001.NewCommonHouse()
	director := sample001.NewDirector(commonHouse)
	director.Construct()
	fmt.Println(director.Construct())

	highBuildingHouse := sample001.NewHighBuildingHouse()
	director = sample001.NewDirector(highBuildingHouse)
	director.Construct()
	fmt.Println(director.Construct())

}
