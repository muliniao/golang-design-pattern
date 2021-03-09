package standard

/*
	抽象工厂
*/

type AbstractFactory interface {
	NewProduct() Product
}
