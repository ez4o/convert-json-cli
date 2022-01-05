package model

import (
	"encoding/json"
	"model/struct_type"
)

type Struct struct {
	Name   string
	Fields []Field
}

func (s *Struct) AddField(f Field) {
	s.Fields = append(s.Fields, f)
}

func Parse(name string, jsonString string) ([]Struct, error) {
	var s Struct = Struct{Name: name}
	var abstractStructs []Struct = make([]Struct, 0)
	var result map[string]interface{}

	s.Fields = make([]Field, 0)

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return abstractStructs, err
	}

	for key, value := range result {
		f := CreateField(key, value)
		s.AddField(f)

		if f.StructType == struct_type.NonPrimitive {
			b, err := json.Marshal(value)
			if err != nil {
				return abstractStructs, err
			}

			structs, err := Parse(f.TypeName, string(b))
			if err != nil {
				return abstractStructs, err
			}

			abstractStructs = append(abstractStructs, structs...)
		} else if f.StructType == struct_type.NonPrimitiveArray {
			b, err := json.Marshal(value.([]interface{})[0])
			if err != nil {
				return abstractStructs, err
			}

			structs, err := Parse(f.TypeName[:len(f.TypeName)-2], string(b))
			if err != nil {
				return abstractStructs, err
			}

			abstractStructs = append(abstractStructs, structs...)
		}
	}

	abstractStructs = append(abstractStructs, s)
	return abstractStructs, nil
}
