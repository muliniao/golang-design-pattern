package standard

/**
	Context(环境类)
 */
type Context struct {
	Strategy Strategy
}

/**
	传入策略
 */
func NewContext(strategy Strategy) *Context {
	return &Context{
		Strategy: strategy,
	}
}

/**
	传入动态策略
*/
func (context *Context) SetStrategy(strategy Strategy) {
	context.Strategy = strategy
}

func (context *Context) GetStrategy() Strategy {
	return context.Strategy
}

func (context *Context) ConcreteStrategyMethod(){
	context.Strategy.StrategyMethod()
}
