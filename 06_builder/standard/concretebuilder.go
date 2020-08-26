package standard

/**
	ConcreteBuilder(具体的Builder)
 */
type ConcreteBuilder struct {
	product	*Product
}

func NewConcreteBuilder(product *Product) *ConcreteBuilder {
	return &ConcreteBuilder{product: product}
}

/**
	构建对象
 */
func (concreteBuilder *ConcreteBuilder) BuildPartA() Builder {
	concreteBuilder.product.SetPartA("建造 PartA")
	return concreteBuilder
}

func (concreteBuilder *ConcreteBuilder) BuildPartB() Builder {
	concreteBuilder.product.SetPartB("建造 PartB")
	return concreteBuilder
}

func (concreteBuilder *ConcreteBuilder) BuildPartC() Builder {
	concreteBuilder.product.SetPartC("建造 PartC")
	return concreteBuilder
}

/**
	返回构建完成的Product
 */
func (concreteBuilder *ConcreteBuilder) GetResult() *Product {
	return concreteBuilder.product
}
