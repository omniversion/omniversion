package nvm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseListOutput(t *testing.T) {
	vector := "        v14.8.0\n       v14.15.0\n       v14.17.0\n       v14.19.0\n        v16.2.0\n       v16.13.0\n       v16.13.1\n->     v16.13.2\n        v17.3.1\ndefault -> 16 (-> v16.13.2)\nset -> default (-> v16.13.2)\nnode -> stable (-> v17.3.1) (default)\nstable -> 17.3 (-> v17.3.1) (default)\niojs -> N/A (default)\nunstable -> N/A (default)\nlts/* -> lts/gallium (-> N/A)\nlts/argon -> v4.9.1 (-> N/A)\nlts/boron -> v6.17.1 (-> N/A)\nlts/carbon -> v8.17.0 (-> N/A)\nlts/dubnium -> v10.24.1 (-> N/A)\nlts/erbium -> v12.22.10 (-> N/A)\nlts/fermium -> v14.19.0\nlts/gallium -> v16.14.0 (-> N/A)\n"

	assert.False(t, isVersionOutput(vector))

	result, err := parseNvmOutput(vector)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "node", result[0].Name)
	assert.Equal(t, "nvm", result[0].PackageManager)
	assert.Equal(t, 9, len(result[0].Installations))
	assert.Equal(t, "16.13.2", result[0].Current)

	assert.Equal(t, "16.13.2", result[0].Installations[7].Version)
	// the "set" entry is probably a botched attempt to set an alias... :D
	assert.Equal(t, []string{"default", "16", "set"}, result[0].Installations[7].VersionAliases)
	assert.Equal(t, "17.3.1", result[0].Installations[8].Version)
	assert.Equal(t, []string{"node", "stable", "17.3"}, result[0].Installations[8].VersionAliases)
}
