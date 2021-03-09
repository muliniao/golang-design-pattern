package standard

/**
Product(具体产品信息)
*/
type Product struct {
	partA string
	partB string
	partC string
}

func (product *Product) setPartA(partA string) {
	product.partA = partA
}

func (product *Product) setPartB(partB string) {
	product.partB = partB
}

func (product *Product) setPartC(partC string) {
	product.partC = partC
}
