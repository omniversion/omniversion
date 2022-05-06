package npm

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseListParseableOutput(t *testing.T) {
	vector :=
		`/Users/testortestington/Documents/Repositories/foobar/frontend
/Users/testortestington/Documents/Repositories/foobar/frontend/node_modules/__ngcc_entry_points__.json
/Users/testortestington/Documents/Repositories/foobar/frontend/node_modules/@angular-devkit/architect
`

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 3, len(result))

	item := result[0]
	assert.Equal(t, "", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/Users/testortestington/Documents/Repositories/foobar/frontend", item.Installations[0].Location)

	item = result[1]
	assert.Equal(t, "", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/Users/testortestington/Documents/Repositories/foobar/frontend/node_modules/__ngcc_entry_points__.json", item.Installations[0].Location)

	item = result[2]
	assert.Equal(t, "", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/Users/testortestington/Documents/Repositories/foobar/frontend/node_modules/@angular-devkit/architect", item.Installations[0].Location)
}

func TestParseOutdatedParseableOutput(t *testing.T) {
	shared.InjectPackageManager = true
	vector := `/Users/testor/Documents/Repositories/foobar/frontend/node_modules/@angular-eslint/template-parser:@angular-eslint/template-parser@13.2.1:@angular-eslint/template-parser@13.0.1:@angular-eslint/template-parser@13.2.4:frontend`

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "@angular-eslint/template-parser", item.Name)
	assert.Equal(t, "13.0.1", item.Current)
	assert.Equal(t, "13.2.1", item.Wanted)
	assert.Equal(t, "13.2.4", item.Latest)
	assert.Equal(t, "13.0.1", item.Installations[0].Version)
	assert.Equal(t, "/Users/testor/Documents/Repositories/foobar/frontend/node_modules/@angular-eslint/template-parser", item.Installations[0].Location)
}

func TestParseOutdatedParseableWithMissing(t *testing.T) {
	vector := `/srv/foobar/releases/20220420102847/frontend:test1@10.0.0:MISSING:test1@10.0.1
/srv/foobar/releases/20220420102847/frontend:test2@10.0.0:MISSING:test2@10.0.0
/srv/foobar/releases/20220420102847/frontend:@test/test3@10.0.0:MISSING:@test/test3@10.0.1
/srv/foobar/releases/20220420102847/frontend:test4@13.3.4:MISSING:test4@13.3.4
/srv/foobar/releases/20220420102847/frontend:test6@13.3.3:MISSING:test6@13.3.3
/srv/foobar/releases/20220420102847/frontend:test5@13.3.4:MISSING:test5@13.3.4
`
	previousInjectValue := shared.InjectPackageManager
	shared.InjectPackageManager = true
	result, err := parseNpmOutput(vector)
	shared.InjectPackageManager = previousInjectValue

	assert.Nil(t, err)
	assert.Equal(t, 6, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "@test/test3", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "10.0.0", item.Wanted)
	assert.Equal(t, "10.0.1", item.Latest)
	assert.Equal(t, 0, len(item.Installations))

	assert.Equal(t, "test1", result[1].Name)
	assert.Equal(t, "test2", result[2].Name)
	assert.Equal(t, "test4", result[3].Name)
	assert.Equal(t, "test5", result[4].Name)
	assert.Equal(t, "test6", result[5].Name)
}
