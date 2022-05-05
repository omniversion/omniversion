package rvm

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
	"regexp"
)

func parseRvmOutput(input string) ([]PackageMetadata, error) {
	compiledRegex := regexp.MustCompile("(?m)^(?P<current>=)? ?(?P<default>\\*)?([ >])? *ruby-(?P<version>[^ ]*) \\[ (?P<architecture>.*) ]$")
	matches := compiledRegex.FindAllStringSubmatch(input, -1)
	newItem := PackageMetadata{
		Name:           "ruby",
		PackageManager: "rvm",
	}
	installed := make([]InstalledPackage, 0, len(matches))
	for _, match := range matches {
		isCurrent := len(match[compiledRegex.SubexpIndex("current")]) > 0
		isDefault := len(match[compiledRegex.SubexpIndex("default")]) > 0
		version := match[compiledRegex.SubexpIndex("version")]
		installed = append(installed, InstalledPackage{Version: version})
		if isCurrent {
			newItem.Current = version
			newItem.Architecture = match[compiledRegex.SubexpIndex("architecture")]
		}
		if isDefault {
			newItem.Default = version
		}
	}
	newItem.Installations = installed
	return []PackageMetadata{newItem}, nil
}

var ParseCommand = &cobra.Command{
	Use:   "rvm",
	Short: "Parse the output of rvm",
	Long:  `Translate the output of rvm into the omniversion format.`,
	Run:   shared.WrapCommand(parseRvmOutput),
}
