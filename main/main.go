package main

import (
	"fmt"
	"validation"
)

type MyType struct {
	IntField    int64
	StringField string
	NilPtrField *int
}

func main() {
	temp := 5
	myStruct := MyType{5, "test", &temp}

	validation.LogCleaner()

	err := validation.ValidateStruct(&myStruct,
		validation.Field(&myStruct.IntField, validation.UnsignedInt),
		validation.Field(&myStruct.StringField, validation.MatchRequired{"[A-Z]"}, validation.Length{4, 20}),
		validation.Field(&myStruct.NilPtrField, validation.Required))

	if err != nil {
		fmt.Println("Result: Not valid... check the log file '../tmp/logger.log' for more information")
	}

	err = validation.ValidateStructInformative(&myStruct,
		validation.Field(&myStruct.IntField, validation.UnsignedInt),
		validation.Field(&myStruct.StringField, validation.MatchRequired{"[A-Z]"}, validation.Length{10, 20}),
		validation.Field(&myStruct.NilPtrField, validation.Required))

	if err != nil {
		fmt.Println(err)
	}
}
