package main

import "golang-design-pattern/15_strategy/standard"

func main()  {

	concreteStrategy := new(standard.ConcreteStrategy001)
	context := standard.NewContext(concreteStrategy)
	context.ConcreteStrategyMethod()

}
