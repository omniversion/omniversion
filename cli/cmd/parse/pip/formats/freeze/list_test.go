package freeze

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/formats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseListOutput(t *testing.T) {
	vector := "colorama==0.3.7\ndocopt==0.6.2\nidlex==1.13\njedi==0.9.0\n"

	verb, format := formats.DetectVerbAndFormat(vector)
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.FreezeFormat, format)

	result, err := ParseListOutput(vector)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))
}

func TestParseListOutput_InvalidLine(t *testing.T) {
	vector := "colorama==0.3.7\ndocopt==0.6.2\ntest\nidlex==1.13\njedi==0.9.0\n"

	verb, format := formats.DetectVerbAndFormat(vector)
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.FreezeFormat, format)

	result, err := ParseListOutput(vector)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to parse line")
	assert.Equal(t, 4, len(result))
}

func TestParseListOutput_FixMeLine(t *testing.T) {
	vector := "argparse==1.2.1\n## FIXME: could not find svn URL in dependency_links for this package:\ndistribute==0.6.24dev-r0\n"

	verb, format := formats.DetectVerbAndFormat(vector)
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.FreezeFormat, format)

	result, err := ParseListOutput(vector)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
}
