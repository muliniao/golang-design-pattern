package objectAdapter_standard

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

