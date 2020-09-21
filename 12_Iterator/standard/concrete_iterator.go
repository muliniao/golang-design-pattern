package standard

/**
	ConcreteIterator: 具体迭代器
 */
type ConcreteIterator struct {
	List  		[]interface{}
	Position	int
}

func NewConcreteIterator(list []interface{}) *ConcreteIterator {
	return &ConcreteIterator{
		List:     list,
		Position: 0,
	}
}

func (c *ConcreteIterator) First() interface{} {
	c.Position = 0
	return c.List[c.Position]
}

func (c *ConcreteIterator) Next() interface{} {
	if c.HasNext() {
		object := c.List[c.Position]
		c.Position += 1
		return object
	}
	return nil
}

func (c *ConcreteIterator) HasNext() bool {
	if c.Position >= len(c.List) {
		return false
	} else {
		return true
	}
}
