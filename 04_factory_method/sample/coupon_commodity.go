package sample

import "fmt"

/**
	具体产品: Coupon
*/
type CouponCommodityService struct {
	
}

func (c *CouponCommodityService) SendCommodity(uId string, commodityId string, bizId string, Map map[string]string) error {
	
	couponService := new(CouponService)
	result, err := couponService.SendCoupon(uId, commodityId, bizId)
	if err != nil {
		return err
	}
	
	if result != 0000 {
		return fmt.Errorf("[SendCouponCommodity] send commodity err [+%v] ", err)
	}

	return nil
}



// ignore
type CouponService struct {
	
}

func (c *CouponService) SendCoupon(uId string, commodityId string, bizId string) (resp int64, err error) {
	return 0, nil
}
