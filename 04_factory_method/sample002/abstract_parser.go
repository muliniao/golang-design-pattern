package sample002

type IRuleConfigParser interface {
	Parser(configText string) string
}

