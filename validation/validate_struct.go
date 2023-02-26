package validation

import (
	"fmt"
	"log"
	"reflect"
)

// Type for a field that requires validation, which contains a pointer to a specific field of a struct
// and a variety of rules that user wants to apply to validation of a structure's field
type (
	FieldType struct {
		fieldPtr interface{}
		rules    RuleSlice // string -> rules
	}
)

// Funcation that returns a pointer of a FieldType
// Used in validation.ValidateStruct when user is specifing a field to validate and rules for validation
func Field(field interface{}, rules ...Rule) *FieldType {
	return &FieldType{
		fieldPtr: field,
		rules:    rules,
	}
}

// Funcation for validating a struct and its fields with specific rules
// It return nil if the struct and its fields are valid and the one specific error if it occures
// If an error occurs it sends a message about the error to the log file, but only the first error.
func ValidateStruct(data interface{}, fields ...*FieldType) error {
	var insideError error
	aLog, err := logCreation()
	if err != nil {
		log.Fatal(err, ErrLogFileNotOpened)
	}

	defer func() {
		writeToLog(aLog, fmt.Errorf("%w (data type: %v)", insideError, reflect.ValueOf(data).Type()))
	}()

	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Ptr || !val.IsNil() && val.Elem().Kind() != reflect.Struct {
		insideError = ErrNotValidatable
		return ErrNotValidatable
	}
	if val.IsNil() {
		insideError = ErrNilStruct
		return ErrNilStruct
	}

	for _, field := range fields {
		err := Validate(field.fieldPtr, field.rules)
		if err != nil {
			insideError = err
			return err
		}
	}

	return nil
}

// Funcation for validating a struct and its fields with specific passed rules
// Function is more for informative use, it doesn't return a specific error, but logs out errors
// to the log file that occur during the validation
func ValidateStructInformative(data interface{}, fields ...*FieldType) error {
	aLog, err := logCreation()
	if err != nil {
		log.Fatal(err, ErrLogFileNotOpened)
	}

	val := reflect.ValueOf(data)
	var validBool bool = false

	if val.Kind() != reflect.Ptr || !val.IsNil() && val.Elem().Kind() != reflect.Struct {
		aLog.Println("*terminated: ", ErrNotValidatable)
		return ErrNotValidatable
	}
	if val.IsNil() {
		aLog.Println("*terminated: ", ErrNilStruct)
		return ErrNilStruct
	}

	for _, field := range fields {
		err := Validate(field.fieldPtr, field.rules)
		if err != nil {
			aLog.Println(err)
			validBool = true
		}
	}

	if validBool {
		aLog.Printf("Struct '%v' is NOT valid", val.Type())
		return ErrStructNotValid
	} else {
		aLog.Printf("Struct '%v' is valid", val.Type())
		return nil
	}
}
