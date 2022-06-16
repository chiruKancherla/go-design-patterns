package house

type iBuilder interface {
	withWindowType()
	withDoorType()
}
