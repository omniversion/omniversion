package json

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/omniversion/omniversion/cli/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePackageLockJson(t *testing.T) {
	vector := "{\n  \"name\": \"test\",\n  \"version\": \"1.0.0\",\n  \"lockfileVersion\": 2,\n  \"requires\": true,\n  \"packages\": {\n    \"\": {\n      \"name\": \"test\",\n      \"version\": \"1.0.0\",\n      \"license\": \"ISC\",\n      \"dependencies\": {\n        \"underscore\": \"^1.0.3\"\n      },\n      \"devDependencies\": {\n        \"async\": \"2.1.1\"\n      },\n      \"peerDependencies\": {\n        \"moment\": \"2.21.0\"\n      }\n    },\n    \"node_modules/async\": {\n      \"version\": \"2.1.1\",\n      \"resolved\": \"https://registry.npmjs.org/async/-/async-2.1.1.tgz\",\n      \"integrity\": \"sha1-4RttEAQ/IlTvthohFj2EDM3bjSg=\",\n      \"dev\": true,\n      \"dependencies\": {\n        \"lodash\": \"^4.14.0\"\n      }\n    },\n    \"node_modules/lodash\": {\n      \"version\": \"4.17.21\",\n      \"resolved\": \"https://registry.npmjs.org/lodash/-/lodash-4.17.21.tgz\",\n      \"integrity\": \"sha512-v2kDEe57lecTulaDIuNTPy3Ry4gLGJ6Z1O3vE1krgXZNrsQ+LFTGHVxVjcXPs17LhbZVGedAJv8XZ1tvj5FvSg==\",\n      \"dev\": true\n    },\n    \"node_modules/moment\": {\n      \"version\": \"2.21.0\",\n      \"resolved\": \"https://registry.npmjs.org/moment/-/moment-2.21.0.tgz\",\n      \"integrity\": \"sha512-TCZ36BjURTeFTM/CwRcViQlfkMvL1/vFISuNLO5GkcVm1+QHfbSiNqZuWeMFjj1/3+uAjXswgRk30j1kkLYJBQ==\",\n      \"peer\": true,\n      \"engines\": {\n        \"node\": \"*\"\n      }\n    },\n    \"node_modules/underscore\": {\n      \"version\": \"1.13.3\",\n      \"resolved\": \"https://registry.npmjs.org/underscore/-/underscore-1.13.3.tgz\",\n      \"integrity\": \"sha512-QvjkYpiD+dJJraRA8+dGAU4i7aBbb2s0S3jA45TFOvg2VgqvdCDd/3N6CqA8gluk1W91GLoXg5enMUx560QzuA==\"\n    }\n  },\n  \"dependencies\": {\n    \"async\": {\n      \"version\": \"2.1.1\",\n      \"resolved\": \"https://registry.npmjs.org/async/-/async-2.1.1.tgz\",\n      \"integrity\": \"sha1-4RttEAQ/IlTvthohFj2EDM3bjSg=\",\n      \"dev\": true,\n      \"requires\": {\n        \"lodash\": \"^4.14.0\"\n      }\n    },\n    \"lodash\": {\n      \"version\": \"4.17.21\",\n      \"resolved\": \"https://registry.npmjs.org/lodash/-/lodash-4.17.21.tgz\",\n      \"integrity\": \"sha512-v2kDEe57lecTulaDIuNTPy3Ry4gLGJ6Z1O3vE1krgXZNrsQ+LFTGHVxVjcXPs17LhbZVGedAJv8XZ1tvj5FvSg==\",\n      \"dev\": true\n    },\n    \"moment\": {\n      \"version\": \"2.21.0\",\n      \"resolved\": \"https://registry.npmjs.org/moment/-/moment-2.21.0.tgz\",\n      \"integrity\": \"sha512-TCZ36BjURTeFTM/CwRcViQlfkMvL1/vFISuNLO5GkcVm1+QHfbSiNqZuWeMFjj1/3+uAjXswgRk30j1kkLYJBQ==\",\n      \"peer\": true\n    },\n    \"underscore\": {\n      \"version\": \"1.13.3\",\n      \"resolved\": \"https://registry.npmjs.org/underscore/-/underscore-1.13.3.tgz\",\n      \"integrity\": \"sha512-QvjkYpiD+dJJraRA8+dGAU4i7aBbb2s0S3jA45TFOvg2VgqvdCDd/3N6CqA8gluk1W91GLoXg5enMUx560QzuA==\"\n    }\n  }\n}\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.PackageLockJsonFile, verb)
	assert.Equal(t, formats.JsonFormat, format)

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := ParsePackageLockJsonFile(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 5, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, "2.1.1", item.Locked)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, 0, len(item.Installations))
	assert.Equal(t, types.DevDependency, item.Type)

	item = result[1]
	assert.Equal(t, "lodash", item.Name)
	assert.Equal(t, "4.17.21", item.Locked)
	assert.Equal(t, types.DevDependency, item.Type)

	item = result[2]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, "2.21.0", item.Locked)
	assert.Equal(t, types.PeerDependency, item.Type)

	item = result[3]
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "", item.Locked)

	item = result[4]
	assert.Equal(t, "underscore", item.Name)
	assert.Equal(t, "1.13.3", item.Locked)
	assert.Equal(t, types.ProdDependency, item.Type)
}
