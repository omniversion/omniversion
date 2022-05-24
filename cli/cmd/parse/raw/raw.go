package raw

import (
	"fmt"
	"github.com/omniversion/omniversion/cli/cmd/parse/raw/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
	"regexp"
)

var name = ""
var regex = ""

func parseRawOutput(input string) ([]PackageMetadata, error) {
	if regex == "" {
		newItem := item.New(name)
		newItem.Current = shared.CleanVersion(input)
		return []PackageMetadata{*newItem}, nil
	} else {
		compiledRegex, err := regexp.Compile(regex)
		if err != nil {
			return []PackageMetadata{}, fmt.Errorf("invalid regex: %v", err)
		}
		matches := compiledRegex.FindAllStringSubmatch(input, -1)
		results := make([]PackageMetadata, 0, len(matches))
		for _, match := range matches {
			packageName := name
			if compiledRegex.SubexpIndex("name") != -1 {
				packageName = match[compiledRegex.SubexpIndex("name")]
			}
			newItem := item.New(packageName)
			packageVersion := input
			if compiledRegex.SubexpIndex("version") != -1 {
				packageVersion = match[compiledRegex.SubexpIndex("version")]
			}
			newItem.Current = shared.CleanVersion(packageVersion)
			results = append(results, *newItem)
		}
		return results, nil
	}
}

var ParseCommand = func() *cobra.Command {
	parseCommand := &cobra.Command{
		Use:   "raw",
		Short: "Parse raw output",
		Long:  `Translate arbitrary output into the omniversion format.`,
		Run:   shared.WrapCommand(parseRawOutput),
	}
	parseCommand.PersistentFlags().StringVarP(&regex, "regex", "r", "", "An optional GoLang-style regex containing a named groups `name` and `version`. The `name` group can be omitted if a --name parameter is provided. If no regex is given, the entire string will be used as the version.")
	parseCommand.PersistentFlags().StringVarP(&name, "name", "n", "", "The name of the package. This is required if the regex contains no `name` group.")
	return parseCommand
}()
