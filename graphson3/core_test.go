package graphson3

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetParse(t *testing.T) {
	g := GraphSONv3Parser{}
	set, err := g.parseSet([]byte(set30))

	assert.Nil(t, err)
	assert.Len(t, set, 3)

	assert.Equal(t, 1, set[0].Value)
	assert.Equal(t, "person", set[1].Value)
	assert.Equal(t, true, set[2].Value)
}

func TestMapParse(t *testing.T) {
	g := GraphSONv3Parser{}
	m, err := g.parseFlatMap([]byte(map30))

	assert.Nil(t, err)
	assert.NotEmpty(t, m)
}

func TestClassParse(t *testing.T) {
	g := GraphSONv3Parser{}
	out, err := g.parseClass([]byte(class30))

	assert.Nil(t, err)
	assert.Equal(t, "java.io.File", out)
}

func TestUUIDParse(t *testing.T) {
	g := GraphSONv3Parser{}
	out, err := g.parseUUID([]byte(uuid30))

	assert.Nil(t, err)
	assert.Equal(t, "41d2e28a-20a4-4ab0-b379-d810dede3786", out)
}

func TestInt32Parse(t *testing.T) {
	g := GraphSONv3Parser{}
	out, err := g.parseInt32([]byte(integer30))

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(out).Kind(), reflect.Int)
}

func TestInt64Parse(t *testing.T) {
	g := GraphSONv3Parser{}
	out, err := g.parseInt64([]byte(long30))

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(out).Kind(), reflect.Int64)
}

func TestFloat32Parse(t *testing.T) {
	g := GraphSONv3Parser{}
	out, err := g.parseFloat32([]byte(double30))

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(out).Kind(), reflect.Float32)
}

func TestFloat64Parse(t *testing.T) {
	g := GraphSONv3Parser{}
	out, err := g.parseFloat64([]byte(float30))

	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(out).Kind(), reflect.Float64)
}

func TestTimestampParse(t *testing.T) {
	g := GraphSONv3Parser{}
	out, err := g.parseTimestamp([]byte(timestamp30))

	assert.Nil(t, err)

	year, month, day := out.Date()
	assert.Equal(t, 1974, year)
	assert.Equal(t, time.September, month)
	assert.Equal(t, 11, day)

	out, err = g.parseTimestamp([]byte(date30))

	assert.Nil(t, err)

	year, month, day = out.Date()
	assert.Equal(t, 1974, year)
	assert.Equal(t, time.September, month)
	assert.Equal(t, 11, day)
}
