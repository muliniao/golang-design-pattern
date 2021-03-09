package standard

/**
ConcreteBuilder(具体的Builder)
*/
type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		product: &Product{},
	}
}

/**
构建对象
*/
func (concreteBuilder *ConcreteBuilder) BuildPartA() {
	concreteBuilder.product.setPartA("建造 PartA")
}

func (concreteBuilder *ConcreteBuilder) BuildPartB() {
	concreteBuilder.product.setPartB("建造 PartB")
}

func (concreteBuilder *ConcreteBuilder) BuildPartC() {
	concreteBuilder.product.setPartC("建造 PartC")
}

/**
返回构建完成的Product
*/
func (concreteBuilder *ConcreteBuilder) GetResult() *Product {
	return concreteBuilder.product
}
