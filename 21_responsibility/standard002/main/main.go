package main

import "golang-design-pattern/21_responsibility/standard002"

func main() {

	chain := standard002.NewHandlerChain()
	chain.AddHandler(new(standard002.ConcreteHandler001))
	chain.AddHandler(new(standard002.ConcreteHandler002))
	chain.Handle()

}
