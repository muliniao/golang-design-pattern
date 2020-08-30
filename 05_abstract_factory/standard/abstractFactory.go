package standard

type AbstractFactory interface {
	NewProduct001() Product001
	NewProduct002() Product002
}
