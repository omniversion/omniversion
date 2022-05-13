package _default

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/item"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func ParseListOutput(input string) ([]PackageMetadata, error) {
	lines := strings.Split(input, "\n")
	lineRegex := regexp.MustCompile(`^(?P<name>\S+)\s+(?P<version>\S+)(\s+(?P<location>\S+))?$`)
	results := make([]PackageMetadata, 0, len(lines))

	// skip the first two lines
	var errors *multierror.Error
	for _, line := range lines[2:] {
		if line != "" {
			match := lineRegex.FindStringSubmatch(line)
			if match == nil {
				errors = multierror.Append(errors, fmt.Errorf("failed to parse line: %q", line))
				continue
			}
			name := match[lineRegex.SubexpIndex("name")]
			version := match[lineRegex.SubexpIndex("version")]
			location := match[lineRegex.SubexpIndex("location")]
			if name != "" && version != "" {
				newItem := item.New(name)
				newItem.Current = version
				newItem.Installations = []InstalledPackage{{
					Version:  version,
					Location: location,
				}}
				results = append(results, *newItem)
			}
		}
	}

	return results, errors.ErrorOrNil()
}
