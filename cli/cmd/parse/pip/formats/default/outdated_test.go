package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseOutdatedOutput(t *testing.T) {
	vector := "Package    Version Latest Type\n---------- ------- ------ -----\nretry      0.8.1   0.9.1  wheel\nsetuptools 20.6.7  21.0.0 wheel\n"

	verb, format := formats.DetectVerbAndFormat(vector)
	assert.Equal(t, formats.OutdatedCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := ParseOutdatedOutput(vector)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	retry := result[0]
	assert.Equal(t, "retry", retry.Name)
	assert.Equal(t, "pip", retry.PackageManager)
	assert.Equal(t, "0.8.1", retry.Current)
	assert.Equal(t, "0.9.1", retry.Latest)
	assert.Equal(t, 1, len(retry.Installations))
	assert.Equal(t, "0.8.1", retry.Installations[0].Version)

	setuptools := result[1]
	assert.Equal(t, "setuptools", setuptools.Name)
	assert.Equal(t, "pip", setuptools.PackageManager)
	assert.Equal(t, "20.6.7", setuptools.Current)
	assert.Equal(t, 1, len(setuptools.Installations))
	assert.Equal(t, "20.6.7", setuptools.Installations[0].Version)
}
