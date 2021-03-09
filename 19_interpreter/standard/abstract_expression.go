package standard

/**
Abstract Expression
*/
type AbstractExpression interface {
	Interpret(info string) interface{}
}
