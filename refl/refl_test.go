package refl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComposedOf(t *testing.T) {
	type GrandFather struct {
		AnyField string
	}
	type Father struct {
		MoreField string
		GrandFather
	}
	type Son struct {
		OtherField string
		Father
	}
	type Uncle struct {
		MiscField string
	}

	// direct composition
	assert.True(t, ComposedOf(Son{}, Father{}))
	assert.True(t, ComposedOf(Father{}, GrandFather{}))

	// passing addresses
	assert.True(t, ComposedOf(&Son{}, &Father{}))

	// indirect composition
	assert.True(t, ComposedOf(Son{}, GrandFather{}))

	// no relationship
	assert.False(t, ComposedOf(Uncle{}, Son{}))
}

func TestNestedFields(t *testing.T) {
	type Father struct {
		FieldA string
	}
	type Son struct {
		FieldB string
		Father
	}

	assert.Equal(t, 2, len(NestedFields(Son{})))
	assert.Equal(t, 2, len(NestedFields(&Son{})))
}
