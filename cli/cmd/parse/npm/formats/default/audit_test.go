package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAuditOutput_NoLockfile(t *testing.T) {
	vector := "npm ERR! code ENOLOCK\nnpm ERR! audit This command requires an existing lockfile.\nnpm ERR! audit Try creating one first with: npm i --package-lock-only\nnpm ERR! audit Original error: loadVirtual requires existing shrinkwrap file\n\nnpm ERR! A complete log of this run can be found in:\nnpm ERR!     /Users/testor/.npm/_logs/2022-05-09T07_09_09_064Z-debug.log\n"

	strippedInput, stderrOutput, _ := stderr.Strip(vector)
	verb, format := formats.DetectVerbAndFormat(strippedInput, stderrOutput)
	assert.Equal(t, formats.AuditCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	result, err := ParseAuditOutput(strippedInput, stderrOutput)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(result))
}

func TestParseAuditOutput(t *testing.T) {
	vector := "# npm audit report\n\nasync  <2.6.4\nSeverity: high\nPrototype Pollution in async - https://github.com/advisories/GHSA-fwr7-v2mv-hh25\nfix available via `npm audit fix --force`\nWill install async@2.6.4, which is outside the stated dependency range\nnode_modules/async\n\nmoment  <2.29.2\nSeverity: high\nPath Traversal: 'dir/../../filename' in moment.locale - https://github.com/advisories/GHSA-8hfj-j24r-96c4\nfix available via `npm audit fix --force`\nWill install moment@2.29.3, which is outside the stated dependency range\nnode_modules/moment\n\n2 high severity vulnerabilities\n\nTo address all issues, run:\n  npm audit fix --force\n"

	verb, format := formats.DetectVerbAndFormat(vector, stderr.Output{})
	assert.Equal(t, formats.AuditCommand, verb)
	assert.Equal(t, formats.DefaultFormat, format)

	result, err := ParseAuditOutput(vector, stderr.Output{})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	item := result[0]
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, 1, len(item.Advisories))
	assert.Equal(t, "high", item.Advisories[0].Severity)
	assert.Equal(t, "<2.6.4", item.Advisories[0].VulnerableVersions)

	item = result[1]
	assert.Equal(t, "moment", item.Name)
	assert.Equal(t, 1, len(item.Advisories))
	assert.Equal(t, "high", item.Advisories[0].Severity)
	assert.Equal(t, "<2.29.2", item.Advisories[0].VulnerableVersions)
}
