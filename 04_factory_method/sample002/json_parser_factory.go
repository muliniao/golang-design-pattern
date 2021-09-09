package sample002

type JsonRuleConfigParserFactory struct {

}

func (j *JsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return new(JsonRuleConfigParser)
}