package standard

/**
	Facade:外观
 */
type Facade struct {
	subSystem001 *SubSystem001
	subSystem002 *SubSystem002
}

func NewFacade(subSystem001 *SubSystem001, subSystem002 *SubSystem002) *Facade {
	return &Facade{
		subSystem001: subSystem001,
		subSystem002: subSystem002,
	}
}

func (facade *Facade) Method() {
	facade.subSystem001.method001()
	facade.subSystem002.method002()
}
