package npm

import (
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func isInDefaultAuditFormat(input string) bool {
	return strings.HasPrefix(input, "# npm audit report")
}

func parseAsDefaultAuditOutput(input string, result *[]PackageMetadata) *multierror.Error {
	listRegex := regexp.MustCompile(`(?m)^(?P<name>\S+) {2}(?P<version>.*)\n\s*(Severity: (?P<severity>.*)|Depends on vulnerable versions of (?P<dependency>.*))`)
	items := listRegex.FindAllStringSubmatch(input, -1)
	for _, foundItem := range items {
		name := foundItem[listRegex.SubexpIndex("name")]
		version := foundItem[listRegex.SubexpIndex("version")]
		severity := foundItem[listRegex.SubexpIndex("severity")]
		newItem := PackageMetadata{
			Name: name,
			Advisories: []Advisory{{
				Severity:           severity,
				VulnerableVersions: version,
			}},
		}
		if shared.InjectPackageManager {
			newItem.PackageManager = "npm"
		}
		*result = append(*result, newItem)
	}
	return nil
}
