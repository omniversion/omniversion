package formats

import (
	"encoding/json"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"regexp"
	"strings"
)

type InputFormat string
type Verb string

const (
	DefaultFormat   InputFormat = "default"
	JsonFormat      InputFormat = "json"
	ParseableFormat InputFormat = "parseable"
)

const (
	AuditCommand        Verb = "audit"
	ListCommand         Verb = "list"
	OutdatedCommand     Verb = "outdated"
	VersionCommand      Verb = "version"
	PackageJsonFile     Verb = "package.json"
	PackageLockJsonFile Verb = "package-lock.json"
)

func DetectVerbAndFormat(input string, stderrOutput stderr.Output) (Verb, InputFormat) {
	if strings.HasPrefix(input, "# npm audit report") {
		return AuditCommand, DefaultFormat
	}
	if regexp.MustCompile("(?m)^└──|^└─┬|^`--|^`-+").MatchString(input) {
		return ListCommand, DefaultFormat
	}
	if stderrOutput.AuditLockfileMissing {
		if stderrOutput.Json != nil {
			return AuditCommand, JsonFormat
		} else {
			return AuditCommand, DefaultFormat
		}
	}
	var jsonData map[string]interface{}
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &jsonData)
	if jsonUnmarshallErr == nil {
		if _, ok := jsonData["lockfileVersion"]; ok {
			return PackageLockJsonFile, JsonFormat
		}
		if _, ok := jsonData["auditReportVersion"]; ok {
			return AuditCommand, JsonFormat
		}
		if _, ok := jsonData["advisories"]; ok {
			return AuditCommand, JsonFormat
		}
		if _, ok := jsonData["vulnerabilities"]; ok {
			return AuditCommand, JsonFormat
		}
		if _, ok := jsonData["scripts"]; ok {
			return PackageJsonFile, JsonFormat
		}
		if _, ok := jsonData["license"]; ok {
			return PackageJsonFile, JsonFormat
		}
		if _, ok := jsonData["dependencies"]; ok {
			return ListCommand, JsonFormat
		}
		if npmItem, ok := jsonData["npm"]; ok {
			if _, ok := npmItem.(string); ok {
				return VersionCommand, JsonFormat
			}
		}
		return OutdatedCommand, JsonFormat
	}
	if strings.HasPrefix(input, "{") {
		// starting with `{`, but invalid JSON
		// we are probably dealing with the default `npm --versions` command
		// which outputs an object with unquoted keys
		// that can't be parsed as JSON (but can be parsed as YAML)
		return VersionCommand, DefaultFormat
	}
	firstLine := strings.Split(input, "\n")[0]
	if strings.Contains(firstLine, "Package") &&
		strings.Contains(firstLine, "Current") &&
		strings.Contains(firstLine, "Wanted") &&
		strings.Contains(firstLine, "Latest") {
		return OutdatedCommand, DefaultFormat
	}
	if len(strings.Split(firstLine, ":")) >= 4 {
		return OutdatedCommand, ParseableFormat
	}
	return ListCommand, ParseableFormat
}
