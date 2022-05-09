package json

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePackageJson(t *testing.T) {
	vector := "{\n  \"name\": \"test\",\n  \"version\": \"1.0.0\",\n  \"description\": \"Test app\",\n  \"main\": \"index.js\",\n  \"scripts\": {\n    \"test\": \"test\"\n  },\n  \"repository\": {\n    \"type\": \"git\",\n    \"url\": \"git+https://github.com/example.com/test.git\"\n  },\n  \"keywords\": [\n    \"test\"\n  ],\n  \"author\": \"Testor Testington\",\n  \"license\": \"ISC\",\n  \"bugs\": {\n    \"url\": \"https://github.com/example.com/test/issues\"\n  },\n  \"homepage\": \"https://github.com/example.com/test#readme\",\n  \"dependencies\": {\n    \"underscore\": \"^1.0.3\"\n  },\n  \"devDependencies\": {\n    \"async\": \"2.1.1\"\n  },\n  \"peerDependencies\": {\n    \"moment\": \"2.21.0\"\n  }\n}\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.PackageJsonFile, verb)
	assert.Equal(t, formats.JsonFormat, format)

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := ParsePackageJsonFile(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))

	item := result[0]
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "2.1.1", item.Wanted)

	assert.Equal(t, 0, len(item.Installations))

	item = result[1]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "2.21.0", item.Wanted)

	item = result[2]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "1.0.0", item.Current)
	assert.Equal(t, "", item.Wanted)

	item = result[3]
	assert.Equal(t, "underscore", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "^1.0.3", item.Wanted)
}
