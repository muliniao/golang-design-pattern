package standard

type ConcreteMediator struct {
	ColleagueList []AbstractColleague
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{
		ColleagueList: make([]AbstractColleague, 0),
	}
}

func (concreteMediator *ConcreteMediator) Register(colleague AbstractColleague) {
	concreteMediator.ColleagueList = append(concreteMediator.ColleagueList, colleague)
	colleague.SetMedium(concreteMediator)
}

/**
	核心方法:处理所有concrete colleague的业务逻辑
 */
func (concreteMediator *ConcreteMediator) Relay(relayColleague AbstractColleague) {

	for _, v := range concreteMediator.ColleagueList {
		if v == relayColleague {
			v.Receive()
		}
	}

}
