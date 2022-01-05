package util

import (
	"encoding/json"
	"model"
)

func GetAbstractStructs(jsonString string) ([]model.Struct, error) {
	if jsonString[0] == '[' {
		var result []interface{}

		err := json.Unmarshal([]byte(jsonString), &result)
		if err != nil {
			return nil, err
		}

		b, err := json.Marshal(result[0])
		if err != nil {
			return nil, err
		}

		abstractStructs, err := model.Parse("Nested", string(b))
		if err != nil {
			return nil, err
		}

		abstractStructs = append(abstractStructs, model.Struct{Name: "AutoGenerated", Fields: []model.Field{{Index: "", TypeName: "Nested[]"}}})

		return abstractStructs, nil
	} else {
		abstractStructs, err := model.Parse("AutoGenerated", jsonString)
		if err != nil {
			return nil, err
		}

		return abstractStructs, nil
	}
}
