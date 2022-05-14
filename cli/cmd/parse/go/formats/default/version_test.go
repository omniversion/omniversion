package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/go/formats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseVersion(t *testing.T) {
	vector := "go version go1.17.5 darwin/amd64\n"

	verb, format := formats.DetectVerbAndFormat(vector)
	assert.Equal(t, formats.VersionCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	result, err := ParseVersionOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "go", item.Name)
	assert.Equal(t, "1.17.5", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "1.17.5", item.Installations[0].Version)
}
