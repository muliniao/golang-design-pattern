package main

import (
	simpleFactory "golang-design-pattern/00_simple_factory"
)

func main() {

	factory := simpleFactory.NewFactory("Volvo")
	factory.Operation()
}
