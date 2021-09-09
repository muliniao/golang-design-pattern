package main

import (
	"fmt"

	"golang-design-pattern/04_factory_method/sample002"
)

func main () {

	factoryType := "json"
	configText := "configText"
	parser := sample002.NewRuleConfigParserFactoryMap().GetParserFactory(factoryType).CreateParser()
	ruleConfig := parser.Parser(configText)
	fmt.Println(ruleConfig)

}
