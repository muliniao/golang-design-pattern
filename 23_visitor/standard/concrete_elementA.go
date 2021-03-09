package standard

/**
Concrete Element
*/
type ConcreteElementA struct {
}

func (concreteElementA ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(concreteElementA)
}

func (concreteElementA ConcreteElementA) OperationA() string {
	return "具体元素A的操作"
}
