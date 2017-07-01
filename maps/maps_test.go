package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mapData map[string]string

func init() {
	mapData = make(map[string]string)
	mapData["valid-bool"] = "true"
	mapData["invalid-bool"] = "abcd"

	mapData["valid-float"] = "3.14"
	mapData["invalid-float"] = "3.1.4"

	mapData["valid-int"] = "12345"
	mapData["invalid-int"] = "abcd"

	mapData["key"] = "value"
}

func TestGetBoolean(t *testing.T) {
	assert.True(t, GetBoolean(mapData, "valid-bool", false))
	assert.False(t, GetBoolean(mapData, "invalid-bool", false))
}

func TestGetFloat64(t *testing.T) {
	assert.Equal(t, float64(3.14), GetFloat64(mapData, "valid-float", float64(1.23)))
	assert.Equal(t, float64(1.23), GetFloat64(mapData, "invalid-float", float64(1.23)))
}

func TestGetFloat(t *testing.T) {
	assert.Equal(t, 3.14, GetFloat64(mapData, "valid-float", 1.23))
	assert.Equal(t, 1.23, GetFloat64(mapData, "invalid-float", 1.23))
}

func TestGetString(t *testing.T) {
	assert.Equal(t, "value", GetString(mapData, "key", "default"))
	assert.Equal(t, "default", GetString(mapData, "no-key", "default"))
}

func TestGetInt64(t *testing.T) {
	assert.Equal(t, int64(12345), GetInt64(mapData, "valid-int", int64(123)))
	assert.Equal(t, int64(123), GetInt64(mapData, "invalid-int", int64(123)))
}

func TestGetInt(t *testing.T) {
	assert.Equal(t, 12345, GetInt(mapData, "valid-int", 123))
	assert.Equal(t, 123, GetInt(mapData, "invalid-int", 123))
}

func TestContains(t *testing.T) {
	assert.Equal(t, true, Contains(mapData, "key"))
}
