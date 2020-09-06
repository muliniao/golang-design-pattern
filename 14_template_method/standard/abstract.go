package standard

import (
	"fmt"
)

type CaffeineBeverageTemplateMethod interface {
	Brew()
	AddCondiment()
	CustomWantsCondiment() bool
}

type CoffeeWithWithHook struct {
	CaffeineBeverageTemplateMethod
}

type CoffeeBrewAndCondiment struct {
}

func (o *CoffeeBrewAndCondiment) Brew() {
	fmt.Println("boiling")
}

func (o *CoffeeBrewAndCondiment) AddCondiment() {
	fmt.Println("adding condiment")
}

func (o *CoffeeBrewAndCondiment) CustomWantsCondiment() bool {
	return true
}


func (o *CoffeeWithWithHook) PrepareRecipe() {
	o.BoilWater()
	o.CaffeineBeverageTemplateMethod.Brew()
	o.PourInCup()
	if o.CaffeineBeverageTemplateMethod.CustomWantsCondiment() {
		o.CaffeineBeverageTemplateMethod.AddCondiment()
	}
}

func (o *CoffeeWithWithHook) BoilWater() {
	fmt.Println("boiling water")
}

func (o *CoffeeWithWithHook) PourInCup() {
	fmt.Println("pouring into cup")
}

func NewCaffeineBeverageWithHook() *CoffeeWithWithHook {
	return &CoffeeWithWithHook{
		&CoffeeBrewAndCondiment{},
	}
}
