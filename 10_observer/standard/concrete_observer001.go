package standard

import "fmt"

/**
ConcreteObserver001: 具体观察者
*/
type ConcreteObserver001 struct {
}

func (concreteObserver *ConcreteObserver001) Response() {
	fmt.Println("具体观察者001做出反应")
}
