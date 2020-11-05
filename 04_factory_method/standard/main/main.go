package main

import "golang-design-pattern/04_factory_method/standard"

func main() {

	var abstractFactory standard.AbstractFactory = new(standard.ConcreteFactory001)
	abstractFactory.NewProduct().Show()

	var abstractFactory1 standard.AbstractFactory = new(standard.ConcreteFactory002)
	abstractFactory1.NewProduct().Show()

}
