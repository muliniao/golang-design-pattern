package standard

/**
	UnsharedConcreteFlyWeight:非享元
	外部状态: 指随环境改变而改变的不可以共享的部分。享元模式的实现要领就是区分应用中的这两种状态，并将外部状态外部化
 */
type UnsharedConcreteFlyWeight struct {
	Info	string
}

func NewUnsharedConcreteFlyWeight(info string) *UnsharedConcreteFlyWeight {
	return &UnsharedConcreteFlyWeight{
		Info: info,
	}
}

func (unsharedConcreteFlyWeight *UnsharedConcreteFlyWeight) GetInfo() string {
	return unsharedConcreteFlyWeight.Info
}

func (unsharedConcreteFlyWeight *UnsharedConcreteFlyWeight) SetInfo(info string) {
	unsharedConcreteFlyWeight.Info = info
}
