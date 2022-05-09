package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseListOutput_NotInstalled(t *testing.T) {
	vector := "test@1.0.0 /Users/testor/Documents/Repositories/test\n├── UNMET DEPENDENCY async@2.1.1\n├── UNMET DEPENDENCY moment@2.21.0\n└── UNMET DEPENDENCY underscore@^1.0.3\n\nnpm ERR! code ELSPROBLEMS\nnpm ERR! missing: async@2.1.1, required by test@1.0.0\nnpm ERR! missing: moment@2.21.0, required by test@1.0.0\nnpm ERR! missing: underscore@^1.0.3, required by test@1.0.0\n\nnpm ERR! A complete log of this run can be found in:\nnpm ERR!     /Users/testor/.npm/_logs/2022-05-09T07_11_49_925Z-debug.log\n"

	strippedInput, stderrOutput, _ := stderr.Strip(vector)
	verb, format := formats.DetectVerbAndFormat(strippedInput, stderrOutput)
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	result, err := ParseListOutput(strippedInput, stderrOutput)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))

	item := result[0]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "1.0.0", item.Wanted)
	assert.Equal(t, "1.0.0", item.Current)

	item = result[1]
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, "2.1.1", item.Wanted)
	assert.Equal(t, "", item.Current)

	item = result[2]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, "2.21.0", item.Wanted)
	assert.Equal(t, "", item.Current)

	item = result[3]
	assert.Equal(t, "underscore", item.Name)
	assert.Equal(t, "^1.0.3", item.Wanted)
	assert.Equal(t, "", item.Current)
}

func TestParseListOutput(t *testing.T) {
	vector := "test@1.0.0 /Users/testor/Documents/Repositories/test\n├─┬ async@2.1.1\n│ └── lodash@4.17.21\n├── moment@2.21.0\n└── underscore@1.13.3\n\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	result, err := ParseListOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 5, len(result))

	item := result[0]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "1.0.0", item.Wanted)
	assert.Equal(t, "1.0.0", item.Current)
	assert.Equal(t, []string{"async", "moment", "underscore"}, item.Dependencies)

	item = result[1]
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, "2.1.1", item.Wanted)
	assert.Equal(t, "2.1.1", item.Current)
	assert.Equal(t, []string{"lodash"}, item.Dependencies)

	item = result[2]
	assert.Equal(t, "lodash", item.Name)
	assert.Equal(t, "4.17.21", item.Wanted)
	assert.Equal(t, "4.17.21", item.Current)
	assert.Nil(t, item.Dependencies)

	item = result[3]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, "2.21.0", item.Wanted)
	assert.Equal(t, "2.21.0", item.Current)
	assert.Nil(t, item.Dependencies)

	item = result[4]
	assert.Equal(t, "underscore", item.Name)
	assert.Equal(t, "1.13.3", item.Wanted)
	assert.Equal(t, "1.13.3", item.Current)
	assert.Nil(t, item.Dependencies)
}
