package apt

import (
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

func parseAptOutput(input string) ([]PackageMetadata, error) {
	compiledRegex := regexp.MustCompile(`(?m)^(?P<name>.*?)/(?P<sources>\S*) (?P<version>\S*) (?P<architecture>\S*)( \[(?P<installed>installed)?(,(?P<automatic>automatic))?(,upgradable to: (?P<latest>.*))?(upgradable from: (?P<outdatedVersion>.*))?])?$`)
	matches := compiledRegex.FindAllStringSubmatch(input, -1)
	result := make([]PackageMetadata, 0, len(matches))
	var allErrors *multierror.Error
	for _, match := range matches {
		newDependency := PackageMetadata{}
		if shared.InjectPackageManager {
			newDependency.PackageManager = "apt"
		}
		version := match[compiledRegex.SubexpIndex("version")]
		outdatedVersion := match[compiledRegex.SubexpIndex("outdatedVersion")]
		if outdatedVersion == "" {
			isInstalled := len(match[compiledRegex.SubexpIndex("installed")]) > 0
			latest := match[compiledRegex.SubexpIndex("latest")]
			newDependency.Latest = latest
			if isInstalled {
				newDependency.Current = version
				newDependency.Installations = []InstalledPackage{{
					Version: version,
				}}
			} else {
				newDependency.Wanted = version
			}
		} else {
			// we are dealing with the output of an `outdated` command
			newDependency.Current = outdatedVersion
			newDependency.Installations = []InstalledPackage{{
				Version: outdatedVersion,
			}}
			newDependency.Latest = version
		}
		for groupIndex, groupName := range compiledRegex.SubexpNames() {
			if groupIndex != 0 && groupName != "" {
				value := match[groupIndex]
				if len(value) > 0 {
					switch groupName {
					case "name":
						newDependency.Name = value
					case "architecture":
						newDependency.Architecture = value
					case "sources":
						newDependency.Sources = strings.Split(value, ",")
					}
				}
			}
		}
		result = append(result, newDependency)
	}
	return result, allErrors.ErrorOrNil()
}

var ParseCommand = &cobra.Command{
	Use:     "apt",
	Short:   "Parse the output of apt",
	Long:    `Translate the output of apt into the omniversion format.`,
	Run:     shared.WrapCommand(parseAptOutput),
	Aliases: []string{"apt-get", "aptitude"},
}
