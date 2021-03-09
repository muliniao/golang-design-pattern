package standard

import "fmt"

type ConcreteColleague001 struct {
	AbstractColleague AbstractColleague
}

func NewConcreteColleague001() ConcreteColleague001 {
	return ConcreteColleague001{}
}

func (concreteColleague001 *ConcreteColleague001) Receive() {
	fmt.Println("具体同事类001收到请求")
}

func (concreteColleague001 *ConcreteColleague001) Send() {
	fmt.Println("具体同事类001发出请求")
	concreteColleague001.AbstractColleague.mediator.Relay(concreteColleague001.AbstractColleague)
}
