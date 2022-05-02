package parse

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLockfileVersionMismatchWarning(t *testing.T) {
	vector := `npm WARN read-shrinkwrap This version of npm is compatible with lockfileVersion@1, but package-lock.json was generated for lockfileVersion@2. I'll try to do my best with it!
/srv/foobar/releases/20220420102847/frontend`
	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(result))
}

func TestMissingWarning(t *testing.T) {
	vector := `npm ERR! missing: test@1.5.8, required by foo/bar@13.1.1
npm ERR! missing: test2@2.3.1, required by @some/other@13.1.1`
	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.Pm)
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "", item.Version)
	assert.Equal(t, "1.5.8", item.Wanted)

	item = result[1]
	assert.Equal(t, "npm", item.Pm)
	assert.Equal(t, "test2", item.Name)
	assert.Equal(t, "", item.Version)
	assert.Equal(t, "2.3.1", item.Wanted)
}

func TestExtraneousWarning(t *testing.T) {
	vector := `npm ERR! extraneous: foo/bar@2.0.0-beta.16 /srv/foobar/releases/20220420102847/frontend/node_modules/test/parent1
npm ERR! extraneous: foo/bar2@2.0.0-beta.13 /srv/foobar/releases/20220420102847/frontend/node_modules/test/parent2`
	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.Pm)
	assert.Equal(t, "foo/bar", item.Name)
	assert.Equal(t, "2.0.0-beta.16", item.Version)
	assert.Equal(t, 1, len(item.Installed))
	assert.Equal(t, "/srv/foobar/releases/20220420102847/frontend/node_modules/test/parent1", item.Installed[0].Location)

	item = result[1]
	assert.Equal(t, "npm", item.Pm)
	assert.Equal(t, "foo/bar2", item.Name)
	assert.Equal(t, "2.0.0-beta.13", item.Version)
	assert.Equal(t, 1, len(item.Installed))
	assert.Equal(t, "/srv/foobar/releases/20220420102847/frontend/node_modules/test/parent2", item.Installed[0].Location)
}

func TestParseNpmOutputWithAt(t *testing.T) {
	vector := `npm ERR! missing: @foo/bar@2.0.0-beta.13 /srv/foobar/releases/20220420102847/frontend/node_modules/test/parent`
	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.Pm)
	assert.Equal(t, "@foo/bar", item.Name)
	assert.Equal(t, "2.0.0-beta.13", item.Wanted)
	assert.Equal(t, 1, len(item.Installed))
	assert.Equal(t, "/srv/foobar/releases/20220420102847/frontend/node_modules/test/parent", item.Installed[0].Location)
}

func TestParseNpmKeyedDependencies(t *testing.T) {
	vector := `{
  "version": "1.0.0",
  "name": "foobar",
  "dependencies": {
    "test/dep": {
      "version": "0.1301.2",
      "resolved": "https://registry.npmjs.org/test/dep/-/dep-0.1301.2.tgz"
    }
  }
}`
	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.Pm)
	assert.Equal(t, "test/dep", item.Name)
	assert.Equal(t, "0.1301.2", item.Version)
}

func TestParseNpmFlatJson(t *testing.T) {
	vector := `{
  "foobar": {
    "current": "9.1.1",
    "wanted": "9.1.1",
    "latest": "10.7.0",
    "dependent": "frontend",
    "location": "/Users/me/Documents/Repositories/covfefe/frontend/node_modules/foobar"
  }
}`
	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.Pm)
	assert.Equal(t, "foobar", item.Name)
	assert.Equal(t, "9.1.1", item.Version)
	assert.Equal(t, "9.1.1", item.Wanted)
	assert.Equal(t, "10.7.0", item.Latest)
}

func TestParseNpmListWithMissing(t *testing.T) {
	vector := `/srv/foobar/releases/20220420102847/frontend:test1@10.0.0:MISSING:test1@10.0.1
/srv/foobar/releases/20220420102847/frontend:test2@10.0.0:MISSING:test2@10.0.0
/srv/foobar/releases/20220420102847/frontend:@test/test3@10.0.0:MISSING:@test/test3@10.0.0
/srv/foobar/releases/20220420102847/frontend:test4@13.3.4:MISSING:test4@13.3.4
/srv/foobar/releases/20220420102847/frontend:test5@13.3.4:MISSING:test5@13.3.4
/srv/foobar/releases/20220420102847/frontend:test6@13.3.3:MISSING:test6@13.3.3
`
	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 6, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.Pm)
	assert.Equal(t, "test1", item.Name)
	assert.Equal(t, "", item.Version)
	assert.Equal(t, "10.0.0", item.Wanted)
	assert.Equal(t, "10.0.1", item.Latest)
	assert.Equal(t, 0, len(item.Installed))

	assert.Equal(t, "test2", result[1].Name)
	assert.Equal(t, "@test/test3", result[2].Name)
	assert.Equal(t, "test4", result[3].Name)
	assert.Equal(t, "test5", result[4].Name)
	assert.Equal(t, "test6", result[5].Name)
}

