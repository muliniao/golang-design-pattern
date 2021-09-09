package sample002

type RuleConfigParserFactoryMap struct {
	CachedFactories map[string]IRuleConfigParserFactory
}

func NewRuleConfigParserFactoryMap() *RuleConfigParserFactoryMap {

	r := &RuleConfigParserFactoryMap{
		CachedFactories: make(map[string]IRuleConfigParserFactory),
	}

	r.CachedFactories["json"] = new(JsonRuleConfigParserFactory)
	r.CachedFactories["xml"] = new(XMLRuleConfigParserFactory)

	return r
}

func (r *RuleConfigParserFactoryMap) GetParserFactory(factoryType string) IRuleConfigParserFactory {
	return r.CachedFactories[factoryType]
}
