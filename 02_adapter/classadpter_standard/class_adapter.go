package classadpter_standard

/**
适配器
*/
type ClassAdapter struct {
	// 继承
	Adaptee
}

func (classAdapter *ClassAdapter) Request() {
	classAdapter.Adaptee.SpecificRequest()
}
