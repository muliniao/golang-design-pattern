package standard

import "fmt"

/**
ConcreteObserver002: 具体观察者
*/
type ConcreteObserver002 struct {
}

func (concreteObserver *ConcreteObserver002) Response() {
	fmt.Println("具体观察者002做出反应")
}
