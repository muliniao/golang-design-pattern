package standard

/**
ConcreteCommand: 具体命令角色
*/
type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{
		receiver: receiver,
	}
}

func (concreteCommand *ConcreteCommand) Execute() {
	concreteCommand.receiver.Action()
}
