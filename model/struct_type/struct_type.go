package struct_type

type StructType int

const (
	Primitive StructType = iota
	NonPrimitive
	PrimitiveArray
	NonPrimitiveArray
)
