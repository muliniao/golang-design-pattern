package main

import "golang-design-pattern/21_responsibility/standard"

func main() {

	handler001 := new(standard.ConcreteHandler001)
	handler002 := new(standard.ConcreteHandler002)

	handler001.AbstractHandler.SetNext(handler002)

	handler001.HandleRequest("two")

}
