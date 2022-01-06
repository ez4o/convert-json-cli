package main

import (
	"ez4o.com/convert-json-cli/model"
)

func main() {
	const jsonString = `[
    {
      "_id": 1,
      "deposit": 66666.66,
      "name": "test",
      "jobs": [
        {
          "company": {
            "name": "test_company",
            "bankrupt": false
          },
          "title": ["test"]
        }
      ]
    }
  ]`

	writers := []model.IWriter{
		&model.GoWriter{},
		&model.DartWriter{},
		&model.CppWriter{},
		&model.ProtobufWriter{},
		&model.JavaWriter{},
		&model.CWriter{},
		&model.KotlinWriter{},
		&model.PythonWriter{},
	}

	for _, writer := range writers {
		jc := model.JSONConverter{Writer: writer}
		err := jc.Convert(jsonString)
		if err != nil {
			panic(err)
		}
	}
}
