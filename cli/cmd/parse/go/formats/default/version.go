package _default

import (
	"fmt"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
)

func ParseVersionOutput(input string) ([]PackageMetadata, error) {
	versionRegex := regexp.MustCompile(`(?m)^go version go(?P<version>\S+) (?P<arch>\S+)$`)

	match := versionRegex.FindStringSubmatch(input)
	if match == nil {
		return []PackageMetadata{}, fmt.Errorf("failed to parse version data: %q", input)
	} else {
		version := match[versionRegex.SubexpIndex("version")]
		arch := match[versionRegex.SubexpIndex("arch")]
		return []PackageMetadata{{
			Name:         "go",
			Current:      version,
			Architecture: arch,
			Installations: []InstalledPackage{{
				Version: version,
			}},
		}}, nil
	}
}
