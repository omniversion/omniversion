package npm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseListOutput(t *testing.T) {
	vector := `foobar@3.12.0 /Users/testortestington/Documents/Repositories/foobar/frontend
├── covfefe@ extraneous
├── @test/test1@0.1301.2
├── test/test2@13.1.2
└── test3.js@0.11.4
`

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))

	item := result[0]
	assert.Equal(t, "@test/test1", item.Name)
	assert.Equal(t, "0.1301.2", item.Wanted)
	assert.Equal(t, "", item.Current)

	item = result[1]
	assert.Equal(t, "covfefe", item.Name)
	assert.Equal(t, "", item.Wanted)
	assert.Equal(t, "", item.Current)

	item = result[2]
	assert.Equal(t, "test/test2", item.Name)
	assert.Equal(t, "13.1.2", item.Wanted)
	assert.Equal(t, "", item.Current)

	item = result[3]
	assert.Equal(t, "test3.js", item.Name)
	assert.Equal(t, "0.11.4", item.Wanted)
	assert.Equal(t, "", item.Current)
}

func TestParseNpmOutdatedOutput(t *testing.T) {
	vector :=
		`Package                                        Current          Wanted          Latest  Location                                             Depended by
@angular-devkit/architect                     0.1301.2        0.1301.4        0.1303.5  node_modules/@angular-devkit/architect               frontend
@angular-devkit/build-angular                   13.1.2          13.3.5          13.3.5  node_modules/@angular-devkit/build-angular           frontend
@tiptap/core                            2.0.0-beta.162  2.0.0-beta.175  2.0.0-beta.176  node_modules/@tiptap/core                            frontend
selenium-webdriver                        4.0.0-beta.3           4.1.2           4.1.2  node_modules/selenium-webdriver                      frontend
`

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))

	item := result[0]
	assert.Equal(t, "@angular-devkit/architect", item.Name)
	assert.Equal(t, "0.1301.2", item.Current)
	assert.Equal(t, "0.1301.4", item.Wanted)
	assert.Equal(t, "0.1303.5", item.Latest)

	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "0.1301.2", item.Installations[0].Version)
	assert.Equal(t, "node_modules/@angular-devkit/architect", item.Installations[0].Location)

	item = result[1]
	assert.Equal(t, "@angular-devkit/build-angular", item.Name)
	assert.Equal(t, "13.1.2", item.Current)
	assert.Equal(t, "13.3.5", item.Wanted)
	assert.Equal(t, "13.3.5", item.Latest)

	item = result[2]
	assert.Equal(t, "@tiptap/core", item.Name)
	assert.Equal(t, "2.0.0-beta.162", item.Current)
	assert.Equal(t, "2.0.0-beta.175", item.Wanted)
	assert.Equal(t, "2.0.0-beta.176", item.Latest)

	item = result[3]
	assert.Equal(t, "selenium-webdriver", item.Name)
	assert.Equal(t, "4.0.0-beta.3", item.Current)
	assert.Equal(t, "4.1.2", item.Wanted)
	assert.Equal(t, "4.1.2", item.Latest)
}

func TestParseNpmAuditOutput(t *testing.T) {
	vector := "# npm audit report\n\nasync  >=3.0.0 <3.2.2 || <2.6.4\nSeverity: high\nPrototype Pollution in async - https://github.com/advisories/GHSA-fwr7-v2mv-hh25\nPrototype Pollution in async - https://github.com/advisories/GHSA-fwr7-v2mv-hh25\nfix available via `npm audit fix --force`\nWill install pdf2json@2.0.1, which is a breaking change\nnode_modules/async\nnode_modules/jake/node_modules/async\nnode_modules/mail-listener2/node_modules/async\nnode_modules/pdf2json/node_modules/async\nnode_modules/portfinder/node_modules/async\nnode_modules/protractor-jasmine2-html-reporter/node_modules/async\n  jake  8.0.1 - 10.8.4\n  Depends on vulnerable versions of async\n  node_modules/jake\n  mail-listener2  >=0.1.5\n  Depends on vulnerable versions of async\n  node_modules/mail-listener2\n  pdf2json  0.2.6 - 0.7.1 || 1.2.0 - 1.3.1\n  Depends on vulnerable versions of async\n  node_modules/pdf2json\n\nejs  <3.1.7\nSeverity: high\nTemplate injection in ejs - https://github.com/advisories/GHSA-phwq-j96m-2c2q\nfix available via `npm audit fix`\nnode_modules/ejs\n\nminimist  <1.2.6\nSeverity: critical\nPrototype Pollution in minimist - https://github.com/advisories/GHSA-xvch-5gv4-984h\nfix available via `npm audit fix`\nnode_modules/minimist\n\nmoment  <2.29.2\nSeverity: high\nPath Traversal: 'dir/../../filename' in moment.locale - https://github.com/advisories/GHSA-8hfj-j24r-96c4\nfix available via `npm audit fix`\nnode_modules/moment\n\nnode-forge  <1.3.0\nSeverity: high\nImproper Verification of Cryptographic Signature in node-forge - https://github.com/advisories/GHSA-cfm4-qjh2-4765\nfix available via `npm audit fix`\nnode_modules/node-forge\n  selfsigned  1.1.1 - 1.10.14\n  Depends on vulnerable versions of node-forge\n  node_modules/selfsigned\n    webpack-dev-server  2.5.0 - 4.7.2\n    Depends on vulnerable versions of selfsigned\n    node_modules/webpack-dev-server\n      @angular-devkit/build-angular  <=13.2.0-rc.1\n      Depends on vulnerable versions of webpack-dev-server\n      node_modules/@angular-devkit/build-angular\n\n11 vulnerabilities (10 high, 1 critical)\n\nTo address issues that do not require attention, run:\n  npm audit fix\n\nTo address all issues (including breaking changes), run:\n  npm audit fix --force"

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	// we only count the vulnerable packages themselves,
	// not the packages that depend on them...
	assert.Equal(t, 5, len(result))

	item := result[0]
	assert.Equal(t, "async", item.Name)
	assert.Equal(t, 1, len(item.Advisories))
	assert.Equal(t, "high", item.Advisories[0].Severity)
	assert.Equal(t, ">=3.0.0 <3.2.2 || <2.6.4", item.Advisories[0].VulnerableVersions)
}

func TestParseDefaultVersionOutput(t *testing.T) {
	vector := `{
  foobar: "3.12.0",
  npm: "6.14.16",
  ares: "1.18.1",
  brotli: "1.0.9",
  cldr: "40.0",
  icu: "70.1",
  llhttp: "2.1.4",
  modules: "83",
  napi: "8",
  nghttp2: "1.42.0",
  node: "14.19.1",
  openssl: "1.1.1n",
  tz: "2021a3",
  unicode: "14.0",
  uv: "1.42.0",
  v8: "8.4.371.23-node.85",
  zlib: "1.2.11"
}`

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 17, len(result))
}
