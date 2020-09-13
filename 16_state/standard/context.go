package standard

import "fmt"

/**
	Context:上下文
 */
type Context struct {
	state	State
}

func NewContext(state State) *Context {
	return &Context{
		state: state,
	}
}

func (context *Context) Operation001(ctx *Context) {
	context.state.Operation001(ctx)
}

func (context *Context) Operation002(ctx *Context) {
	context.state.Operation002(ctx)
}

func (context *Context) SetState(state State) {
	context.state = state
}

func (context *Context) GetState() State {
	return context.state
}

func (context *Context) GetCurrentState() string {
	fmt.Println("当前状态:" + context.state.GetCurrentState())
	return context.state.GetCurrentState()
}
