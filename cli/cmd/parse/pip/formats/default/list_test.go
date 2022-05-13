package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/formats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseListOutput(t *testing.T) {
	vector := "Package Version\n------- -------\ndocopt  0.6.2\nidlex   1.13\njedi    0.9.0"

	verb, format := formats.DetectVerbAndFormat(vector)
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	result, err := ParseListOutput(vector)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(result))
}

func TestParseListOutput_EditableMode(t *testing.T) {
	vector := "Package          Version  Editable project location\n---------------- -------- -------------------------------------\npip              21.2.4\npip-test-package 0.1.1    /home/you/.venv/src/pip-test-package\nsetuptools       57.4.0\nwheel            0.36.2"

	verb, format := formats.DetectVerbAndFormat(vector)
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	result, err := ParseListOutput(vector)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))

	assert.Equal(t, "pip-test-package", result[1].Name)
	assert.Equal(t, "0.1.1", result[1].Current)
	assert.Equal(t, "/home/you/.venv/src/pip-test-package", result[1].Installations[0].Location)
}

func TestParseListOutput_InvalidLine(t *testing.T) {
	vector := "Package          Version  Editable project location\n---------------- -------- -------------------------------------\ntest"

	result, err := ParseListOutput(vector)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to parse line")
	assert.Zero(t, len(result))
}
