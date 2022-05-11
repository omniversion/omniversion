package json

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseListOutput_NotInstalled(t *testing.T) {
	vector := "{\n  \"version\": \"1.0.0\",\n  \"name\": \"test\",\n  \"problems\": [\n    \"missing: async@2.1.1, required by test@1.0.0\",\n    \"missing: moment@2.21.0, required by test@1.0.0\",\n    \"missing: underscore@^1.0.3, required by test@1.0.0\"\n  ],\n  \"dependencies\": {\n    \"async\": {\n      \"required\": \"2.1.1\",\n      \"missing\": true,\n      \"problems\": [\n        \"missing: async@2.1.1, required by test@1.0.0\"\n      ]\n    },\n    \"moment\": {\n      \"missing\": true,\n      \"problems\": [\n        \"missing: moment@2.21.0, required by test@1.0.0\"\n      ]\n    },\n    \"underscore\": {\n      \"missing\": true,\n      \"problems\": [\n        \"missing: underscore@^1.0.3, required by test@1.0.0\"\n      ]\n    }\n  }\n}\nnpm ERR! code ELSPROBLEMS\nnpm ERR! missing: async@2.1.1, required by test@1.0.0\nnpm ERR! missing: moment@2.21.0, required by test@1.0.0\nnpm ERR! missing: underscore@^1.0.3, required by test@1.0.0\n{\n  \"error\": {\n    \"code\": \"ELSPROBLEMS\",\n    \"summary\": \"missing: async@2.1.1, required by test@1.0.0\\nmissing: moment@2.21.0, required by test@1.0.0\\nmissing: underscore@^1.0.3, required by test@1.0.0\",\n    \"detail\": \"\"\n  }\n}\n\nnpm ERR! A complete log of this run can be found in:\nnpm ERR!     /Users/testor/.npm/_logs/2022-05-09T10_30_40_760Z-debug.log\n"

	strippedInput, stderrOutput, _ := stderr.Strip(vector)
	verb, format := formats.DetectVerbAndFormat(strippedInput, stderrOutput)
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.JsonFormat, format)

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := ParseListOutput(strippedInput, stderrOutput)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "2.1.1", item.Wanted)
	assert.True(t, *item.Missing)

	item = result[1]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "2.21.0", item.Wanted)
	assert.True(t, *item.Missing)

	item = result[2]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "1.0.0", item.Current)
	assert.Equal(t, "", item.Wanted)

	item = result[3]
	assert.Equal(t, "underscore", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "^1.0.3", item.Wanted)
	assert.True(t, *item.Missing)
}

func TestParseListOutput(t *testing.T) {
	vector := "{\n  \"version\": \"1.0.0\",\n  \"name\": \"test\",\n  \"dependencies\": {\n    \"async\": {\n      \"version\": \"2.1.1\",\n      \"resolved\": \"https://registry.npmjs.org/async/-/async-2.1.1.tgz\",\n      \"dependencies\": {\n        \"lodash\": {\n          \"version\": \"4.17.21\",\n          \"resolved\": \"https://registry.npmjs.org/lodash/-/lodash-4.17.21.tgz\"\n        }\n      }\n    },\n    \"moment\": {\n      \"version\": \"2.21.0\",\n      \"resolved\": \"https://registry.npmjs.org/moment/-/moment-2.21.0.tgz\"\n    },\n    \"underscore\": {\n      \"version\": \"1.13.3\",\n      \"resolved\": \"https://registry.npmjs.org/underscore/-/underscore-1.13.3.tgz\"\n    }\n  }\n}\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.JsonFormat, format)

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := ParseListOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, "2.1.1", item.Current)
	assert.Equal(t, "2.1.1", item.Wanted)
	assert.Equal(t, []string{"lodash"}, item.Dependencies)

	item = result[1]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, "2.21.0", item.Current)
	assert.Equal(t, "2.21.0", item.Wanted)
	assert.Nil(t, item.Dependencies)

	item = result[2]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "1.0.0", item.Current)
	assert.Equal(t, "", item.Wanted)
	assert.Nil(t, item.Dependencies)

	item = result[3]
	assert.Equal(t, "underscore", item.Name)
	assert.Equal(t, "1.13.3", item.Current)
	assert.Equal(t, "1.13.3", item.Wanted)
	assert.Nil(t, item.Dependencies)
}

func TestParseListOutput_WithProblems(t *testing.T) {
	vector := "{\n\t\"version\": \"3.12.0\",\n\t\"name\": \"foobar\",\n\t\"problems\": [\n\t\t\"extraneous: __ngcc_entry_points__.json@ /Users/testortestington/Documents/Repositories/foobar/frontend/node_modules/__ngcc_entry_points__.json\"\n\t],\n\t\"dependencies\": {\n\t\t\"__ngcc_entry_points__.json\": {\n\t\t\t\"extraneous\": true,\n\t\t\t\"problems\": [\n\t\t\t\t\"extraneous: __ngcc_entry_points__.json@ /Users/test/testortestington/Repositories/foobar/frontend/node_modules/__ngcc_entry_points__.json\"\n\t\t\t]\n\t\t},\n\t\t\"test\": {\n\t\t\t\"version\": \"0.1301.2\",\n\t\t\t\"resolved\": \"https://registry.npmjs.org/test/test/-/test-0.1301.2.tgz\"\n\t\t}\n\t}\n}"

	strippedInput, stderrOutput, _ := stderr.Strip(vector)
	verb, format := formats.DetectVerbAndFormat(strippedInput, stderrOutput)
	assert.Equal(t, formats.ListCommand, verb)
	assert.Equal(t, formats.JsonFormat, format)

	result, err := ParseListOutput(strippedInput, stderrOutput)

	assert.Nil(t, err)
	assert.Equal(t, 3, len(result))
	item := result[0]
	assert.Equal(t, "__ngcc_entry_points__.json", item.Name)
	assert.Equal(t, "", item.PackageManager)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, "", item.Current)
	assert.True(t, *item.Extraneous)

	item = result[1]
	assert.Equal(t, "foobar", item.Name)
	assert.Equal(t, "3.12.0", item.Current)

	item = result[2]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "0.1301.2", item.Wanted)
	assert.Equal(t, "0.1301.2", item.Current)
}

func TestParseListOutput_InvalidJson(t *testing.T) {
	vector := "{\n\t\"version\":"

	result, err := ParseListOutput(vector, stderr.Output{})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
	assert.Zero(t, len(result))
}
