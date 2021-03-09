package standard

/**
State:抽象角色，触发各状态流转的操作步骤
*/
type State interface {
	Operation001(context *Context)
	Operation002(context *Context)
	GetCurrentState() string
}
