package rvm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRvmSimpleOutput(t *testing.T) {
	vector := `   ruby-2.6.3 [ x86_64 ]
   ruby-2.6.6 [ x86_64 ]
   ruby-2.7.2 [ x86_64 ]
=* ruby-3.1.0 [ x86_64 ]

# => - current
# =* - current && default
#  * - default

`

	result, err := parseRvmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "ruby", item.Name)
	assert.Equal(t, "rvm", item.PackageManager)

	assert.Equal(t, "3.1.0", item.Default)
	assert.Equal(t, "3.1.0", item.Current)
	assert.Equal(t, "x86_64", item.Architecture)

	assert.Equal(t, 4, len(item.Installations))
	assert.Equal(t, "2.6.3", item.Installations[0].Version)
	assert.Equal(t, "2.6.6", item.Installations[1].Version)
	assert.Equal(t, "2.7.2", item.Installations[2].Version)
	assert.Equal(t, "3.1.0", item.Installations[3].Version)
}

func TestParseRvmOutputWithNonDefaultCurrent(t *testing.T) {
	vector := `   ruby-2.6.3 [ no ]
=> ruby-2.6.6 [ x86_64 ]
 * ruby-2.7.2 [ no ]
   ruby-3.1.0 [ no ]

# => - current
# =* - current && default
#  * - default

`

	result, err := parseRvmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "ruby", item.Name)
	assert.Equal(t, "rvm", item.PackageManager)

	assert.Equal(t, "2.7.2", item.Default)
	assert.Equal(t, "2.6.6", item.Current)
	assert.Equal(t, "x86_64", item.Architecture)
}

func TestParseRvmVersionOutput(t *testing.T) {
	vector := `rvm 1.29.12 (latest) by Michal Papis, Piotr Kuczynski, Wayne E. Seguin [https://rvm.io]
`

	result, err := parseRvmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "rvm", item.Name)
	assert.Equal(t, "1.29.12", item.Current)
	assert.Equal(t, "1.29.12", item.Latest)
	assert.Equal(t, "Michal Papis, Piotr Kuczynski, Wayne E. Seguin", item.Author)
	assert.Equal(t, "https://rvm.io", item.Homepage)
}
