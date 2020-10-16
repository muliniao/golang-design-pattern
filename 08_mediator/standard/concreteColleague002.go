package standard

import "fmt"

type ConcreteColleague002 struct {
	AbstractColleague
}

func NewConcreteColleague002() ConcreteColleague002 {
	return ConcreteColleague002{}
}

func (concreteColleague02 *ConcreteColleague002) Receive() {
	fmt.Println("具体同事类002收到请求")
}

func (concreteColleague002 *ConcreteColleague002) Send() {
	fmt.Println("具体同事类002发出请求")
	concreteColleague002.mediator.Relay(concreteColleague002.AbstractColleague)
}


