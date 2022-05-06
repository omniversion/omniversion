package npm

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
)

func stripProblems(input string, result *[]PackageMetadata) (string, *multierror.Error) {
	problemRegex := regexp.MustCompile(`(?m)npm ERR! (?P<problem>[^:]*): (?P<name>\S*)@(?P<version>[^\s,]*)(, required by (?P<requiredBy>\S*))?( (?P<location>.*))?`)
	var allErrors *multierror.Error
	foundProblems := problemRegex.FindAllStringSubmatch(input, -1)
	for _, foundProblem := range foundProblems {
		newItem := PackageMetadata{}
		if shared.InjectPackageManager {
			newItem.PackageManager = "npm"
		}
		problemKind := foundProblem[problemRegex.SubexpIndex("problem")]
		for groupIndex, groupName := range problemRegex.SubexpNames() {
			if groupIndex != 0 && groupName != "" {
				value := foundProblem[groupIndex]
				if len(value) > 0 {
					switch groupName {
					case "name":
						newItem.Name = value
					case "version":
						switch problemKind {
						case "missing":
							newItem.Wanted = value
						case "extraneous":
							newItem.Current = value
						default:
							allErrors = multierror.Append(allErrors, fmt.Errorf("unknown npm problem kind: %q", problemKind))
						}
					case "location":
						newItem.Installations = []InstalledPackage{{Location: value}}
					}
				}
			}
		}
		*result = append(*result, newItem)
	}

	strippedInput := problemRegex.ReplaceAllLiteralString(input, "")
	return strippedInput, allErrors
}
