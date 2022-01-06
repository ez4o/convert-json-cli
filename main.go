package main

import (
	"ez4o.com/convert-json-cli/model"
)

func main() {
	const jsonString = `[
    {
      "_id": 1,
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

	writers := []model.IWriter{
		&model.DartWriter{},
		&model.CppWriter{},
		&model.GoWriter{},
	}

	for _, writer := range writers {
		jc := model.JSONConverter{Writer: writer}
		err := jc.Convert(jsonString)
		if err != nil {
			panic(err)
		}
	}
}
