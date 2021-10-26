package standard

type ConcreteIterator struct {
	Cursor	int
	List 	[]interface{}
}

func NewConcreteIterator(list []interface{}) *ConcreteIterator {
	return &ConcreteIterator{
		Cursor: 0,
		List:   list,
	}
}

func (c *ConcreteIterator) HasNext() bool {
	return c.Cursor != len(c.List) - 1
}

func (c *ConcreteIterator) Next() interface{} {
	c.Cursor += 1
	object := c.List[c.Cursor]
	return object
}

func (c *ConcreteIterator) CurrentItem() interface{} {
	return c.List[c.Cursor]
}
