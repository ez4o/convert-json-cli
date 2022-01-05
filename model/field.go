package model

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Field struct {
	Index    string
	TypeName string
}

func CreateField(index string, value interface{}) (Field, []Struct, error) {
	var f Field = Field{Index: index}
	var additionalStructs []Struct = make([]Struct, 0)
	var typeName string = reflect.TypeOf(value).String()

	switch typeName {
	case "[]interface {}":
		// array
		arrayTypeName, isPrimitive := IsPrimitive(value.([]interface{})[0])

		if isPrimitive {
			// primitive array
			f.TypeName = arrayTypeName + "[]"
		} else {
			// non-primitive array
			if index[len(index)-1:] == "s" {
				f.TypeName = strings.Title(index[:len(index)-1]) + "[]"
			} else {
				f.TypeName = strings.Title(index) + "[]"
			}

			b, err := json.Marshal(value.([]interface{})[0])
			if err != nil {
				return f, nil, err
			}

			structs, err := Parse(f.TypeName[:len(f.TypeName)-2], string(b))
			if err != nil {
				return f, nil, err
			}

			additionalStructs = append(additionalStructs, structs...)
		}

		break
	case "map[string]interface {}":
		f.TypeName = strings.Title(index)

		b, err := json.Marshal(value)
		if err != nil {
			return f, nil, err
		}

		structs, err := Parse(f.TypeName, string(b))
		if err != nil {
			return f, nil, err
		}

		additionalStructs = append(additionalStructs, structs...)

		break
	default:
		f.TypeName = typeName
	}

	return f, additionalStructs, nil
}

func IsPrimitive(v interface{}) (string, bool) {
	typeName := reflect.TypeOf(v).String()
	if typeName == "[]interface {}" || reflect.TypeOf(v).String() == "map[string]interface {}" {
		return typeName, false
	}

	return typeName, true
}
