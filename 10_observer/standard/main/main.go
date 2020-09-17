package main

import (
	"fmt"
	"golang-design-pattern/10_observer/standard"
)

func main()  {

	observerList := standard.NewConcreteSubject()

	concreteObserver001 := new(standard.ConcreteObserver001)
	concreteObserver002 := new(standard.ConcreteObserver002)

	observerList.AddObserver(concreteObserver001)
	observerList.AddObserver(concreteObserver002)

	fmt.Println("通知各个观察者")
	observerList.NotifyObserver()

}
