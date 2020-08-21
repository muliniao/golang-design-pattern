package decorator_standard

/**
	装饰者(抽象)
 */
type Decorator struct {
	component	Component
}

func NewDecorator(component Component) *Decorator {
	return &Decorator{
		component: component,
	}
}

func (decorator *Decorator) Operation() {
	decorator.component.Operation()
}

