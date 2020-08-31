package main

import "golang-design-pattern/01_facade/standard"

func main() {

	facade := standard.NewFacade(standard.NewSubSystem001(), standard.NewSubSystem002())
	facade.Method()

}
