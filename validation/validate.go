package validation

import (
	"errors"
)

// Errors that are sent to a user and to the log file
// User can work with returned errors to get the information on the specific error that occured
// during the validation
var (
	ErrNilStruct                  = errors.New("error: data structure has to be defined and not nil")
	ErrLogFileNotOpened           = errors.New("error: log file was not successfully opened")
	ErrStructNotValid             = errors.New("error: struct is not valid")
	ErrValueNotUnsignedInt        = errors.New("error: possible values - uint8, uint16, uint32, uint64, uint")
	ErrRegexNotSatisfied          = errors.New("error: regex expression was not satisfied")
	ErrNotValidatable             = errors.New("error: data type must be a pointer on struct or nil")
	ErrFieldNotSpecified          = errors.New("error: field cannot be nil because of the non-nil field value requirement")
	ErrOnlyStringValue            = errors.New("error: value must be a string")
	ErrStringLengthIsNotSatisfied = errors.New("error: string's length is beyond the limit of min value or max value")
)

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
