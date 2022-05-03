package rvm

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
	"regexp"
)

func parseRvmOutput(input string) ([]Dependency, error) {
	compiledRegex := regexp.MustCompile("(?m)^(?P<current>=)? ?(?P<default>\\*)?([ >])? *ruby-(?P<version>[^ ]*) \\[ (?P<architecture>.*) ]$")
	matches := compiledRegex.FindAllStringSubmatch(input, -1)
	newItem := Dependency{
		Name: "ruby",
		Pm:   "rvm",
	}
	installed := make([]InstalledDependency, 0, len(matches))
	for _, match := range matches {
		isCurrent := len(match[compiledRegex.SubexpIndex("current")]) > 0
		isDefault := len(match[compiledRegex.SubexpIndex("default")]) > 0
		version := match[compiledRegex.SubexpIndex("version")]
		installed = append(installed, InstalledDependency{Version: version})
		if isCurrent {
			newItem.Version = version
			newItem.Architecture = match[compiledRegex.SubexpIndex("architecture")]
		}
		if isDefault {
			newItem.Default = version
		}
	}
	newItem.Installed = installed
	return []Dependency{newItem}, nil
}

var ParseCommand = &cobra.Command{
	Use:   "rvm",
	Short: "Parse the output of rvm",
	Long:  `Transform the output of rvm into a common format.`,
	Run:   shared.WrapCommand(parseRvmOutput),
}
