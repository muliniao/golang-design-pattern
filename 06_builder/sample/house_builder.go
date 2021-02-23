package sample

type HouseBuilder interface {
	BuildBasic()
	BuildWall()
	BuildRoof()
	GetHouse() *House
}
