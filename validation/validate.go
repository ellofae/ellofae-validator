package validation

import (
	"errors"
	"fmt"
	"reflect"
)

// Add "github.com/pkg/errors" package for a Wrapf func

var (
	ErrStructNotValid             = errors.New("error: struct is not valid")
	ErrValueNotUnsignedInt        = errors.New("error: possible values - uint8, uint16, uint32, uint64, uint")
	ErrRegexNotSatisfied          = errors.New("error: regex expression was not satisfied")
	ErrNotValidatable             = errors.New("error: data type must be a pointer on struct or nil")
	ErrFieldNotSpecified          = errors.New("error: field cannot be nil because of the non-nil field value requirement")
	ErrOnlyStringValue            = errors.New("error: value must be a string")
	ErrStringLengthIsNotSatisfied = errors.New("error: string's length is beyond the limit of min value or max value")
)

// Function that prints out all errors that occured during the validation from the slice where they are contained
func logErrorsToUser(errorList []error, structName interface{}) {
	fmt.Printf("(%s) Validation errors occured:\n", reflect.ValueOf(structName).Elem().Type())
	for _, err := range errorList {
		fmt.Println("\t", err)
	}
}

// Function that validates fields of a struct
// Where first argument is a field value and second argument is a slice of rules for validation of the field
func Validate(value interface{}, rules RuleSlice) error {
	for _, rule := range rules {
		err := rule.Specifier(value)
		if err != nil {
			return err
		}
	}

	return nil
}
