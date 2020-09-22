package main

import (
	"fmt"
	"golang-design-pattern/18_flyweight/standard"
)

func main ()  {

	flyWeight := standard.NewFlyWeightFactory()
	flyWeight.GetFlyWeight("weibo").Operation(standard.NewUnsharedConcreteFlyWeight("Sina"))
	flyWeight.GetFlyWeight("qq").Operation(standard.NewUnsharedConcreteFlyWeight("Tengxun"))

	flyWeight.GetFlyWeight("weibo").Operation(standard.NewUnsharedConcreteFlyWeight("Sina"))
	flyWeight.GetFlyWeight("qq").Operation(standard.NewUnsharedConcreteFlyWeight("Tengxun"))

	fmt.Println("")
	flyWeight.GetSize()

}
