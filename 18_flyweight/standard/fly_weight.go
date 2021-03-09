package standard

/**
FlyWeight:抽象享元
*/
type FlyWeight interface {
	Operation(unsharedConcreteFlyWeight *UnsharedConcreteFlyWeight)
}
