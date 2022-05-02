package parse

import (
	models2 "github.com/omniversion/omniversion/cli/models"
	"github.com/spf13/cobra"
	"regexp"
)

func parseRvmOutput(input string) ([]models2.Dependency, error) {
	compiledRegex := regexp.MustCompile("(?m)^(?P<current>=)? ?(?P<default>\\*)?([ >])? *ruby-(?P<version>[^ ]*) \\[ (?P<architecture>.*) ]$")
	matches := compiledRegex.FindAllStringSubmatch(input, -1)
	newItem := models2.Dependency{
		Name: "ruby",
		Pm:   "rvm",
	}
	installed := make([]models2.InstalledDependency, 0, len(matches))
	for _, match := range matches {
		isCurrent := len(match[compiledRegex.SubexpIndex("current")]) > 0
		isDefault := len(match[compiledRegex.SubexpIndex("default")]) > 0
		version := match[compiledRegex.SubexpIndex("version")]
		installed = append(installed, models2.InstalledDependency{Version: version})
		if isCurrent {
			newItem.Version = version
			newItem.Architecture = match[compiledRegex.SubexpIndex("architecture")]
		}
		if isDefault {
			newItem.Default = version
		}
	}
	newItem.Installed = installed
	return []models2.Dependency{newItem}, nil
}

var RvmCmd = &cobra.Command{
	Use:   "rvm",
	Short: "Parse the output of rvm",
	Long:  `Transform the output of rvm into a common format.`,
	Run:   wrapCommand(parseRvmOutput),
}
