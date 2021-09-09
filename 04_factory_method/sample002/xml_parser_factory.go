package sample002

type XMLRuleConfigParserFactory struct {

}

func (x *XMLRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return new(XMLRuleConfigParser)
}
