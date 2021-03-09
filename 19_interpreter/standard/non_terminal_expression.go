package standard

/**
Non Terminal Expression
*/
type NonTerminalExpression struct {
	abstractExpression001 AbstractExpression
	abstractExpression002 AbstractExpression
}

/**
非终结表达式的处理
*/
func (*NonTerminalExpression) Interpret(info string) interface{} {
	return nil
}
