package classadpter_standard

/**
适配器
*/
type ClassAdapter struct {
	//如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现继承
	//如果一个struct嵌套了另一个【有名】的结构体，那么这个模式叫做组合
	//如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现多重继承
	Adaptee
}

func (classAdapter *ClassAdapter) Request() {
	classAdapter.Adaptee.SpecificRequest()
}
