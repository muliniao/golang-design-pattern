package objectAdapter_standard

/**
	适配器
 */
type ObjectAdapter struct {
	adaptee *Adaptee
}

func NewObjectAdapter(adaptee *Adaptee) *ObjectAdapter {
	return &ObjectAdapter{
		adaptee: adaptee,
	}
}

func (objectAdapter *ObjectAdapter) Request () {
	objectAdapter.adaptee.SpecificRequest()
}

