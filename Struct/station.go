package Struct

type FuelStand struct {
	Id    int
	Type  string
	Queue chan *Car
}
