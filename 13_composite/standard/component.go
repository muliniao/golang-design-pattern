package standard

/**
Component:抽象构建角色
*/
type Component interface {
	Add(component Component)
	Remove(component Component)
	GetChild(i int) Component
	Operation()
}
