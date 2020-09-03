package standard

/**
	Composite:树枝角色
 */
type Composite struct {
	Children []Component
}

func NewComposite() *Composite {
	return &Composite{
		Children: make([]Component, 0),
	}
}

func (composite *Composite) Add(component Component) {
	composite.Children = append(composite.Children, component)
}

func (composite *Composite) Remove(component Component) {
	for key, value := range composite.Children {
		if value == component {
			newChildren := append(composite.Children[:key], composite.Children[key+1:]...)
			composite.Children = newChildren
		}
	}
}

func (composite *Composite) GetChild(i int) Component {
	return composite.Children[i]
}

func (composite *Composite) Operation() {
	for _, value := range composite.Children {
		value.(Component).Operation()
	}
}
