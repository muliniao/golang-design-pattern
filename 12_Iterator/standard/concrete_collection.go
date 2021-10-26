package standard

type ConcreteCollection struct {
	List []interface{}
}

func NewConcreteCollection() *ConcreteCollection {
	return &ConcreteCollection{
		List: make([]interface{}, 0),
	}
}

func (l *ConcreteCollection) CreateIterator() Iterator {
	return NewConcreteIterator(l.List)
}