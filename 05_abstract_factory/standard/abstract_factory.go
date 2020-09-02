package standard

/**
	抽象工厂
 */
type AbstractFactory interface {
	NewProduct001() Product001
	NewProduct002() Product002
}
