package graphson3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEdge(t *testing.T) {
	g := GraphSONv3Parser{}
	edge, err := g.ParseEdge([]byte(edge30))
	assert.Nil(t, err)
	assert.Equal(t, int64(13), edge.ID)
	assert.Equal(t, "develops", edge.Label)

	assert.Equal(t, "person", edge.OutVLabel)
	assert.Equal(t, int64(1), edge.OutV)

	assert.Equal(t, "software", edge.InVLabel)
	assert.Equal(t, int64(10), edge.InV)

	assert.Equal(t, 2009, edge.Properties["since"].Value.Value)
}

func TestParseProperty(t *testing.T) {
	g := GraphSONv3Parser{}
	property, err := g.ParseProperty([]byte(property30))
	assert.Nil(t, err)
	assert.Equal(t, 2009, property.Value.Value)
	assert.Equal(t, "since", property.Key)
}
