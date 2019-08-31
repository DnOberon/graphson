package graphson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValuePair(t *testing.T) {
	valuePair, err := parseValuePair([]byte(valuePairSimple30))
	assert.Nil(t, err)
	assert.Equal(t, valuePair.Value, float64(1))
	assert.Equal(t, valuePair.ValueTypeName, "g:Int64")

}
