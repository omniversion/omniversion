package _go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseListOutput(t *testing.T) {
	vector := "github.com/omniversion/omniversion/cli\ngithub.com/BurntSushi/toml v1.1.0\ngithub.com/cpuguy83/go-md2man/v2 v2.0.1\ngithub.com/davecgh/go-spew v1.1.1\ngithub.com/hashicorp/errwrap v1.0.0\ngithub.com/hashicorp/go-multierror v1.1.1\ngithub.com/inconshreveable/mousetrap v1.0.0\ngithub.com/mitchellh/mapstructure v1.5.0\ngithub.com/pmezard/go-difflib v1.0.0\ngithub.com/russross/blackfriday/v2 v2.1.0\ngithub.com/sirupsen/logrus v1.8.1\ngithub.com/spf13/cobra v1.4.0\ngithub.com/spf13/pflag v1.0.5\ngithub.com/stretchr/objx v0.1.0\ngithub.com/stretchr/testify v1.7.1\ngolang.org/x/crypto v0.0.0-20191011191535-87dc89f01550\ngolang.org/x/mod v0.5.1\ngolang.org/x/sys v0.0.0-20191026070338-33540a1f6037\ngolang.org/x/tools v0.0.0-20191119224855-298f0cb1881e\ngolang.org/x/xerrors v0.0.0-20191011141410-1b5146add898\ngopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405\ngopkg.in/yaml.v2 v2.4.0\ngopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c\n"

	result, err := ParseListOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 23, len(result))

	item := result[0]
	assert.Equal(t, "omniversion/cli", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "github.com/omniversion/omniversion/cli", item.InstallPath)

	item = result[1]
	assert.Equal(t, "toml", item.Name)
	assert.Equal(t, "1.1.0", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "github.com/BurntSushi/toml", item.InstallPath)

	item = result[2]
	assert.Equal(t, "go-md2man/v2", item.Name)

	item = result[15]
	assert.Equal(t, "crypto", item.Name)

	item = result[22]
	assert.Equal(t, "yaml.v3", item.Name)
}

func TestParseListOutput_VersionsFlag(t *testing.T) {
	vector := "github.com/omniversion/omniversion/cli\ngithub.com/BurntSushi/toml v0.1.0 v0.2.0 v0.3.0 v0.3.1 v0.4.0 v0.4.1 v1.0.0 v1.1.0\ngithub.com/cpuguy83/go-md2man/v2 v2.0.0 v2.0.1 v2.0.2\ngithub.com/davecgh/go-spew v1.0.0 v1.1.0 v1.1.1\ngithub.com/hashicorp/errwrap v1.0.0 v1.1.0\ngithub.com/hashicorp/go-multierror v1.0.0 v1.1.0 v1.1.1\ngithub.com/inconshreveable/mousetrap v1.0.0\ngithub.com/mitchellh/mapstructure v1.0.0 v1.1.0 v1.1.1 v1.1.2 v1.2.0 v1.2.1 v1.2.2 v1.2.3 v1.3.0 v1.3.1 v1.3.2 v1.3.3 v1.4.0 v1.4.1 v1.4.2 v1.4.3 v1.5.0\ngithub.com/pmezard/go-difflib v1.0.0\ngithub.com/russross/blackfriday/v2 v2.0.0 v2.0.1 v2.1.0-pre.1 v2.1.0\ngithub.com/sirupsen/logrus v0.1.0 v0.1.1 v0.2.0 v0.3.0 v0.4.0 v0.4.1 v0.5.0 v0.5.1 v0.6.0 v0.6.1 v0.6.2 v0.6.3 v0.6.4 v0.6.5 v0.6.6 v0.7.0 v0.7.1 v0.7.2 v0.7.3 v0.8.0 v0.8.1 v0.8.2 v0.8.3 v0.8.4 v0.8.5 v0.8.6 v0.8.7 v0.9.0 v0.10.0 v0.11.0 v0.11.1 v0.11.2 v0.11.3 v0.11.4 v0.11.5 v1.0.0 v1.0.1 v1.0.3 v1.0.4 v1.0.5 v1.0.6 v1.1.0 v1.1.1 v1.2.0 v1.3.0 v1.4.0 v1.4.1 v1.4.2 v1.5.0 v1.6.0 v1.7.0 v1.7.1 v1.8.0 v1.8.1\ngithub.com/spf13/cobra v0.0.1 v0.0.2 v0.0.3 v0.0.4 v0.0.5 v0.0.6 v0.0.7 v1.0.0 v1.1.0 v1.1.1 v1.1.2 v1.1.3 v1.2.0 v1.2.1 v1.3.0 v1.4.0\ngithub.com/spf13/pflag v1.0.0 v1.0.1 v1.0.2 v1.0.3 v1.0.4 v1.0.5-rc1 v1.0.5\ngithub.com/stretchr/objx v0.1.0 v0.1.1 v0.2.0 v0.3.0 v0.4.0\ngithub.com/stretchr/testify v1.1.1 v1.1.2 v1.1.3 v1.1.4 v1.2.0 v1.2.1 v1.2.2 v1.3.0 v1.4.0 v1.5.0 v1.5.1 v1.6.0 v1.6.1 v1.7.0 v1.7.1\ngolang.org/x/crypto\ngolang.org/x/mod v0.1.0 v0.2.0 v0.3.0 v0.4.0 v0.4.1 v0.4.2 v0.5.0 v0.5.1 v0.6.0-dev\ngolang.org/x/sys\ngolang.org/x/tools v0.1.0 v0.1.1 v0.1.2 v0.1.3 v0.1.4 v0.1.5 v0.1.6 v0.1.7 v0.1.8 v0.1.9 v0.1.10\ngolang.org/x/xerrors\ngopkg.in/check.v1\ngopkg.in/yaml.v2 v2.0.0 v2.1.0 v2.1.1 v2.2.0 v2.2.1 v2.2.2 v2.2.3 v2.2.4 v2.2.5 v2.2.6 v2.2.7 v2.2.8 v2.3.0 v2.4.0\ngopkg.in/yaml.v3\n"
	result, err := ParseListOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 23, len(result))

	item := result[0]
	assert.Equal(t, "omniversion/cli", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "github.com/omniversion/omniversion/cli", item.InstallPath)

	item = result[1]
	assert.Equal(t, "toml", item.Name)
	assert.Equal(t, "1.1.0", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "github.com/BurntSushi/toml", item.InstallPath)
	assert.Equal(t, 1, len(item.Sources))
	assert.Equal(t, []string{"0.1.0", "0.2.0", "0.3.0", "0.3.1", "0.4.0", "0.4.1", "1.0.0", "1.1.0"}, item.Sources[0].Versions)

	item = result[2]
	assert.Equal(t, "go-md2man/v2", item.Name)
	assert.Equal(t, 1, len(item.Sources))
	assert.Equal(t, []string{"2.0.0", "2.0.1", "2.0.2"}, item.Sources[0].Versions)

	item = result[15]
	assert.Equal(t, "crypto", item.Name)

	item = result[22]
	assert.Equal(t, "yaml.v3", item.Name)
}
