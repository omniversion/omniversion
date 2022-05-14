package nvm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNvmrcFile(t *testing.T) {
	vector := "v14.8.0"

	assert.False(t, isVersionOutput(vector))
	assert.True(t, isNvmrcFile(vector))

	result, err := parseNvmrcFile(vector)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "node", result[0].Name)
	assert.Equal(t, "v14.8.0", result[0].Wanted)
}
