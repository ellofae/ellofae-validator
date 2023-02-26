package validation

import (
	"errors"
	"testing"
)

type MyType struct {
	IntField    int
	StringField string
	PtrField    *int
}

func TestValidateStruct(t *testing.T) {
	t.Run("RequiredTest", func(t *testing.T) {
		myType := &MyType{
			IntField: 225,
		}

		if err := ValidateStruct(myType, Field(&myType.IntField, UnsignedInt)); !errors.Is(err, ErrValueNotUnsignedInt) {
			t.Errorf("expected: %v, recived %v", ErrValueNotUnsignedInt, err)
		}
	})

	t.Run("StringLengthTEst", func(t *testing.T) {
		myType := &MyType{
			StringField: "test field",
		}

		if err := ValidateStruct(myType, Field(&myType.StringField, Length{20, 40})); !errors.Is(err, ErrStringLengthIsNotSatisfied) {
			t.Errorf("expected: %v, recived %v", ErrStringLengthIsNotSatisfied, err)
		}
	})

	t.Run("StringLengthTestTwo", func(t *testing.T) {
		myType := &MyType{
			StringField: "test field",
		}

		if err := ValidateStruct(myType, Field(&myType.StringField, Length{5, 40})); !errors.Is(err, nil) {
			t.Errorf("expected: %v, recived %v", nil, err)
		}
	})

	t.Run("StringRegexTest", func(t *testing.T) {
		myType := &MyType{
			StringField: "test",
		}

		if err := ValidateStruct(myType, Field(&myType.StringField, MatchRequired{"[a-z]"})); !errors.Is(err, nil) {
			t.Errorf("expected: %v, recived %v", nil, err)
		}
	})

	t.Run("StringRegexTestTwo", func(t *testing.T) {
		myType := &MyType{
			StringField: "test",
		}

		if err := ValidateStruct(myType, Field(&myType.StringField, MatchRequired{"[A-Z0-9]"})); !errors.Is(err, ErrRegexNotSatisfied) {
			t.Errorf("expected: %v, recived %v", ErrRegexNotSatisfied, err)
		}
	})

	t.Run("RequiredNotNilTest", func(t *testing.T) {
		myType := &MyType{
			PtrField: nil,
		}

		if err := ValidateStruct(myType, Field(&myType.PtrField, Required)); !errors.Is(err, ErrFieldNotSpecified) {
			t.Errorf("expected: %v, recived %v", ErrFieldNotSpecified, err)
		}
	})

	t.Run("RequiredNotNilTwo", func(t *testing.T) {
		value := 5
		myType := &MyType{
			PtrField: &value,
		}

		if err := ValidateStruct(myType, Field(&myType.PtrField, Required)); !errors.Is(err, nil) {
			t.Errorf("expected: %v, recived %v", nil, err)
		}
	})
}

func TestInValidateStructValidInput(t *testing.T) {
	t.Run("NotDefinedStrctTest", func(t *testing.T) {
		var myType *MyType

		if err := ValidateStruct(myType); !errors.Is(err, ErrNilStruct) {
			t.Errorf("expected: %v, recived: %v", ErrNilStruct, err)
		}
	})

	t.Run("DefinedStrctTest", func(t *testing.T) {
		var myType *MyType = &MyType{}

		if err := ValidateStruct(myType); !errors.Is(err, nil) {
			t.Errorf("expected: %v, recived: %v", nil, err)
		}
	})

	t.Run("NotValidableStruct", func(t *testing.T) {
		var myType MyType

		if err := ValidateStruct(myType); !errors.Is(err, ErrNotValidatable) {
			t.Errorf("expected: %v, recived: %v", ErrNotValidatable, err)
		}
	})

	t.Run("NotValidableStruct", func(t *testing.T) {
		var myType int

		if err := ValidateStruct(myType); !errors.Is(err, ErrNotValidatable) {
			t.Errorf("expected: %v, recived: %v", ErrNotValidatable, err)
		}
	})

	t.Run("ValidStruct", func(t *testing.T) {
		var myType int

		if err := ValidateStruct(myType); !errors.Is(err, ErrNotValidatable) {
			t.Errorf("expected: %v, recived: %v", ErrNotValidatable, err)
		}
	})

	t.Run("NotValidStruct", func(t *testing.T) {
		var myType *MyType

		if err := ValidateStruct(myType); !errors.Is(err, ErrNilStruct) {
			t.Errorf("expected: %v, recived: %v", ErrNilStruct, err)
		}
	})
}
