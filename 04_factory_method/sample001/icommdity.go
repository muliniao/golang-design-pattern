package sample001

/**
抽象产品: Commodity

模拟发放多种奖品
- 1.优惠券
- 2.实物商品
- 3.爱奇艺的兑换卡
*/
type ICommodity interface {
	SendCommodity(uId string, commodityId string, bizId string, Map map[string]string) error
}
