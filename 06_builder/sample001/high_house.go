package sample001

type HighBuildingHouse struct {
	house *House
}

func NewHighBuildingHouse() *HighBuildingHouse {
	return &HighBuildingHouse{
		house: &House{},
	}
}

func (h *HighBuildingHouse) BuildBasic() {
	h.house.Basic = "高楼房子的地基"
}

func (h *HighBuildingHouse) BuildWall() {
	h.house.Wall = "高楼房子的砌墙"
}

func (h *HighBuildingHouse) BuildRoof() {
	h.house.Roof = "高楼房子的封顶"
}

func (h *HighBuildingHouse) GetHouse() *House {
	return h.house
}
