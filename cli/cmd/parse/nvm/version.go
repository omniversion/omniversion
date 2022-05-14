package nvm

import (
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

const versionOutputRegex = `(?m)^(?P<version>\S+)$`

func isVersionOutput(input string) bool {
	return strings.Count(input, "\n") == 1 && regexp.MustCompile(versionOutputRegex).MatchString(input)
}

func parseAsVersionOutput(input string) ([]PackageMetadata, error) {
	compiledRegex := regexp.MustCompile(versionOutputRegex)
	match := compiledRegex.FindStringSubmatch(input)
	version := match[compiledRegex.SubexpIndex("version")]
	return []PackageMetadata{{
		Name:           "nvm",
		Current:        version,
		PackageManager: "nvm",
		Installations: []InstalledPackage{{
			Version: version,
		}},
	}}, nil
}
