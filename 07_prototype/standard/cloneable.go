package standard

/**
	抽象原型类
 */
type Cloneable interface {
	Clone() Cloneable
}
