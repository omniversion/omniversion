package rubygems

import (
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

	result, err := parseRubygemsOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	assert.Equal(t, `Marc-Andre Lafortune`, result[0].Author)

	assert.Equal(t, "prime", result[0].Name)
	assert.Equal(t, 1, len(result[0].Installed))
	assert.Equal(t, "0.1.2", result[0].Version)
	assert.Equal(t, "/usr/local/rvm/rubies/ruby-3.1.0/lib/ruby/gems/3.1.0", result[0].Installed[0].Location)
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
	assert.Equal(t, "", result[0].Version)
	assert.Equal(t, 3, len(result[0].Installed))
	assert.Equal(t, "2.3.3", result[0].Default)
	assert.Equal(t, "2.3.7", result[0].Installed[0].Version)
	assert.Equal(t, "2.3.3", result[0].Installed[1].Version)
	assert.Equal(t, "2.1.4", result[0].Installed[2].Version)
}

func TestParseRubygemsOutdatedOutput(t *testing.T) {
	vector := `bigdecimal (3.1.1 < 3.1.2)
bundler (2.3.7 < 2.3.12)`

	result, err := parseRubygemsOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	item := result[0]
	assert.Equal(t, "bigdecimal", item.Name)
	assert.Equal(t, "3.1.1", item.Version)
	assert.Equal(t, 1, len(item.Installed))
	assert.Equal(t, "3.1.1", item.Installed[0].Version)
	assert.Equal(t, "3.1.2", item.Latest)

	item = result[1]
	assert.Equal(t, "bundler", item.Name)
	assert.Equal(t, "2.3.7", item.Version)
	assert.Equal(t, 1, len(item.Installed))
	assert.Equal(t, "2.3.7", item.Installed[0].Version)
	assert.Equal(t, "2.3.12", item.Latest)
}
