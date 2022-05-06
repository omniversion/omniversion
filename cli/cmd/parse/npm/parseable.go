package npm

import (
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func parseAsList(input string, result *[]PackageMetadata) *multierror.Error {
	listRegex := regexp.MustCompile(`(?m)^(?P<location>[^\n:]*):(?P<wantedPackage>[^\n:]*)@(?P<wantedVersion>[^\n:]*):((?P<currentPackage>[^\n:]*)@(?P<currentVersion>[^\n:]*)|MISSING):(?P<latestPackage>[^\n:]*)@(?P<latestVersion>[^\n:]*)(:(?P<dir>.*))?$`)
	items := listRegex.FindAllStringSubmatch(input, -1)
	var allErrors *multierror.Error = nil
	for _, foundItem := range items {
		newItem := PackageMetadata{}
		if shared.InjectPackageManager {
			newItem.PackageManager = "npm"
		}
		currentVersion := foundItem[listRegex.SubexpIndex("currentVersion")]
		for groupIndex, groupName := range listRegex.SubexpNames() {
			if groupIndex != 0 && groupName != "" {
				value := foundItem[groupIndex]
				if len(value) > 0 {
					switch groupName {
					case "latestPackage":
						newItem.Name = value
					case "wantedVersion":
						newItem.Wanted = value
					case "currentVersion":
						newItem.Current = value
					case "latestVersion":
						newItem.Latest = strings.Trim(value, "\n")
					case "location":
						if currentVersion != "" {
							newItem.Installations = []InstalledPackage{{Location: value, Version: currentVersion}}
						}
					}
				}
			}
		}
		*result = append(*result, newItem)
	}
	return allErrors
}
