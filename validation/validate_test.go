package validation

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("NotUnsignedInt value test", func(t *testing.T) {
		var value int = 20
		var rules RuleSlice = []Rule{UnsignedInt}

		if err := errors.Unwrap(Validate(&value, rules)); !errors.Is(err, ErrValueNotUnsignedInt) {
			t.Errorf("Expected err = %v, recived err = %v", ErrValueNotUnsignedInt, err)
		}
	})

	t.Run("UnsignedInt value test", func(t *testing.T) {
		var value uint64 = 20
		var rules RuleSlice = []Rule{UnsignedInt}

		if err := errors.Unwrap(Validate(&value, rules)); !errors.Is(err, nil) {
			t.Errorf("Expected err = %v, recived err = %v", nil, err)
		}
	})

	t.Run("OnlyStringForRegex test ", func(t *testing.T) {
		notstr := 12
		var rules RuleSlice = []Rule{MatchRequired{"[A-Z]"}}

		if err := Validate(&notstr, rules); !errors.Is(err, ErrOnlyStringValue) {
			t.Errorf("Expected err = %v, recived err = %v", ErrOnlyStringValue, err)
		}
	})
	t.Run("ValidRegexString test ", func(t *testing.T) {
		str := "TEST"
		var rules RuleSlice = []Rule{MatchRequired{"[A-Z]"}}

		if err := errors.Unwrap(Validate(&str, rules)); !errors.Is(err, nil) {
			t.Errorf("Expected err = %v, recived err = %v", nil, err)
		}
	})

	t.Run("NotValidRegexString test ", func(t *testing.T) {
		str := "test"
		var rules RuleSlice = []Rule{MatchRequired{"[A-Z]"}}

		if err := Validate(&str, rules); !errors.Is(err, ErrRegexNotSatisfied) {
			t.Errorf("Expected err = %v, recived err = %v", ErrRegexNotSatisfied, err)
		}
	})

	t.Run("CorrectLength test ", func(t *testing.T) {
		str := "test"
		var rules RuleSlice = []Rule{Length{10, 20}}

		if err := Validate(&str, rules); !errors.Is(err, ErrStringLengthIsNotSatisfied) {
			t.Errorf("Expected err = %v, recived err = %v", ErrStringLengthIsNotSatisfied, err)
		}
	})

	t.Run("IncorrectLength test ", func(t *testing.T) {
		str := "test"
		var rules RuleSlice = []Rule{Length{2, 10}}

		if err := Validate(&str, rules); !errors.Is(err, nil) {
			t.Errorf("Expected err = %v, recived err = %v", nil, err)
		}
	})

	t.Run("CorrectRequired test ", func(t *testing.T) {
		value := 10
		var definedVar *int = &value

		var rules RuleSlice = []Rule{Required}

		if err := Validate(&definedVar, rules); !errors.Is(err, nil) {
			t.Errorf("Expected err = %v, recived err = %v", nil, err)
		}
	})

	t.Run("IncorrectRequired test ", func(t *testing.T) {
		var definedVar *int

		var rules RuleSlice = []Rule{Required}

		if err := Validate(&definedVar, rules); !errors.Is(err, ErrFieldNotSpecified) {
			t.Errorf("Expected err = %v, recived err = %v", ErrFieldNotSpecified, err)
		}
	})

}
