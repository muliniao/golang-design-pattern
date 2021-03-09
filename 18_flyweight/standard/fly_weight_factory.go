package standard

import "fmt"

/**
FlyWeightFactory:享元工厂
*/
type FlyWeightFactory struct {
	Pool map[string]FlyWeight
}

func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{
		Pool: make(map[string]FlyWeight),
	}
}

func (flyWeightFactory *FlyWeightFactory) GetFlyWeight(key string) FlyWeight {

	if value, ok := flyWeightFactory.Pool[key]; ok {
		return value
	}

	flyWeightFactory.Pool[key] = NewConcreteFlyWeight(key)
	return flyWeightFactory.Pool[key]
}

func (flyWeightFactory *FlyWeightFactory) GetSize() {
	fmt.Println(len(flyWeightFactory.Pool))
}
