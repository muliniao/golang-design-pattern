package _0_simple_factory

/**
	简单工厂模式
 */
type Factory interface {
	Operation()
}

/**
	传入Type,返回实例
 */
func NewFactory(carType string) Factory{
	switch carType {
	case "Audi":
		return &Audio{}
	case "Volvo":
		return &Volvo{}
	default:
		return nil
	}
	return nil
}
