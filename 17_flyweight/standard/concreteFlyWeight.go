package standard

import "fmt"

/**
	ConcreteFlyWeight:具体享元
*/
type ConcreteFlyWeight struct {
	Key	string
}

func NewConcreteFlyWeight (key string) *ConcreteFlyWeight {
	return &ConcreteFlyWeight{
		Key: key,
	}
}

func (concreteFlyWeight *ConcreteFlyWeight) Operation(unsharedConcreteFlyWeight *UnsharedConcreteFlyWeight) {
	fmt.Printf("非享元信息:[%s]", unsharedConcreteFlyWeight.GetInfo())
	fmt.Printf("具体享元[%s]被调用", concreteFlyWeight.Key)
}


