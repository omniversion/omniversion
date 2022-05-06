package homebrew

import (
	"fmt"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func isVersionCommandOutput(input string) bool {
	return strings.HasPrefix(input, "Homebrew ")
}

func parseAsVersionCommandOutput(input string) ([]PackageMetadata, error) {
	compiledRegex := regexp.MustCompile(`(?m)^Homebrew (?P<version>\S*)$`)
	match := compiledRegex.FindStringSubmatch(input)
	if match == nil {
		// if the input is actually `rvm version` output,
		// we need to adapt the regex
		// if not, we don't need to worry about it and can keep trying to parse it in other formats
		return []PackageMetadata{}, fmt.Errorf("unexpected input format")
	} else {
		version := match[compiledRegex.SubexpIndex("version")]
		newItem := PackageMetadata{
			Name:          "homebrew",
			Current:       version,
			Installations: []InstalledPackage{{Version: version}},
		}
		return []PackageMetadata{newItem}, nil
	}
}
