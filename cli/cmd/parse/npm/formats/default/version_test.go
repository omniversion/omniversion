package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseVersionOutput(t *testing.T) {
	vector := "{\n  test: '1.0.0',\n  npm: '8.1.2',\n  node: '16.13.2',\n  v8: '9.4.146.24-node.14',\n  uv: '1.42.0',\n  zlib: '1.2.11',\n  brotli: '1.0.9',\n  ares: '1.18.1',\n  modules: '93',\n  nghttp2: '1.45.1',\n  napi: '8',\n  llhttp: '6.0.4',\n  openssl: '1.1.1l+quic',\n  cldr: '39.0',\n  icu: '69.1',\n  tz: '2021a',\n  unicode: '13.0',\n  ngtcp2: '0.1.0-DEV',\n  nghttp3: '0.1.0-DEV'\n}\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.VersionCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	result, err := ParseVersionOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 19, len(result))
}

func TestParseVersionOutput_InvalidJson(t *testing.T) {
	vector := "{\n\t\"version\":"

	result, err := ParseVersionOutput(vector, stderr.Output{})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid version data")
	assert.Zero(t, len(result))
}
