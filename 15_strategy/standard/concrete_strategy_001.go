package standard

import "fmt"

/**
	ConcreteStrategy001(具体策略类)
 */
type ConcreteStrategy001 struct {

}

func (concreteStrategy001 *ConcreteStrategy001) StrategyMethod() {
	fmt.Println("具体策略001")
}
