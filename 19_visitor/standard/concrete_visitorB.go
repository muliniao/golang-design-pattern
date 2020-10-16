package standard

import "fmt"

/**
	Concrete Visitor
*/
type ConcreteVisitorB struct {

}

func (concreteVisitorB *ConcreteVisitorB) VisitConcreteElementA(element ConcreteElementA) {
	fmt.Println("具体访问者B访问--->" + element.OperationA())
}

func (concreteVisitorB *ConcreteVisitorB) VisitConcreteElementB(element ConcreteElementB) {
	fmt.Println("具体访问者B访问--->" + element.OperationB())
}
