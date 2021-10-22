package main

import "golang-design-pattern/21_responsibility/standard001"

func main() {

	handler001 := new(standard001.ConcreteHandler001)
	handler002 := new(standard001.ConcreteHandler002)

	handler001.AbstractHandler.SetNext(handler002)

	handler001.HandleRequest("two")

}
