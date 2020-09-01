package standard

/**
	UnsharedConcreteFlyWeight:非享元
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
