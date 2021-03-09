package sample

import "fmt"

/**
具体产品: Card
*/
type CardCommodityService struct {
}

func (c *CardCommodityService) SendCommodity(uId string, commodityId string, bizId string, Map map[string]string) error {

	mobile := queryUserMobile(uId)

	iQiYiCardService := new(iQiYiCardService)
	token, err := iQiYiCardService.grantToken(mobile, bizId)
	if err != nil {
		return err
	}

	if token == "" {
		return fmt.Errorf("[SendCardCommodity] token is nil")
	}

	return err

}

// ignore
func queryUserMobile(uId string) string {
	return "12345678901"
}

type iQiYiCardService struct {
}

func (i *iQiYiCardService) grantToken(mobile string, bizId string) (string, error) {

	return "!!@@!adcdbafasf", nil

}
