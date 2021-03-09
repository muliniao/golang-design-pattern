package main

import "golang-design-pattern/08_mediator/standard"

func main() {

	/**
	Not Complete
	*/
	var mediator standard.Mediator = standard.NewConcreteMediator()

	var c1 standard.ConcreteColleague001

	mediator.Register(c1.AbstractColleague)

}
