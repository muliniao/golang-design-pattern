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
func (concreteBuilder *ConcreteBuilder) BuildPartA() {
	concreteBuilder.product.SetPartA("建造 PartA")
}

func (concreteBuilder *ConcreteBuilder) BuildPartB() {
	concreteBuilder.product.SetPartB("建造 PartB")
}

func (concreteBuilder *ConcreteBuilder) BuildPartC() {
	concreteBuilder.product.SetPartC("建造 PartC")
}

/**
	返回构建完成的Product
 */
func (concreteBuilder *ConcreteBuilder) GetResult() *Product {
	return concreteBuilder.product
}
