package validation

import (
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
func ValidateStruct(data interface{}, fields ...*FieldType) error {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Ptr || !val.IsNil() && val.Elem().Kind() != reflect.Struct {
		return ErrNotValidatable
	}
	if val.IsNil() {
		return nil
	}

	for _, field := range fields {
		err := Validate(field.fieldPtr, field.rules)
		if err != nil {
			return err
		}
	}

	return nil
}

// Funcation for validating a struct and its fields with specific passed rules
// Function is more for informative use, it doesn't return a specific error, but prints out errros that occured
func ValidateStructInformative(data interface{}, fields ...*FieldType) error {
	errorList := make([]error, 0) // list of occured errors
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Ptr || !val.IsNil() && val.Elem().Kind() != reflect.Struct {
		return ErrNotValidatable
	}
	if val.IsNil() {
		return nil
	}

	for _, field := range fields {
		err := Validate(field.fieldPtr, field.rules)
		if err != nil {
			// the first error of the field being not valid is added and then starts validating the next field
			errorList = append(errorList, err)
		}
	}

	if len(errorList) != 0 {
		logErrorsToUser(errorList, data)
		return ErrStructNotValid
	} else {
		log.Printf("Struct '%v' is valid", val.Type())
		return nil
	}
}
