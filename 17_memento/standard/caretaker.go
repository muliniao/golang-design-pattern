package standard

/**
Caretaker: 管理者
*/
type Caretaker struct {
	memento *Memento
}

func (c *Caretaker) SetMemento(memento *Memento) {
	c.memento = memento
}

func (c *Caretaker) GetMemento() *Memento {
	return c.memento
}
