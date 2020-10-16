package standard

import "fmt"

/**
	Concrete Visitor
 */
type ConcreteVisitorA struct {

}

func (concreteVisitorA *ConcreteVisitorA) VisitConcreteElementA(element ConcreteElementA) {
	fmt.Println("具体访问者A访问--->" + element.OperationA())
}

func (concreteVisitorA *ConcreteVisitorA) VisitConcreteElementB(element ConcreteElementB) {
	fmt.Println("具体访问者A访问--->" + element.OperationB())
}
