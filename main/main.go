package main

import (
	"log"
	"validation"
)

type MyType struct {
	IntField    uint64
	StringField string
	NilPtrField *int
}

func main() {
	temp := 5
	myStruct := MyType{5, "test", &temp}

	err := validation.ValidateStruct(&myStruct,
		validation.Field(&myStruct.IntField, validation.UnsignedInt),
		validation.Field(&myStruct.StringField, validation.MatchRequired{"[a-z]"}, validation.Length{4, 20}),
		validation.Field(&myStruct.NilPtrField, validation.Required))

	if err != nil {
		log.Println(err)
	}
}
