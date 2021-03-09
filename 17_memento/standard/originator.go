package standard

/**
Originator: 发起人
*/
type Originator struct {
	State string
}

func (o *Originator) CreateMemento() *Memento {
	return NewMemento(o.State)
}

func (o *Originator) RestoreMemento(m *Memento) {
	o.SetState(m.State)
}

func (o *Originator) SetState(state string) {
	o.State = state
}

func (o *Originator) GetState() string {
	return o.State
}
