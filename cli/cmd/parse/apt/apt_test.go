package apt

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAptSimpleOutput(t *testing.T) {
	vector := `Listing... Done
autoconf/bionic,now 2.69-11 all [installed]
autotools-dev/bionic-updates,now 20180224.1 all [installed,automatic]
mde-netfilter/insiders-fast,bionic,now 100.69.32 amd64 [installed,upgradable to: 100.69.45]`

	result, err := parseAptOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 3, len(result))

	item := result[0]
	assert.Equal(t, "autoconf", item.Name)
	assert.Equal(t, "2.69-11", item.Current)
	assert.Equal(t, "all", item.Architecture)
	assert.Equal(t, 2, len(item.Sources))
	assert.Equal(t, "bionic", item.Sources[0])
	assert.Equal(t, "now", item.Sources[1])

	item = result[1]
	assert.Equal(t, "autotools-dev", item.Name)
	assert.Equal(t, "20180224.1", item.Current)
	assert.Equal(t, "all", item.Architecture)
	assert.Equal(t, 2, len(item.Sources))
	assert.Equal(t, "bionic-updates", item.Sources[0])
	assert.Equal(t, "now", item.Sources[1])

	item = result[2]
	assert.Equal(t, "mde-netfilter", item.Name)
	assert.Equal(t, "100.69.32", item.Current)
	assert.Equal(t, "amd64", item.Architecture)
	assert.Equal(t, 3, len(item.Sources))
	assert.Equal(t, "insiders-fast", item.Sources[0])
	assert.Equal(t, "bionic", item.Sources[1])
	assert.Equal(t, "now", item.Sources[2])
	assert.Equal(t, "100.69.45", item.Latest)
}

func TestParseAptNotInstalledOutput(t *testing.T) {
	vector := `Listing... Done
zssh/bionic 1.5c.debian.1-4 amd64`

	result, err := parseAptOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Empty(t, item.PackageManager)
	assert.Equal(t, "zssh", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "1.5c.debian.1-4", item.Wanted)
	assert.Equal(t, "amd64", item.Architecture)
	assert.Equal(t, 1, len(item.Sources))
	assert.Equal(t, "bionic", item.Sources[0])
}

func TestParseAptOutdatedOutput(t *testing.T) {
	vector := `Listing... Done
mde-netfilter/insiders-fast 100.69.45 amd64 [upgradable from: 100.69.32]
N: There are 4 additional versions. Please use the '-a' switch to see them.`

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := parseAptOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "apt", item.PackageManager)
	assert.Equal(t, "mde-netfilter", item.Name)
	assert.Equal(t, "100.69.32", item.Current)
	assert.Equal(t, "100.69.45", item.Latest)
	assert.Equal(t, "100.69.32", item.Installations[0].Version)
	assert.Equal(t, "amd64", item.Architecture)
	assert.Equal(t, 1, len(item.Sources))
	assert.Equal(t, "insiders-fast", item.Sources[0])
}
