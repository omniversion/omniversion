package json

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseVersionOutput(t *testing.T) {
	vector := "{\n  \"test\": \"1.0.0\",\n  \"npm\": \"8.1.2\",\n  \"node\": \"16.13.2\",\n  \"v8\": \"9.4.146.24-node.14\",\n  \"uv\": \"1.42.0\",\n  \"zlib\": \"1.2.11\",\n  \"brotli\": \"1.0.9\",\n  \"ares\": \"1.18.1\",\n  \"modules\": \"93\",\n  \"nghttp2\": \"1.45.1\",\n  \"napi\": \"8\",\n  \"llhttp\": \"6.0.4\",\n  \"openssl\": \"1.1.1l+quic\",\n  \"cldr\": \"39.0\",\n  \"icu\": \"69.1\",\n  \"tz\": \"2021a\",\n  \"unicode\": \"13.0\",\n  \"ngtcp2\": \"0.1.0-DEV\",\n  \"nghttp3\": \"0.1.0-DEV\"\n}\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.VersionCommand, verb)
	assert.Equal(t, formats.JsonFormat, format)

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := ParseVersionOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 19, len(result))
}
