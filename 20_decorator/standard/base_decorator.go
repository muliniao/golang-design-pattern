package standard

/**
装饰者(基础)
*/
type BaseDecorator struct {
	component Component
}

func NewBaseDecorator(component Component) *BaseDecorator {
	return &BaseDecorator{
		component: component,
	}
}

func (BaseDecorator *BaseDecorator) Operation() {
	BaseDecorator.component.Operation()
}
