package sample

type PromotionObserver struct {
	promotionService PromotionService
}

func (p * PromotionObserver) HandleRegisterSuccess(userID int) {
	p.promotionService.issueNewUserExperienceCash(userID)
}




type PromotionService struct {

}

func (p *PromotionService) issueNewUserExperienceCash(userID int)  {

}
