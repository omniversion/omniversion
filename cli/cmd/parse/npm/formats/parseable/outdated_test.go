package parseable

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseOutdatedOutput_NotInstalled(t *testing.T) {
	vector := ":underscore@1.13.3:MISSING:underscore@1.13.3:test"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.OutdatedCommand, verb)
	assert.Equal(t, formats.ParseableFormat, format)

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
	assert.Zero(t, len(item.Installations))
	assert.True(t, item.Missing)
}

func TestParseOutdatedOutput(t *testing.T) {
	vector := "/Users/testor/Documents/Repositories/test/node_modules/async:async@2.1.1:async@2.1.1:async@3.2.3:test\n/Users/testor/Documents/Repositories/test/node_modules/moment:moment@2.21.0:moment@2.21.0:moment@2.29.3:test\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.OutdatedCommand, verb)
	assert.Equal(t, formats.ParseableFormat, format)

	result, err := ParseOutdatedOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	item := result[0]
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, "2.1.1", item.Wanted)
	assert.Equal(t, "2.1.1", item.Current)
	assert.Equal(t, "3.2.3", item.Latest)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/Users/testor/Documents/Repositories/test/node_modules/async", item.Installations[0].Location)

	item = result[1]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, "2.21.0", item.Wanted)
	assert.Equal(t, "2.21.0", item.Current)
	assert.Equal(t, "2.29.3", item.Latest)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/Users/testor/Documents/Repositories/test/node_modules/moment", item.Installations[0].Location)
}

func TestParseOutdatedOutput_WithMissing(t *testing.T) {
	vector := "/srv/foobar/releases/20220420102847/frontend:test1@10.0.0:MISSING:test1@10.0.1\n/srv/foobar/releases/20220420102847/frontend:test2@10.0.0:MISSING:test2@10.0.0\n/srv/foobar/releases/20220420102847/frontend:@test/test3@10.0.0:MISSING:@test/test3@10.0.1\n/srv/foobar/releases/20220420102847/frontend:test4@13.3.4:MISSING:test4@13.3.4\n/srv/foobar/releases/20220420102847/frontend:test6@13.3.3:MISSING:test6@13.3.3\n/srv/foobar/releases/20220420102847/frontend:test5@13.3.4:MISSING:test5@13.3.4\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.OutdatedCommand, verb)
	assert.Equal(t, formats.ParseableFormat, format)

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := ParseOutdatedOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 6, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "test1", item.Name)
	assert.Equal(t, "10.0.0", item.Wanted)
	assert.Equal(t, "10.0.1", item.Latest)
	assert.Equal(t, 0, len(item.Installations))

	assert.Equal(t, "test2", result[1].Name)
	assert.Equal(t, "@test/test3", result[2].Name)
	assert.Equal(t, "test4", result[3].Name)
	assert.Equal(t, "test6", result[4].Name)
	assert.Equal(t, "test5", result[5].Name)
}