func TestParseNpmList(t *testing.T) {
	vector := `/Users/testor/Documents/Repositories/foobar/frontend/node_modules/@angular-eslint/template-parser:@angular-eslint/template-parser@13.2.1:@angular-eslint/template-parser@13.0.1:@angular-eslint/template-parser@13.2.4:frontend`

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.Pm)
	assert.Equal(t, "@angular-eslint/template-parser", item.Name)
	assert.Equal(t, "13.0.1", item.Version)
	assert.Equal(t, "13.2.1", item.Wanted)
	assert.Equal(t, "13.2.4", item.Latest)
	assert.Equal(t, "13.0.1", item.Installed[0].Version)
	assert.Equal(t, "/Users/testor/Documents/Repositories/foobar/frontend/node_modules/@angular-eslint/template-parser", item.Installed[0].Location)
}

func TestParseNpmAuditJson(t *testing.T) {
	vector := `{
  "advisories": {
    "1067323": {
      "findings": [
        {
          "version": "1.2.0",
          "paths": [
            "@angular-devkit/build-angular>webpack-dev-server>selfsigned>node-forge"
          ]
        }
      ],
      "metadata": null,
      "vulnerable_versions": "<1.3.0",
      "module_name": "node-forge",
      "severity": "high",
      "github_advisory_id": "GHSA-x4jg-mjrx-434g",
      "cves": [
        "CVE-2022-24772"
      ],
      "access": "public",
      "patched_versions": ">=1.3.0",
      "cvss": {
        "score": 7.5,
        "vectorString": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:H/A:N"
      },
      "updated": "2022-03-30T20:07:57.000Z",
      "recommendation": "Upgrade to version 1.3.0 or later",
      "cwe": [
        "CWE-347"
      ],
      "found_by": null,
      "deleted": null,
      "id": 1067323,
      "references": "- https://github.com/digitalbazaar/forge/security/advisories/GHSA-x4jg-mjrx-434g\n- https://nvd.nist.gov/vuln/detail/CVE-2022-24772\n- https://github.com/digitalbazaar/forge/commit/3f0b49a0573ef1bb7af7f5673c0cfebf00424df1\n- https://github.com/digitalbazaar/forge/commit/bb822c02df0b61211836472e29b9790cc541cdb2\n- https://github.com/advisories/GHSA-x4jg-mjrx-434g",
      "created": "2022-03-18T23:10:28.000Z",
      "reported_by": null,
      "title": "Improper Verification of Cryptographic Signature in node-forge",
      "npm_advisory_id": null,
      "overview": "### Impact\n\nRSA PKCS#1 v1.5 signature verification code does not check for tailing garbage bytes after decoding a DigestInfo ASN.1 structure. This can allow padding bytes to be removed and garbage data added to forge a signature when a low public exponent is being used.\n\n",
      "url": "https://github.com/advisories/GHSA-x4jg-mjrx-434g"
    }
  }
}
`

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.Pm)
	assert.Equal(t, "node-forge", item.Name)
	assert.Equal(t, "1.2.0", item.Version)
	assert.Equal(t, 1, len(item.Advisories))

	advisory := item.Advisories[0]
	assert.Equal(t, "public", advisory.Access)
	assert.Equal(t, 7.5, advisory.CVSSScore)
	assert.Equal(t, 1067323, advisory.Id)
	assert.True(t, strings.HasPrefix(advisory.Overview, "### Impact"))
	assert.Equal(t, ">=1.3.0", advisory.PatchedVersions)
	assert.Equal(t, "Upgrade to version 1.3.0 or later", advisory.Recommendation)
	assert.True(t, strings.HasPrefix(advisory.References, "- https://github.com/digitalbazaar/forge/security/advisories/GHSA-x4jg-mjrx-434g"))
	assert.Equal(t, "high", advisory.Severity)
	assert.Equal(t, "Improper Verification of Cryptographic Signature in node-forge", advisory.Title)
	assert.Equal(t, "https://github.com/advisories/GHSA-x4jg-mjrx-434g", advisory.Url)
	assert.Equal(t, "<1.3.0", advisory.VulnerableVersions)
}

