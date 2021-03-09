package standard

/**
Concrete Element
*/
type ConcreteElementB struct {
}

func (concreteElementB ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(concreteElementB)
}

func (concreteElementB ConcreteElementB) OperationB() string {
	return "具体元素B的操作"
}
