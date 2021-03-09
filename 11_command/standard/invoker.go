package standard

import "fmt"

/**
Invoker: 调用者/请求者角色
*/
type Invoker struct {
	command Command
}

func NewInvoker(command Command) *Invoker {
	return &Invoker{
		command: command,
	}
}

func (invoker *Invoker) SetCommand(command Command) {
	invoker.command = command
}

func (invoker *Invoker) Call() {
	fmt.Println("调用者执行命令Command")
	invoker.command.Execute()
}
