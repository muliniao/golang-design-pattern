package sample

type Director struct {
	houseBuilder HouseBuilder
}

func NewDirector(houseBuilder HouseBuilder) Director {
	return Director{
		houseBuilder: houseBuilder,
	}
}

func (director *Director) Construct() *House {
	director.houseBuilder.BuildBasic()
	director.houseBuilder.BuildWall()
	director.houseBuilder.BuildRoof()

	return director.houseBuilder.GetHouse()
}
