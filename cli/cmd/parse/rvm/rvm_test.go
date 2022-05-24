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

func TestParseRvmVersionOutput_Manual(t *testing.T) {
	vector := "rvm 1.29.12 (manual) by Michal Papis, Piotr Kuczynski, Wayne E. Seguin [https://rvm.io]\n"

	result, err := parseRvmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "rvm", item.Name)
	assert.Equal(t, "1.29.12", item.Current)
	assert.Equal(t, "", item.Latest)
	assert.Equal(t, "Michal Papis, Piotr Kuczynski, Wayne E. Seguin", item.Author)
	assert.Equal(t, "https://rvm.io", item.Homepage)
}

func TestParseRvmVersionOutput_PathWarning(t *testing.T) {
	vector := "Warning! PATH is not properly set up, /Users/testor/.rvm/gems/ruby-3.1.0/bin is not at first place.\n         Usually this is caused by shell initialization files. Search for PATH=... entries.\n         You can also re-add RVM to your profile by running: rvm get stable --auto-dotfiles\n         To fix it temporarily in this shell session run: rvm use ruby-3.1.0\n         To ignore this error add rvm_silence_path_mismatch_check_flag=1 to your ~/.rvmrc file.\nrvm 1.29.12 (latest) by Michal Papis, Piotr Kuczynski, Wayne E. Seguin [https://rvm.io]\n"

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
