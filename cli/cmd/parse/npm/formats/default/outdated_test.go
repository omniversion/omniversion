package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseOutdatedOutput_NotInstalled(t *testing.T) {
	vector := "Package     Current  Wanted  Latest  Location  Depended by\nunderscore  MISSING  1.13.3  1.13.3  -         test\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.OutdatedCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	result, err := ParseOutdatedOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "underscore", item.Name)
	assert.Equal(t, "1.13.3", item.Wanted)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "1.13.3", item.Latest)
	assert.Zero(t, len(item.Installations))
}

func TestParseOutdatedOutput(t *testing.T) {
	vector := "Package  Current  Wanted  Latest  Location             Depended by\nasync      2.1.1   2.1.1   3.2.3  node_modules/async   test\nmoment    2.21.0  2.21.0  2.29.3  node_modules/moment  test\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.OutdatedCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	result, err := ParseOutdatedOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	item := result[0]
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, "2.1.1", item.Wanted)
	assert.Equal(t, "2.1.1", item.Current)
	assert.Equal(t, "3.2.3", item.Latest)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "2.1.1", item.Installations[0].Version)

	item = result[1]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, "2.21.0", item.Wanted)
	assert.Equal(t, "2.21.0", item.Current)
	assert.Equal(t, "2.29.3", item.Latest)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "2.21.0", item.Installations[0].Version)
}
