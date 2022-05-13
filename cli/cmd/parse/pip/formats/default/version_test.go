package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/formats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseVersionOutput(t *testing.T) {
	vector := "pip 22.0.4 from /Users/testor/Documents/Repositories/omniversion/python/env/lib/python3.10/site-packages/pip (python 3.10)\n"

	verb, format := formats.DetectVerbAndFormat(vector)
	assert.Equal(t, formats.VersionCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	result, err := ParseVersionOutput(vector)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	pip := result[0]
	assert.Equal(t, "pip", pip.Name)
	assert.Equal(t, "pip", pip.PackageManager)
	assert.Equal(t, "22.0.4", pip.Current)
	assert.Equal(t, 1, len(pip.Installations))
	assert.Equal(t, "/Users/testor/Documents/Repositories/omniversion/python/env/lib/python3.10/site-packages/pip", pip.Installations[0].Location)
	assert.Equal(t, "22.0.4", pip.Installations[0].Version)

	python := result[1]
	assert.Equal(t, "python", python.Name)
	assert.Equal(t, "pip", python.PackageManager)
	assert.Equal(t, "3.10", python.Current)
	assert.Equal(t, 1, len(python.Installations))
	assert.Equal(t, "", python.Installations[0].Location)
	assert.Equal(t, "3.10", python.Installations[0].Version)
}

func TestParseVersionOutput_Invalid(t *testing.T) {
	vector := "test"

	result, err := ParseVersionOutput(vector)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid version data")
	assert.Zero(t, len(result))
}
