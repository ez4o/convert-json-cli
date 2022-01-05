package model

import (
	"encoding/json"
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
		return nil, err
	}

	for key, value := range result {
		f, additionalStructs, err := CreateField(key, value)
		if err != nil {
			return nil, err
		}

		s.AddField(f)
		abstractStructs = append(abstractStructs, additionalStructs...)
	}

	abstractStructs = append(abstractStructs, s)
	return abstractStructs, nil
}
