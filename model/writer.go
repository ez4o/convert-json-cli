package model

type IWriter interface {
	SetOutputPath(outputPath string)
	SetNested(nested bool)
	Write(abstractStructs []Struct) error
	GetStruct(abstractStruct Struct) string
	GetFields(fields []Field) string
	GetField(index int, field Field) string
	GetTypeName(typeName string) string
}
