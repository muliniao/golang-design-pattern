package standard

/**
	Visitor
 */
type Visitor interface {
	VisitConcreteElementA(element ConcreteElementA)
	VisitConcreteElementB(element ConcreteElementB)
}

