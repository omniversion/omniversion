package rubygems

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRubygemsSimpleOutput(t *testing.T) {
	vector := `prime (0.1.2)
    Author: Marc-Andre Lafortune
    Homepage: https://github.com/ruby/prime
    Licenses: Ruby, BSD-2-Clause
    Installed at: /usr/local/rvm/rubies/ruby-3.1.0/lib/ruby/gems/3.1.0

    Prime numbers and factorization library.`

	previousInjectValue := shared.InjectPackageManager
	shared.InjectPackageManager = true
	result, err := parseRubygemsOutput(vector)
	shared.InjectPackageManager = previousInjectValue

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	assert.Equal(t, `Marc-Andre Lafortune`, result[0].Author)
	assert.Equal(t, "rubygems", result[0].PackageManager)
	assert.Equal(t, "prime", result[0].Name)
	assert.Equal(t, 1, len(result[0].Installations))
	assert.Equal(t, "0.1.2", result[0].Current)
	assert.Equal(t, "/usr/local/rvm/rubies/ruby-3.1.0/lib/ruby/gems/3.1.0", result[0].Installations[0].Location)
}

func TestParseRubygemsMultilineOutput(t *testing.T) {
	vector := `bundler (2.3.7, 2.3.3, 2.1.4)
    Authors: André Arko, Samuel Giddins, Colby Swandale, Hiroshi
    Shibata, David Rodríguez, Grey Baker, Stephanie Morillo, Chris
    Morris, James Wen, Tim Moore, André Medeiros, Jessica Lynn Suttles,
    Terence Lee, Carl Lerche, Yehuda Katz
    Homepage: https://bundler.io
    License: MIT
    Installed at (2.3.7): /usr/local/rvm/gems/ruby-3.1.0
                 (2.3.3, default): /usr/local/rvm/rubies/ruby-3.1.0/lib/ruby/gems/3.1.0
                 (2.1.4): /usr/local/rvm/gems/ruby-3.1.0

    The best way to manage your application's dependencies.
    Bar none.

`

	result, err := parseRubygemsOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	assert.Equal(t, `André Arko, Samuel Giddins, Colby Swandale, Hiroshi
    Shibata, David Rodríguez, Grey Baker, Stephanie Morillo, Chris
    Morris, James Wen, Tim Moore, André Medeiros, Jessica Lynn Suttles,
    Terence Lee, Carl Lerche, Yehuda Katz`, result[0].Author)

	assert.Equal(t, "bundler", result[0].Name)
	assert.Equal(t, "", result[0].Current)
	assert.Equal(t, 3, len(result[0].Installations))
	assert.Equal(t, "2.3.3", result[0].Default)
	assert.Equal(t, "2.3.7", result[0].Installations[0].Version)
	assert.Equal(t, "2.3.3", result[0].Installations[1].Version)
	assert.Equal(t, "2.1.4", result[0].Installations[2].Version)
}

func TestParseRubygemsOutdatedOutput(t *testing.T) {
	vector := `bigdecimal (3.1.1 < 3.1.2)
bundler (2.3.7 < 2.3.12)`

	previousInjectValue := shared.InjectPackageManager
	shared.InjectPackageManager = true
	result, err := parseRubygemsOutput(vector)
	shared.InjectPackageManager = previousInjectValue

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	item := result[0]
	assert.Equal(t, "bigdecimal", item.Name)
	assert.Equal(t, "rubygems", item.PackageManager)
	assert.Equal(t, "3.1.1", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "3.1.1", item.Installations[0].Version)
	assert.Equal(t, "3.1.2", item.Latest)

	item = result[1]
	assert.Equal(t, "bundler", item.Name)
	assert.Equal(t, "2.3.7", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "2.3.7", item.Installations[0].Version)
	assert.Equal(t, "2.3.12", item.Latest)
}

func TestParseRubygemsInvalidVersions(t *testing.T) {
	vector := `test (3.1.1 < 3.1.2 < 4.5.6)`

	result, err := parseRubygemsOutput(vector)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unable to parse package description: \"test\"")
	assert.Zero(t, len(result))
}

func TestIncompletePackageDescription(t *testing.T) {
	vector := `test (1.2.3)

    Test description.
`

	result, err := parseRubygemsOutput(vector)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unable to parse package description: \"test\"")
	assert.Zero(t, len(result))
}

func TestSingleDefaultVersion(t *testing.T) {
	vector := `test (2.3.2)
    Authors: test
    Homepage: example.com
    License: MIT
    Installed at (default): /usr/local/rvm/gems/ruby-3.1.0

    Test description.
`

	result, err := parseRubygemsOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, 1, len(result[0].Installations))

	assert.Equal(t, "2.3.2", result[0].Current)
	assert.Equal(t, "2.3.2", result[0].Default)
	assert.Equal(t, "2.3.2", result[0].Installations[0].Version)
}

func TestParseGemListOutput(t *testing.T) {
	vector := `
*** LOCAL GEMS ***

abbrev (default: 0.1.0)
actioncable (7.0.2.3, 7.0.2, 6.1.0)
actionmailbox (7.0.2.3, 7.0.2)
addressable (2.8.0)
`

	result, err := parseRubygemsOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))

	item := result[0]
	assert.Equal(t, "abbrev", item.Name)
	assert.Equal(t, "0.1.0", item.Current)
	assert.Equal(t, "0.1.0", item.Default)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "0.1.0", item.Installations[0].Version)

	item = result[1]
	assert.Equal(t, "actioncable", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Default)
	assert.Equal(t, 3, len(item.Installations))
	assert.Equal(t, "7.0.2.3", item.Installations[0].Version)
	assert.Equal(t, "7.0.2", item.Installations[1].Version)
	assert.Equal(t, "6.1.0", item.Installations[2].Version)

	item = result[2]
	assert.Equal(t, "actionmailbox", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "", item.Default)
	assert.Equal(t, 2, len(item.Installations))
	assert.Equal(t, "7.0.2.3", item.Installations[0].Version)
	assert.Equal(t, "7.0.2", item.Installations[1].Version)

	item = result[3]
	assert.Equal(t, "addressable", item.Name)
	assert.Equal(t, "2.8.0", item.Current)
	assert.Equal(t, "", item.Default)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "2.8.0", item.Installations[0].Version)
}
