package freeze

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
	lineRegex := regexp.MustCompile(`^(?P<name>\S+)==(?P<version>\S+)$`)
	results := make([]PackageMetadata, 0, len(lines))

	var errors *multierror.Error
	for _, line := range lines {
		if line == "" {
			continue
		}
		match := lineRegex.FindStringSubmatch(line)
		if match == nil {
			errors = multierror.Append(errors, fmt.Errorf("failed to parse line: %q", line))
			continue
		}
		name := match[lineRegex.SubexpIndex("name")]
		version := match[lineRegex.SubexpIndex("version")]
		if name != "" && version != "" {
			newItem := item.New(name)
			newItem.Current = version
			newItem.Installations = []InstalledPackage{{
				Version: version,
			}}
			results = append(results, *newItem)
		}
	}

	return results, errors.ErrorOrNil()
}
