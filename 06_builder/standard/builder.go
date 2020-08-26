package standard

/**
	Builder(抽象)
 */
type Builder interface {

	/**
		构建对象
	 */
	BuildPartA() Builder
	BuildPartB() Builder
	BuildPartC() Builder

	/**
		返回构建完成的Product
	 */
	GetResult() *Product
}

