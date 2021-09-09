package sample002

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}
