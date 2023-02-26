package validation

import (
	"errors"
	"testing"
)

func TestValidateStructInformation(t *testing.T) {
	t.Run("NotOnStruct", func(t *testing.T) {
		var temp int = 5

		if err := ValidateStructInformative(temp); !errors.Is(err, ErrNotValidatable) {
			t.Errorf("Expected err = %v, recived err = %v", ErrNotValidatable, err)
		}
	})

	t.Run("NotPointerStruct", func(t *testing.T) {
		v := 5
		var temp *int = &v

		if err := ValidateStructInformative(temp); !errors.Is(err, ErrNotValidatable) {
			t.Errorf("Expected err = %v, recived err = %v", ErrNotValidatable, err)
		}
	})

	t.Run("NilStruct", func(t *testing.T) {
		var myType *MyType

		if err := ValidateStructInformative(myType); !errors.Is(err, ErrNilStruct) {
			t.Errorf("Expected err = %v, recived err = %v", ErrNilStruct, err)
		}
	})

	t.Run("NotValidStruct", func(t *testing.T) {
		myType := &MyType{10, "test", nil}

		if err := ValidateStructInformative(myType, Field(&myType.IntField, UnsignedInt)); !errors.Is(err, ErrStructNotValid) {
			t.Errorf("expected: %v, recived: %v", ErrStructNotValid, err)
		}
	})

	t.Run("NotValidStruct", func(t *testing.T) {
		myType := &MyType{10, "test", nil}

		if err := ValidateStructInformative(myType, Field(&myType.StringField, MatchRequired{"[a-z0-9]"})); !errors.Is(err, nil) {
			t.Errorf("expected: %v, recived: %v", nil, err)
		}
	})
}
