package sample

import "fmt"

/**
具体产品: Goods
*/
type GoodsCommodityService struct {
}

func (g *GoodsCommodityService) SendCommodity(uId string, commodityId string, bizId string, Map map[string]string) error {

	deliver := new(Deliver)
	deliver.SetUserName(uId)
	deliver.SetUserPhone(Map[uId])

	goodsService := new(GoodsService)
	isSuccess := goodsService.DeliverGoods(deliver)

	if !isSuccess {
		return fmt.Errorf("[SendGoodsCommodity] failed to send commodity")
	}

	return nil
}

// ignore
type Deliver struct {
}

func (d *Deliver) SetUserName(uId string) {

}

func (d *Deliver) SetUserPhone(phone string) {

}

type GoodsService struct {
}

func (g *GoodsService) DeliverGoods(deliver *Deliver) bool {
	return true
}
