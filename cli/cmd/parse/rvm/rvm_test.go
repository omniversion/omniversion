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
	assert.Equal(t, "rvm", item.Pm)

	assert.Equal(t, "3.1.0", item.Default)
	assert.Equal(t, "3.1.0", item.Version)
	assert.Equal(t, "x86_64", item.Architecture)

	assert.Equal(t, 4, len(item.Installed))
	assert.Equal(t, "2.6.3", item.Installed[0].Version)
	assert.Equal(t, "2.6.6", item.Installed[1].Version)
	assert.Equal(t, "2.7.2", item.Installed[2].Version)
	assert.Equal(t, "3.1.0", item.Installed[3].Version)
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
	assert.Equal(t, "rvm", item.Pm)

	assert.Equal(t, "2.7.2", item.Default)
	assert.Equal(t, "2.6.6", item.Version)
	assert.Equal(t, "x86_64", item.Architecture)
}
