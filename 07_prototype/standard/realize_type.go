package standard

/**
具体原型类
*/
type RealizeType struct {
	Name string
	Age  int
}

func (r *RealizeType) Clone() Cloneable {
	realizeType := *r
	return &realizeType
}
