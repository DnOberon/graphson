package graphson3

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetParse(t *testing.T) {
	set, err := parseSet([]byte(set30))

	assert.Nil(t, err)
	assert.Len(t, set, 3)

	assert.Equal(t, 1, set[0])
	assert.Equal(t, "person", set[1])
	assert.Equal(t, true, set[2])
}

func TestSetInt32(t *testing.T) {
	out, err := parseInt32([]byte(integer30))

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(out).Kind(), reflect.Int)
}

func TestSetInt64(t *testing.T) {
	out, err := parseInt64([]byte(long30))

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(out).Kind(), reflect.Int64)
}

func TestSetFloat32(t *testing.T) {
	out, err := parseFloat32([]byte(double30))

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(out).Kind(), reflect.Float32)
}

func TestSetFloat64(t *testing.T) {
	out, err := parseFloat64([]byte(float30))

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(out).Kind(), reflect.Float64)
}
