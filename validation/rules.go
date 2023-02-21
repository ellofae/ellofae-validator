package validation

import (
	"fmt"
	"reflect"
	"regexp"
)

type (
	// Type Rule which specifier a rule for validation
	Rule interface {
		Specifier(value interface{}) error
	}

	// Type RuleSlice for a slice of Rules
	RuleSlice []Rule
)

// Variables of rules that can be used to validate:
var (
	Required    = RequiredProperties{undefinedField: true}
	UnsignedInt = UnsignedIntProperties{unsignedInt: true}
)

// Types of rules that can be used to validate:
type (
	Length struct { // to specify the length range of string
		MinValue int
		MaxValue int
	}

	RequiredProperties struct { // Wrapper for a Required variable which specifier that field's value is not undefined(nil)
		undefinedField bool
	}

	MatchRequired struct {
		RegexToMatch string
	}

	UnsignedIntProperties struct {
		unsignedInt bool
	}
)

// Specifies RequiredProperties type as Rule interface
func (r RequiredProperties) Specifier(value interface{}) error {
	val := reflect.ValueOf(value).Elem()
	if val.Kind() == reflect.Ptr && val.IsNil() {
		if r.undefinedField {
			return ErrFieldNotSpecified
		}
	}

	return nil
}

// Specifies Length type as Rule interface
func (l Length) Specifier(value interface{}) error {
	val := reflect.ValueOf(value).Elem()

	if val.Kind() != reflect.String {
		return ErrOnlyStringValue
	}

	strLength := len(val.String())
	if strLength < l.MinValue || strLength > l.MaxValue {
		return fmt.Errorf("%w (current string's length: %d)", ErrStringLengthIsNotSatisfied, strLength)
	}

	return nil
}

// Specifies MatchRequired type as Rule interface
func (m MatchRequired) Specifier(value interface{}) error {
	val := reflect.ValueOf(value).Elem()
	if val.Kind() != reflect.String {
		return ErrOnlyStringValue
	}

	re, err := regexp.Compile(m.RegexToMatch)
	if err != nil {
		return err
	}

	if !re.MatchString(val.String()) {
		return ErrRegexNotSatisfied
	}

	return nil
}

// Add validation for specifing the value is unsigned int
func (u UnsignedIntProperties) Specifier(value interface{}) error {
	val := reflect.ValueOf(value).Elem().Kind()
	if val != reflect.Uint8 && val != reflect.Uint16 && val != reflect.Uint32 && val != reflect.Uint64 && val != reflect.Uint {
		return fmt.Errorf("%w (current field type: %s)", ErrValueNotUnsignedInt, reflect.ValueOf(value).Elem().Type())
	}

	return nil
}

// Specifies RuleSlice type as Rule interface
func (t RuleSlice) Specifier() error {
	return nil
}
