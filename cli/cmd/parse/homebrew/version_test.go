package homebrew

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseHomebrewVersionOutput(t *testing.T) {
	vector := "Homebrew 3.4.10-49-g14ff6be\nHomebrew/homebrew-core (git revision 31e2ed54e34; last commit 2022-05-06)\nHomebrew/homebrew-cask (git revision b817992654; last commit 2022-05-06)\n\n"

	assert.True(t, isVersionCommandOutput(vector))

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := parseHomebrewOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "brew", item.PackageManager)
	assert.Equal(t, "homebrew", item.Name)
	assert.Equal(t, "3.4.10-49-g14ff6be", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "3.4.10-49-g14ff6be", item.Installations[0].Version)
}

func TestParseHomebrewVersionOutput_Invalid(t *testing.T) {
	vector := "Homebrew    "

	assert.True(t, isVersionCommandOutput(vector))

	result, err := parseHomebrewOutput(vector)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unexpected input format")
	assert.Zero(t, len(result))
}
