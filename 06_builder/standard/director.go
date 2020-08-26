package standard

/**
	Director(指挥者)
 */
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) Director {
	return Director{
		builder: builder,
	}
}

/**
	产品构建和组装
 */
func (director *Director) Construct() *Product {

	return director.builder.BuildPartA().BuildPartB().BuildPartC().GetResult()

}
