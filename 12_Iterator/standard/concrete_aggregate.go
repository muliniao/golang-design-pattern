package standard

/**
	ConcreteAggregate: 具体聚合
 */
type ConcreteAggregate struct {
	List	[]interface{}
}

func NewConcreteAggregate() *ConcreteAggregate {
	return &ConcreteAggregate{
		List: make([]interface{}, 0),
	}
}

func (c *ConcreteAggregate) Add(Object interface{}) {
	c.List = append(c.List, Object)
}

func (c *ConcreteAggregate) Remove(Object interface{}) {
	for key, value := range c.List {
		if value == Object {
			newList := append(c.List[:key], c.List[key+1:]...)
			c.List = newList
		}
	}
}

func (c *ConcreteAggregate) GetIterator() Iterator {
	return NewConcreteIterator(c.List)
}
