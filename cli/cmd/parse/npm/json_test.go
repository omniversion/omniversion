package npm

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

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
	previousInjectValue := shared.InjectPackageManager
	shared.InjectPackageManager = true
	result, err := parseNpmOutput(vector)
	shared.InjectPackageManager = previousInjectValue

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "test/dep", item.Name)
	assert.Equal(t, "0.1301.2", item.Wanted)
	assert.Equal(t, "", item.Current)
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
	previousInjectValue := shared.InjectPackageManager
	shared.InjectPackageManager = true
	result, err := parseNpmOutput(vector)
	shared.InjectPackageManager = previousInjectValue

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "foobar", item.Name)
	assert.Equal(t, "9.1.1", item.Current)
	assert.Equal(t, "9.1.1", item.Wanted)
	assert.Equal(t, "10.7.0", item.Latest)
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
	previousInjectValue := shared.InjectPackageManager
	shared.InjectPackageManager = true
	result, err := parseNpmOutput(vector)
	shared.InjectPackageManager = previousInjectValue

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "node-forge", item.Name)
	assert.Equal(t, "1.2.0", item.Current)
	assert.Equal(t, 1, len(item.Advisories))

	advisory := item.Advisories[0]
	assert.Equal(t, 7.5, advisory.CVSSScore)
	assert.Equal(t, "1067323", advisory.Identifier)
	assert.True(t, strings.HasPrefix(advisory.Overview, "### Impact"))
	assert.Equal(t, ">=1.3.0", advisory.PatchedVersions)
	assert.Equal(t, "Upgrade to version 1.3.0 or later", advisory.Recommendation)
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
	previousInjectValue := shared.InjectPackageManager
	shared.InjectPackageManager = true
	result, err := parseNpmOutput(vector)
	shared.InjectPackageManager = previousInjectValue

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "npm", item.PackageManager)
	assert.Equal(t, "minimist", item.Name)
	assert.Equal(t, "1.2.5", item.Current)
	assert.Equal(t, 1, len(item.Advisories))

	advisory := item.Advisories[0]
	assert.Equal(t, 9.8, advisory.CVSSScore)
	assert.Equal(t, "1067342", advisory.Identifier)
	assert.True(t, strings.HasPrefix(advisory.Overview, "Minimist <=1.2.5"))
	assert.Equal(t, ">=1.2.6", advisory.PatchedVersions)
	assert.Equal(t, "Upgrade to version 1.2.6 or later", advisory.Recommendation)
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

func TestInvalidJsonData(t *testing.T) {
	vector := `{
	"test": []
}`

	result, err := parseNpmOutput(vector)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unable to interpret this input")
	assert.Zero(t, len(result))
}

func TestListJsonOutputWithProblems(t *testing.T) {
	vector := `{
	"version": "3.12.0",
	"name": "foobar",
	"problems": [
		"extraneous: __ngcc_entry_points__.json@ /Users/testortestington/Documents/Repositories/foobar/frontend/node_modules/__ngcc_entry_points__.json"
	],
	"dependencies": {
		"__ngcc_entry_points__.json": {
			"extraneous": true,
			"problems": [
				"extraneous: __ngcc_entry_points__.json@ /Users/test/testortestington/Repositories/foobar/frontend/node_modules/__ngcc_entry_points__.json"
			]
		},
		"test": {
			"version": "0.1301.2",
			"resolved": "https://registry.npmjs.org/test/test/-/test-0.1301.2.tgz"
		}
	}
}`

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "__ngcc_entry_points__.json", result[0].Name)
	assert.Equal(t, "", result[0].Wanted)
	assert.Equal(t, "", result[0].Current)
	assert.Equal(t, "test", result[1].Name)
	assert.Equal(t, "0.1301.2", result[1].Wanted)
	assert.Equal(t, "", result[1].Current)
}

func TestParsePackageJson(t *testing.T) {
	vector :=
		`{
  "name": "foobar",
  "version": "3.12.0",
  "license": "MIT",
  "angular-cli": {},
  "lint-staged": {
    "*.ts": [
      "npm run lint:fix",
      "git add"
    ]
  },
  "browser": {
    "fs": false
  },
  "scripts": {
    "ng": "ng",
    "build": "ng build --configuration production --progress=false --configuration=production",
    "start": "ng serve --host 0.0.0.0",
    "test": "ng test",
    "style": "prettier --write 'src/**/*.ts'",
    "lint": "ng lint",
    "lint:fix": "ng lint --fix",
    "pree2e": "./node_modules/protractor/bin/webdriver-manager update",
    "e2e": "ng e2e",
    "aot": "node --max-old-space-size=8192 ./node_modules/@angular/cli/bin/ng build --configuration production --aot",
    "checkin": "git add dist",
    "pre-commit": "run-s lint test build checkin"
  },
  "private": true,
  "dependencies": {
    "@angular-redux/router": "^10.0.0",
    "zone.js": "~0.11.4"
  },
  "devDependencies": {
    "@angular-devkit/architect": "^0.1301.2",
    "@angular-devkit/build-angular": "^13.1.2",
    "selenium-webdriver": "^4.0.0-alpha.7"
  },
  "husky": {
    "hooks": {}
  }
}
`

	result, err := parseNpmOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 5, len(result))

	item := result[0]
	assert.Equal(t, "@angular-devkit/architect", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "^0.1301.2", item.Wanted)

	assert.Equal(t, 0, len(item.Installations))

	item = result[1]
	assert.Equal(t, "@angular-devkit/build-angular", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "^13.1.2", item.Wanted)

	item = result[2]
	assert.Equal(t, "@angular-redux/router", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "^10.0.0", item.Wanted)

	item = result[3]
	assert.Equal(t, "selenium-webdriver", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "^4.0.0-alpha.7", item.Wanted)

	item = result[4]
	assert.Equal(t, "zone.js", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "~0.11.4", item.Wanted)
}
