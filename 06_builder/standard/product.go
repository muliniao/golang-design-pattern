package standard

/**
	Product(具体产品信息)
 */
type Product struct {
	partA string
	partB string
	partC string
}

func (product *Product) SetPartA (partA string) {
	product.partA = partA
}

func (product *Product) SetPartB (partB string) {
	product.partB = partB
}

func (product *Product) SetPartC (partC string) {
	product.partC = partC
}
