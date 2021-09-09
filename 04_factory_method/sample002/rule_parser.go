package sample002

// JsonRuleConfigParser */
type JsonRuleConfigParser struct {

}

func (j *JsonRuleConfigParser) Parser(configText string) string {
	return "json parser"
}

// XMLRuleConfigParser */
type XMLRuleConfigParser struct {

}

func (x *XMLRuleConfigParser) Parser(configText string) string {
	return "xml parser"
}