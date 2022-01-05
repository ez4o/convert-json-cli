package main

import (
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
	err := jc.Convert(jsonString)
	if err != nil {
		panic(err)
	}
}
