package standard001

/**
State:抽象角色，触发各状态流转的操作步骤
*/
type IState interface {
	Event001()
	Event002()
	Event003()
	GetCurrentState() string
}
