package main

import "golang-design-pattern/05_abstract_factory/standard"

func main() {

	var abstractFactory standard.AbstractFactory = new(standard.ConcreteFactory001)
	abstractFactory.NewProduct001().Show()
	abstractFactory.NewProduct002().Show()

}
