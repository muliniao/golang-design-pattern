package standard

type IColleague interface {
	Receive()
	Send()
}

type AbstractColleague struct {
	IColleague
	mediator Mediator
}

func (abstractColleague *AbstractColleague)SetMedium(mediator Mediator) {
	abstractColleague.mediator = mediator
}
