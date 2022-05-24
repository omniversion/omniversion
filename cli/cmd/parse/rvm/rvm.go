package rvm

import (
	"fmt"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
	"regexp"
)

func parseRvmOutput(input string) ([]PackageMetadata, error) {
	result, err := parseAsRvmVersionOutput(input)
	if err == nil {
		return result, nil
	}

	compiledRegex := regexp.MustCompile(`(?m)^(?P<current>=)? ?(?P<default>\*)?([ >])? *ruby-(?P<version>[^ ]*) \[ (?P<architecture>.*) ]$`)
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

func parseAsRvmVersionOutput(input string) ([]PackageMetadata, error) {
	compiledRegex := regexp.MustCompile(`(?m)^rvm (?P<version>\S*)(?P<latest> \(latest\))?(?P<manual> \(manual\))? by (?P<author>.*) \[(?P<url>.*)]$`)
	match := compiledRegex.FindStringSubmatch(input)
	if match == nil {
		// if the input is actually `rvm version` output,
		// we need to adapt the regex
		// if not, we don't need to worry about it and can keep trying to parse it in other formats
		return []PackageMetadata{}, fmt.Errorf("unexpected input format")
	} else {
		version := match[compiledRegex.SubexpIndex("version")]
		latest := match[compiledRegex.SubexpIndex("latest")]
		author := match[compiledRegex.SubexpIndex("author")]
		url := match[compiledRegex.SubexpIndex("url")]
		newItem := PackageMetadata{
			Name:     "rvm",
			Author:   author,
			Current:  version,
			Homepage: url,
		}
		if latest != "" {
			newItem.Latest = version
		}
		return []PackageMetadata{newItem}, nil
	}
}

var ParseCommand = &cobra.Command{
	Use:   "rvm",
	Short: "Parse the output of rvm",
	Long:  `Translate the output of rvm into the omniversion format.`,
	Run:   shared.WrapCommand(parseRvmOutput),
}
