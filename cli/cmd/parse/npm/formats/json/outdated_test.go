package json

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseOutdatedOutput_NotInstalled(t *testing.T) {
	vector := "{\n  \"underscore\": {\n    \"wanted\": \"1.13.3\",\n    \"latest\": \"1.13.3\",\n    \"dependent\": \"test\"\n  }\n}\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.OutdatedCommand, verb)
	assert.Equal(t, formats.JsonFormat, format)

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := ParseOutdatedOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "underscore", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "1.13.3", item.Wanted)
	assert.Equal(t, "1.13.3", item.Latest)
}

func TestParseOutdatedOutput(t *testing.T) {
	vector := "{\n  \"async\": {\n    \"current\": \"2.1.1\",\n    \"wanted\": \"2.1.1\",\n    \"latest\": \"3.2.3\",\n    \"dependent\": \"test\",\n    \"location\": \"/Users/testor/Documents/Repositories/test/node_modules/async\"\n  },\n  \"moment\": {\n    \"current\": \"2.21.0\",\n    \"wanted\": \"2.21.0\",\n    \"latest\": \"2.29.3\",\n    \"dependent\": \"test\",\n    \"location\": \"/Users/testor/Documents/Repositories/test/node_modules/moment\"\n  }\n}\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.OutdatedCommand, verb)
	assert.Equal(t, formats.JsonFormat, format)

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := ParseOutdatedOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, "2.1.1", item.Current)
	assert.Equal(t, "2.1.1", item.Wanted)
	assert.Equal(t, "3.2.3", item.Latest)

	item = result[1]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, "2.21.0", item.Current)
	assert.Equal(t, "2.21.0", item.Wanted)
	assert.Equal(t, "2.29.3", item.Latest)
}

func TestParseOutdatedOutput_InvalidJson(t *testing.T) {
	vector := "{\n\t\"test\": []\n}"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.OutdatedCommand, verb)
	assert.Equal(t, formats.JsonFormat, format)

	result, err := ParseOutdatedOutput(vector, stderr.Output{})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unable to interpret this input")
	assert.Zero(t, len(result))
}
