package model

type IWriter interface {
	Write(abstractStructs []Struct) error
}
