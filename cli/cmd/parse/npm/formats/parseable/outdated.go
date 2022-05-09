package parseable

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func ParseOutdatedOutput(input string, _ stderr.Output) ([]PackageMetadata, error) {
	var result []PackageMetadata
	listRegex := regexp.MustCompile(`(?m)^(?P<location>[^\n:]*):(?P<wantedPackage>[^\n:]*)@(?P<wantedVersion>[^\n:]*):((?P<currentPackage>[^\n:]*)@(?P<currentVersion>[^\n:]*)|MISSING):(?P<latestPackage>[^\n:]*)@(?P<latestVersion>[^\n:]*)(:(?P<dir>.*))?$`)
	items := listRegex.FindAllStringSubmatch(input, -1)
	for _, foundItem := range items {
		name := foundItem[listRegex.SubexpIndex("latestPackage")]
		newItem := item.New(name)
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
				} else {
					if groupName == "currentVersion" {
						isMissing := true
						newItem.Missing = &isMissing
					}
				}
			}
		}
		result = append(result, *newItem)
	}
	return result, nil
}
