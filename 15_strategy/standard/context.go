package standard

/**
Context(环境类)
*/
type Context struct {
	strategy Strategy
}

/**
传入策略
*/
func NewContext(strategy Strategy) *Context {
	return &Context{
		strategy: strategy,
	}
}

/**
传入动态策略
*/
func (context *Context) SetStrategy(strategy Strategy) {
	context.strategy = strategy
}

func (context *Context) GetStrategy() Strategy {
	return context.strategy
}

func (context *Context) ConcreteStrategyMethod() {
	context.strategy.StrategyMethod()
}
