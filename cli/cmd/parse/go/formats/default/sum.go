package _default

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/go/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func ParseGoSumFile(input string) ([]PackageMetadata, error) {
	lines := strings.Split(input, "\n")
	lineRegex := regexp.MustCompile(`^(?P<name>\S+) (?P<version>\S+) (?P<hash>\S+)$`)
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
		version := match[lineRegex.SubexpIndex("version")]
		// ignore entries for `go.mod` files, we only want the module's `zip` file
		if strings.HasSuffix(version, "/go.mod") {
			continue
		}
		name := match[lineRegex.SubexpIndex("name")]
		hash := match[lineRegex.SubexpIndex("hash")]
		if name != "" && version != "" && hash != "" {
			newItem := item.New(name)
			newItem.Aliases = []string{shared.ShortModuleName(name)}
			newItem.Current = shared.CleanVersion(version)
			newItem.Installations = []InstalledPackage{{
				Version: version,
				Hash:    hash,
			}}
			results = append(results, *newItem)
		}
	}

	return results, errors.ErrorOrNil()
}
