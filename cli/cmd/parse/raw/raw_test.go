package raw

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRawOutput_NameWithoutRegex(t *testing.T) {
	vector := "v1.2.3"

	name = "test"
	regex = ""
	result, err := parseRawOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "1.2.3", item.Current)
}

func TestParseRawOutput_NameWithRegex(t *testing.T) {
	vector := "version v1.2.3"

	name = "test"
	regex = "version (?P<version>.*)"
	result, err := parseRawOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "1.2.3", item.Current)
}

func TestParseRawOutput_InvalidRegex(t *testing.T) {
	vector := "version v1.2.3"

	name = "test"
	regex = "version (?P<version>.*"
	result, err := parseRawOutput(vector)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid regex")
	assert.Equal(t, 0, len(result))
}

func TestParseRawOutput_RegexWithName(t *testing.T) {
	vector := "test=v1.2.3"

	name = ""
	regex = `(?m)^(?P<name>.*)=(?P<version>.*)$`
	result, err := parseRawOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "1.2.3", item.Current)
}

func TestParseRawOutput_MultipleVersions(t *testing.T) {
	vector := "test1=v1.2.3\ntest2=v2.3.4\ntest3=v3.4.5"

	name = ""
	regex = `(?m)^(?P<name>\S*)=(?P<version>\S*)$`
	result, err := parseRawOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 3, len(result))

	item := result[0]
	assert.Equal(t, "test1", item.Name)
	assert.Equal(t, "1.2.3", item.Current)
}
