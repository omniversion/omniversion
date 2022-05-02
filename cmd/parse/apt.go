package parse

import (
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion-cli/models"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

func parseAptOutput(input string) ([]models.Dependency, error) {
	compiledRegex := regexp.MustCompile(`(?m)^(?P<name>.*?)/(?P<sources>\S*) (?P<version>\S*) (?P<architecture>\S*)( \[(?P<installed>installed)?(,(?P<automatic>automatic))?(,upgradable to: (?P<latest>.*))?(upgradable from: (?P<outdatedVersion>.*))?])?$`)
	matches := compiledRegex.FindAllStringSubmatch(input, -1)
	result := make([]models.Dependency, 0, len(matches))
	var allErrors *multierror.Error
	for _, match := range matches {
		newDependency := models.Dependency{
			Pm: "apt",
		}
		version := match[compiledRegex.SubexpIndex("version")]
		outdatedVersion := match[compiledRegex.SubexpIndex("outdatedVersion")]
		if outdatedVersion == "" {
			isInstalled := len(match[compiledRegex.SubexpIndex("installed")]) > 0
			latest := match[compiledRegex.SubexpIndex("latest")]
			newDependency.Latest = latest
			if isInstalled {
				newDependency.Version = version
				newDependency.Installed = []models.InstalledDependency{{
					Version: version,
				}}
			} else {
				newDependency.Wanted = version
			}
		} else {
			// we are dealing with the output of an `outdated` command
			newDependency.Version = outdatedVersion
			newDependency.Installed = []models.InstalledDependency{{
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

var AptCmd = &cobra.Command{
	Use:   "apt",
	Short: "Parse the output of apt",
	Long:  `Transform the output of apt into a common format.`,
	Run:   wrapCommand(parseAptOutput),
}
