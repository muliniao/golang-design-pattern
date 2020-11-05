package sample

type ICommodity interface {
	SendCommodity(uId string, commodityId string, bizId string, Map map[string]string) error
}
