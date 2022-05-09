package parseable

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseListOutput_NotInstalled(t *testing.T) {
	vector := "/Users/testor/Documents/Repositories/test\nnpm ERR! code ELSPROBLEMS\nnpm ERR! missing: async@2.1.1, required by test@1.0.0\nnpm ERR! missing: moment@2.21.0, required by test@1.0.0\nnpm ERR! missing: underscore@^1.0.3, required by test@1.0.0\n\nnpm ERR! A complete log of this run can be found in:\nnpm ERR!     /Users/testor/.npm/_logs/2022-05-09T07_13_37_978Z-debug.log\n"

	strippedInput, stderrOutput, _ := stderr.Strip(vector)
	verb, format := formats.DetectVerbAndFormat(strippedInput, stderrOutput)
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.ParseableFormat, format)

	result, err := ParseListOutput(strippedInput, stderrOutput)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))

	item := result[0]
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "2.1.1", item.Wanted)
	assert.Zero(t, len(item.Installations))
	assert.True(t, *item.Missing)

	item = result[1]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "2.21.0", item.Wanted)
	assert.Zero(t, len(item.Installations))

	item = result[2]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/Users/testor/Documents/Repositories/test", item.Installations[0].Location)

	item = result[3]
	assert.Equal(t, "underscore", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "^1.0.3", item.Wanted)
	assert.Zero(t, len(item.Installations))

}

func TestParseListOutput(t *testing.T) {
	vector := "/Users/testor/Documents/Repositories/test\n/Users/testor/Documents/Repositories/test/node_modules/async\n/Users/testor/Documents/Repositories/test/node_modules/moment\n/Users/testor/Documents/Repositories/test/node_modules/underscore\n/Users/testor/Documents/Repositories/test/node_modules/lodash\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.ParseableFormat, format)

	result, err := ParseListOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 5, len(result))

	item := result[0]
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/Users/testor/Documents/Repositories/test/node_modules/async", item.Installations[0].Location)

	item = result[1]
	assert.Equal(t, "lodash", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/Users/testor/Documents/Repositories/test/node_modules/lodash", item.Installations[0].Location)

	item = result[2]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/Users/testor/Documents/Repositories/test/node_modules/moment", item.Installations[0].Location)

	item = result[3]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/Users/testor/Documents/Repositories/test", item.Installations[0].Location)

	item = result[4]
	assert.Equal(t, "underscore", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/Users/testor/Documents/Repositories/test/node_modules/underscore", item.Installations[0].Location)
}
