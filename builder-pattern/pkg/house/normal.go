package house

import (
	"github.com/builder-pattern/house"
)

type IglooHouse struct {
	windowType string
	doorType   string
}

func (i *IglooHouse) getHouse() house.iHouseBuilder {
	return i
}

func (i *IglooHouse) withWindowType() {

}

func (i *IglooHouse) withDoorType() {

}
