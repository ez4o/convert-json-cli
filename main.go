package main

import (
	"ez4o.com/convert-json-cli/model"
)

func main() {
	const jsonString = `{
    "_id": "61d54c090c28b92801f59fa9",
    "index": 0,
    "guid": "56a2f095-e8ce-4dfa-a642-15242b5f4bca",
    "isActive": false,
    "balance": "$1,133.93",
    "picture": "http://placehold.it/32x32",
    "age": 34,
    "eyeColor": "brown",
    "name": "Yolanda Charles",
    "gender": "female",
    "company": "ACCEL",
    "email": "yolandacharles@accel.com",
    "phone": "+1 (891) 456-2260",
    "address": "950 Columbia Street, Juntura, Georgia, 8664",
    "about": "Officia veniam et mollit adipisicing aliquip aute incididunt do aliqua laborum aliquip veniam. Sint do velit culpa duis ut deserunt. Esse laboris eu Lorem amet ad non mollit exercitation mollit irure duis dolor fugiat. Ex in deserunt aute minim minim. Nostrud adipisicing fugiat eu incididunt dolore ullamco voluptate minim minim consectetur cillum veniam deserunt velit. Deserunt ut in laborum tempor. Aliquip magna incididunt et aliqua anim laborum eiusmod ullamco sunt veniam labore ex.\r\n",
    "registered": "2020-12-18T11:01:04 -08:00",
    "latitude": 49.901994,
    "longitude": -78.153293,
    "tags": [
      "voluptate",
      "reprehenderit",
      "exercitation",
      "nisi",
      "amet",
      "voluptate",
      "irure"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Rollins Booker",
        "jobs": [
          {
            "company": {
              "name": "test_company"
            },
            "title": ["test"]
          }
        ]
      },
      {
        "id": 1,
        "name": "Pena Casey",
        "jobs": [
          {
            "company": {
              "name": "test_company"
            },
            "title": ["test"]
          }
        ]
      },
      {
        "id": 2,
        "name": "Holloway Wilkinson",
        "jobs": [
          {
            "company": {
              "name": "test_company"
            },
            "title": ["test"]
          }
        ]
      }
    ],
    "greeting": "Hello, Yolanda Charles! You have 4 unread messages.",
    "favoriteFruit": "apple"
  }`

	jc1 := model.JSONConverter{Writer: &model.GoWriter{}}
	err := jc1.Convert(jsonString)
	if err != nil {
		panic(err)
	}

	jc2 := model.JSONConverter{Writer: &model.DartWriter{}}
	err = jc2.Convert(jsonString)
	if err != nil {
		panic(err)
	}
}
