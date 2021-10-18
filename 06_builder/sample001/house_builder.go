package sample001

type HouseBuilder interface {
	BuildBasic()
	BuildWall()
	BuildRoof()
	GetHouse() *House
}
