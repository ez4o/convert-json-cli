package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Field struct {
	Index    string
	TypeName string
}

func CreateField(index string, value interface{}) (Field, []Struct, error) {
	var f Field = Field{Index: index}
	var additionalStructs []Struct = make([]Struct, 0)
	typeName, _ := IsPrimitive(value)

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

	if strings.Contains(typeName, "float") {
		if _, err := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 16); err == nil {
			return "int", true
		} else if _, err := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 32); err == nil {
			return "int32", true
		} else if _, err := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64); err == nil {
			return "int64", true
		} else if _, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 32); err == nil {
			return "float32", true
		} else if _, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64); err == nil {
			return "float64", true
		}
	}

	return typeName, true
}
