package col

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOne(t *testing.T) {

	type Abc struct {
		FieldString string
		FieldInt    int
		FieldFloat  float64
	}

	var collection = []Abc{
		Abc{
			FieldString: "one",
			FieldInt:    1,
			FieldFloat:  1.1,
		},
		Abc{
			FieldString: "two",
			FieldInt:    2,
			FieldFloat:  2.2,
		},
		Abc{
			FieldString: "three",
			FieldInt:    3,
			FieldFloat:  3.3,
		},
	}

	assert.Equal(t, 1, FindOne(collection, "FieldString", "one").(Abc).FieldInt)

	assert.Equal(t, "two", FindOne(collection, "FieldInt", 2).(Abc).FieldString)

	assert.Equal(t, "three", FindOne(collection, "FieldFloat", 3.3).(Abc).FieldString)

	assert.Nil(t, FindOne(collection, "FieldDoesNotExist", "any"))

	assert.Nil(t, FindOne(collection, "FieldString", "value_does_not_match"))
}
