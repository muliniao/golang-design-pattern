package standard

import "fmt"

/**
	ConcreteStrategy002(具体策略类)
*/
type ConcreteStrategy002 struct {

}

func (concreteStrategy001 *ConcreteStrategy002) StrategyMethod() {
	fmt.Println("具体策略002")
}
