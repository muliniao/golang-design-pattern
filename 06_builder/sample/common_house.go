package sample

type CommonHouse struct {
	house *House
}

func NewCommonHouse() *CommonHouse {
	return &CommonHouse{
		house: &House{},
	}
}

func (c *CommonHouse) BuildBasic() {
	c.house.Basic = "普通房子的地基"
}

func (c *CommonHouse) BuildWall() {
	c.house.Wall = "普通房子的砌墙"
}

func (c *CommonHouse) BuildRoof() {
	c.house.Roof = "普通房子的封顶"
}

func (c *CommonHouse) GetHouse() *House {
	return c.house
}
