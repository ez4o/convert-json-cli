package main

import (
	"encoding/json"
	"fmt"

	"ez4o.com/json-to/model"
)

func main() {
	const jsonString = `[
{
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
}
]`

	jc := model.JSONConverter{Writer: &model.GoWriter{}}

	abstractStructs, err := jc.GetAbstractStructs(jsonString)
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(abstractStructs, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
