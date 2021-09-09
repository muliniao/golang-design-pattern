package sample001

/**
具体工厂: store factory
*/
type StoreFactory struct {
}

func (s *StoreFactory) GetCommodityService(commodityType int64) ICommodity {

	switch commodityType {
	case 1:
		return new(CardCommodityService)
	case 2:
		return new(CouponCommodityService)
	case 3:
		return new(GoodsCommodityService)
	}

	return nil

}
