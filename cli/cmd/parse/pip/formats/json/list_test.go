package json

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/formats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseListOutput(t *testing.T) {
	vector := "[{\"name\": \"colorama\", \"version\": \"0.3.7\"}, {\"name\": \"docopt\", \"version\": \"0.6.2\"}]\n"

	verb, format := formats.DetectVerbAndFormat(vector)
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.JsonFormat, format)

	result, err := ParseListOutput(vector)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
}

func TestParseListOutput_InvalidJson(t *testing.T) {
	vector := "{ test"

	result, err := ParseListOutput(vector)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unable to parse JSON")
	assert.Zero(t, len(result))
}
