package nvm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseVersionOutput(t *testing.T) {
	vector := "0.35.3"

	assert.True(t, isVersionOutput(vector))

	result, err := parseNvmOutput(vector)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "nvm", result[0].Name)
	assert.Equal(t, "nvm", result[0].PackageManager)
	assert.Equal(t, 1, len(result[0].Installations))
}
