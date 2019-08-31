package graphson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseVertex(t *testing.T) {
	vertex, err := parseVertex([]byte(vertex30))
	assert.Nil(t, err)
	assert.Equal(t, int64(1), vertex.ID)
	assert.Equal(t, "person", vertex.Label)

	assert.NotEmpty(t, vertex.Properties["name"])
	assert.NotEmpty(t, vertex.Properties["location"])
	assert.Equal(t, "marko", vertex.Properties["name"][0].Value)
	assert.Equal(t, "san diego", vertex.Properties["location"][0].Value)
}

func TestParseVertexProperty(t *testing.T) {
	property, err := parseVertexProperty([]byte(vertexProperty30))
	assert.Nil(t, err)
	assert.Equal(t, int64(6), property.ID)
	assert.Equal(t, "san diego", property.Value)
	assert.Equal(t, "location", property.Label)

	p, ok := property.Properties["startTime"]
	assert.True(t, ok)
	assert.Equal(t, int64(1997), p)
}
