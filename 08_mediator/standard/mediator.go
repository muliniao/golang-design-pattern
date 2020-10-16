package standard

type Mediator interface {
	Register(colleague AbstractColleague)
	Relay(relayColleague AbstractColleague)
}