func TestParseNpmAuditJson2(t *testing.T) {
	vector := `{
  "advisories": {
    "1067342": {
      "findings": [
        {
          "version": "1.2.5",
          "paths": [
            "webdriver-manager>minimist",
            "ng2-validation>libphonenumber-js>minimist",
            "@angular/compiler-cli>@babel/core>json5>minimist",
            "karma-typescript>istanbul-lib-instrument>@babel/core>json5>minimist",
            "@angular-devkit/build-angular>babel-plugin-istanbul>istanbul-lib-instrument>@babel/core>json5>minimist",
            "pdfmake>svg-to-pdfkit>pdfkit>linebreak>brfs>quote-stream>minimist"
          ]
        }
      ],
      "metadata": null,
      "vulnerable_versions": "<1.2.6",
      "module_name": "minimist",
      "severity": "critical",
      "github_advisory_id": "GHSA-xvch-5gv4-984h",
      "cves": [
        "CVE-2021-44906"
      ],
      "access": "public",
      "patched_versions": ">=1.2.6",
      "cvss": {
        "score": 9.8,
        "vectorString": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H"
      },
      "updated": "2022-04-04T21:39:39.000Z",
      "recommendation": "Upgrade to version 1.2.6 or later",
      "cwe": [
        "CWE-1321"
      ],
      "found_by": null,
      "deleted": null,
      "id": 1067342,
      "references": "- https://nvd.nist.gov/vuln/detail/CVE-2021-44906\n- https://github.com/substack/minimist/issues/164\n- https://github.com/substack/minimist/blob/master/index.js#L69\n- https://snyk.io/vuln/SNYK-JS-MINIMIST-559764\n- https://stackoverflow.com/questions/8588563/adding-custom-properties-to-a-function/20278068#20278068\n- https://github.com/Marynk/JavaScript-vulnerability-detection/blob/main/minimist%20PoC.zip\n- https://github.com/advisories/GHSA-xvch-5gv4-984h",
      "created": "2022-03-18T00:01:09.000Z",
      "reported_by": null,
      "title": "Prototype Pollution in minimist",
      "npm_advisory_id": null,
      "overview": "Minimist <=1.2.5 is vulnerable to Prototype Pollution via file index.js, function setKey() (lines 69-95).",
      "url": "https://github.com/advisories/GHSA-xvch-5gv4-984h"
    }
  }
}
`

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.Pm)
	assert.Equal(t, "minimist", item.Name)
	assert.Equal(t, "1.2.5", item.Version)
	assert.Equal(t, 1, len(item.Advisories))

	advisory := item.Advisories[0]
	assert.Equal(t, "public", advisory.Access)
	assert.Equal(t, 9.8, advisory.CVSSScore)
	assert.Equal(t, 1067342, advisory.Id)
	assert.True(t, strings.HasPrefix(advisory.Overview, "Minimist <=1.2.5"))
	assert.Equal(t, ">=1.2.6", advisory.PatchedVersions)
	assert.Equal(t, "Upgrade to version 1.2.6 or later", advisory.Recommendation)
	assert.True(t, strings.HasPrefix(advisory.References, "- https://nvd.nist.gov/vuln/detail/CVE-2021-44906"))
	assert.Equal(t, "critical", advisory.Severity)
	assert.Equal(t, "Prototype Pollution in minimist", advisory.Title)
	assert.Equal(t, "https://github.com/advisories/GHSA-xvch-5gv4-984h", advisory.Url)
	assert.Equal(t, "<1.2.6", advisory.VulnerableVersions)
}

func TestParseNpmVersion(t *testing.T) {
	vector := `{
  "foobar": "3.12.0",
  "npm": "6.14.16",
  "ares": "1.18.1",
  "brotli": "1.0.9",
  "cldr": "40.0",
  "icu": "70.1",
  "llhttp": "2.1.4",
  "modules": "83",
  "napi": "8",
  "nghttp2": "1.42.0",
  "node": "14.19.1",
  "openssl": "1.1.1n",
  "tz": "2021a3",
  "unicode": "14.0",
  "uv": "1.42.0",
  "v8": "8.4.371.23-node.85",
  "zlib": "1.2.11"
}`

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 17, len(result))
}
