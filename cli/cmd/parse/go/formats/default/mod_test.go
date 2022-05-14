package _default

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseGoModFile(t *testing.T) {
	vector := "module github.com/omniversion/omniversion/cli\n\ngo 1.17\n\nrequire (\n\tgithub.com/BurntSushi/toml v1.1.0\n\tgithub.com/hashicorp/go-multierror v1.1.1\n\tgithub.com/mitchellh/mapstructure v1.5.0\n\tgithub.com/sirupsen/logrus v1.8.1\n\tgithub.com/spf13/cobra v1.4.0\n\tgithub.com/stretchr/testify v1.7.1\n\tgolang.org/x/mod v0.5.1\n\tgopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c\n)\n\nrequire (\n\tgithub.com/davecgh/go-spew v1.1.1 // indirect\n\tgithub.com/hashicorp/errwrap v1.0.0 // indirect\n\tgithub.com/inconshreveable/mousetrap v1.0.0 // indirect\n\tgithub.com/pmezard/go-difflib v1.0.0 // indirect\n\tgithub.com/spf13/pflag v1.0.5 // indirect\n\tgolang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect\n\tgolang.org/x/xerrors v0.0.0-20191011141410-1b5146add898 // indirect\n)\n"

	result, err := ParseGoModFile(vector)

	assert.Nil(t, err)
	assert.Equal(t, 17, len(result))

	item := result[0]
	assert.Equal(t, "github.com/omniversion/omniversion/cli", item.Name)
	assert.Equal(t, []string{"omniversion/cli"}, item.Aliases)
	assert.Equal(t, "", item.Current)
	assert.Zero(t, len(item.Installations))

	item = result[1]
	assert.Equal(t, "go", item.Name)
	assert.Equal(t, "1.17", item.Wanted)
	assert.Zero(t, len(item.Installations))

	item = result[2]
	assert.Equal(t, "github.com/BurntSushi/toml", item.Name)
	assert.Equal(t, []string{"toml"}, item.Aliases)
	assert.Equal(t, "1.1.0", item.Current)
	assert.Zero(t, len(item.Installations))

	item = result[8]
	assert.Equal(t, "golang.org/x/mod", item.Name)
	assert.Equal(t, []string{"mod"}, item.Aliases)

	item = result[9]
	assert.Equal(t, "gopkg.in/yaml.v3", item.Name)
	assert.Equal(t, []string{"yaml.v3"}, item.Aliases)
}
