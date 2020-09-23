package standard

/**
	Memento: 备忘录
 */
type Memento struct {
	State	string
}

func NewMemento(state string) *Memento {
	return &Memento{
		State: state,
	}
}

func (m *Memento) SetState(state string) {
	m.State = state
}

func (m *Memento) GetState() string {
	return m.State
}
