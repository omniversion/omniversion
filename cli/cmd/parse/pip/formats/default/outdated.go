package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/item"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func ParseOutdatedOutput(input string) ([]PackageMetadata, error) {
	lines := strings.Split(input, "\n")
	lineRegex := regexp.MustCompile(`^(?P<name>\S+)\s+(?P<version>\S+)\s+(?P<latest>\S+)\s+(?P<type>\S+)(\s+(?P<location>\S+))?$`)
	results := make([]PackageMetadata, 0, len(lines))

	// skip the first two lines
	for _, line := range lines[2:] {
		if line != "" {
			match := lineRegex.FindStringSubmatch(line)
			name := match[lineRegex.SubexpIndex("name")]
			version := match[lineRegex.SubexpIndex("version")]
			latest := match[lineRegex.SubexpIndex("latest")]
			location := match[lineRegex.SubexpIndex("location")]
			if name != "" && version != "" {
				newItem := item.New(name)
				newItem.Current = version
				newItem.Latest = latest
				newItem.Installations = []InstalledPackage{{
					Version:  version,
					Location: location,
				}}
				results = append(results, *newItem)
			}
		}
	}

	return results, nil
}
