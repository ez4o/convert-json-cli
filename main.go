package main

import (
	"encoding/json"
	"fmt"
	"model"
)

func main() {
	const jsonString = `{
  "id": 1,
  "name": "test",
  "jobs": [
    {
      "company": {
        "name": "test_company"
      },
      "title": ["test"]
    }
  ]
}`

	abstractStructs, err := model.Parse("AutoGenerated", jsonString)
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(abstractStructs, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}