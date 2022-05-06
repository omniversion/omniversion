package npm

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnknownProblemKind(t *testing.T) {
	vector := `
npm ERR! covfefe: name@version
`

	result, err := parseNpmOutput(vector)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unknown npm problem kind")
	assert.Equal(t, 1, len(result))
}

func TestLockfileVersionMismatchWarning(t *testing.T) {
	vector := `npm WARN read-shrinkwrap This version of npm is compatible with lockfileVersion@1, but package-lock.json was generated for lockfileVersion@2. I'll try to do my best with it!
/srv/foobar/releases/20220420102847/frontend`
	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(result))
}

func TestMissingWarning(t *testing.T) {
	vector := `npm ERR! missing: test@1.5.8, required by foo/bar@13.1.1
npm ERR! missing: test2@2.3.1, required by @some/other@13.1.1`
	previousInjectValue := shared.InjectPackageManager
	shared.InjectPackageManager = true
	result, err := parseNpmOutput(vector)
	shared.InjectPackageManager = previousInjectValue

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "1.5.8", item.Wanted)

	item = result[1]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "test2", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "2.3.1", item.Wanted)
}

func TestExtraneousWarning(t *testing.T) {
	vector := `npm ERR! extraneous: foo/bar@2.0.0-beta.16 /srv/foobar/releases/20220420102847/frontend/node_modules/test/parent1
npm ERR! extraneous: foo/bar2@2.0.0-beta.13 /srv/foobar/releases/20220420102847/frontend/node_modules/test/parent2`
	previousInjectValue := shared.InjectPackageManager
	shared.InjectPackageManager = true
	result, err := parseNpmOutput(vector)
	shared.InjectPackageManager = previousInjectValue

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "foo/bar", item.Name)
	assert.Equal(t, "2.0.0-beta.16", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/srv/foobar/releases/20220420102847/frontend/node_modules/test/parent1", item.Installations[0].Location)

	item = result[1]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "foo/bar2", item.Name)
	assert.Equal(t, "2.0.0-beta.13", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/srv/foobar/releases/20220420102847/frontend/node_modules/test/parent2", item.Installations[0].Location)
}

func TestParseNpmOutputWithAt(t *testing.T) {
	vector := `npm ERR! missing: @foo/bar@2.0.0-beta.13 /srv/foobar/releases/20220420102847/frontend/node_modules/test/parent`
	previousInjectValue := shared.InjectPackageManager
	shared.InjectPackageManager = true
	result, err := parseNpmOutput(vector)
	shared.InjectPackageManager = previousInjectValue

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "@foo/bar", item.Name)
	assert.Equal(t, "2.0.0-beta.13", item.Wanted)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "/srv/foobar/releases/20220420102847/frontend/node_modules/test/parent", item.Installations[0].Location)
}
