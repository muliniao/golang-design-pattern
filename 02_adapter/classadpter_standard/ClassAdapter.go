package classadpter_standard

type ClassAdapter struct {
	// 继承
	Adaptee
}

func (classAdapter *ClassAdapter) Request() {
	classAdapter.Adaptee.SpecificRequest()
}


