package model

import (
	struct_type "model/struct_type"
	"reflect"
	"strings"
)

type Field struct {
	Index      string
	TypeName   string
	StructType struct_type.StructType
}

func CreateField(index string, value interface{}) Field {
	f := Field{
		Index: index,
	}

	typeName := reflect.TypeOf(value).String()
	switch typeName {
	case "[]interface {}":
		arrayTypeName, isPrimitive := IsPrimitive(value.([]interface{})[0])
		if isPrimitive {
			f.TypeName = arrayTypeName + "[]"
			f.StructType = struct_type.PrimitiveArray
		} else {
			if index[len(index)-1:] == "s" {
				f.TypeName = strings.Title(index[:len(index)-1]) + "[]"
			} else {
				f.TypeName = strings.Title(index) + "[]"
			}
			f.StructType = struct_type.NonPrimitiveArray
		}
	case "map[string]interface {}":
		f.TypeName = strings.Title(index)
		f.StructType = struct_type.NonPrimitive
	default:
		f.TypeName = typeName
		f.StructType = struct_type.Primitive
	}

	return f
}

func IsPrimitive(v interface{}) (string, bool) {
	typeName := reflect.TypeOf(v).String()
	if typeName == "[]interface {}" || reflect.TypeOf(v).String() == "map[string]interface {}" {
		return typeName, false
	}

	return typeName, true
}
