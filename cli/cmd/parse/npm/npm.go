package npm

import (
	"encoding/json"
	"github.com/hashicorp/go-multierror"
	. "github.com/omniversion/omniversion/cli/cmd/parse/npm/types"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"sort"
)

func parseNpmOutput(input string) ([]PackageMetadata, error) {
	result := make([]PackageMetadata, 0, 100)
	// remove problems that might appear in stderr
	// and would prevent us from parsing the content as JSON
	// this is relevant if stdout and stderr have been merged,
	// e.g. in terminal output copied from the console
	var allErrors *multierror.Error
	input, err := stripProblems(input, &result)
	if err != nil {
		allErrors = multierror.Append(allErrors, err)
	}

	if isInDefaultAuditFormat(input) {
		err = parseAsDefaultAuditOutput(input, &result)
		sort.Slice(result, func(i, j int) bool {
			return result[i].Name < result[j].Name
		})
		return result, multierror.Append(allErrors, err).ErrorOrNil()
	}

	if isInTreeFormat(input) {
		err = parseAsTree(input, &result)
		sort.Slice(result, func(i, j int) bool {
			return result[i].Name < result[j].Name
		})
		return result, multierror.Append(allErrors, err).ErrorOrNil()
	}

	if isInOutdatedTableFormat(input) {
		err = parseAsOutdatedTable(input, &result)
		sort.Slice(result, func(i, j int) bool {
			return result[i].Name < result[j].Name
		})
		return result, multierror.Append(allErrors, err).ErrorOrNil()
	}

	// we might have a list of strings in npm `--parseable` format
	// or valid JSON - so we try to unmarshall it
	packageJson := &NpmPackageJson{}
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &packageJson)

	if jsonUnmarshallErr == nil && packageJson.Version != "" && packageJson.Name != "" {
		err = parseAsPackageJson(packageJson, &result)
		sort.Slice(result, func(i, j int) bool {
			return result[i].Name < result[j].Name
		})
		return result, multierror.Append(allErrors, err).ErrorOrNil()
	}

	// the keys in the default `npm version` output are unquoted,
	// but the yaml parser can deal with this
	npmVersionData := &NpmVersionJson{}
	yamlUnmarshallErr := yaml.Unmarshal([]byte(input), &npmVersionData)
	if yamlUnmarshallErr == nil {
		if _, ok := (*npmVersionData)["npm"]; ok {
			err = parseAsNpmJson(*npmVersionData, &result)
			sort.Slice(result, func(i, j int) bool {
				return result[i].Name < result[j].Name
			})
			return result, multierror.Append(allErrors, err).ErrorOrNil()
		}
	}

	// we might have a list of strings in npm `--parseable` format
	// or valid JSON - so we try to unmarshall it
	dependenciesAsJson := &NpmJson{}
	jsonUnmarshallErr = json.Unmarshal([]byte(input), &dependenciesAsJson)

	if jsonUnmarshallErr == nil {
		err = parseAsJson(input, *dependenciesAsJson, &result)
		sort.Slice(result, func(i, j int) bool {
			return result[i].Name < result[j].Name
		})
		return result, multierror.Append(allErrors, err).ErrorOrNil()
	} else {
		err = parseAsList(input, &result)
		sort.Slice(result, func(i, j int) bool {
			return result[i].Name < result[j].Name
		})
		return result, multierror.Append(allErrors, err).ErrorOrNil()
	}
}

var ParseCommand = &cobra.Command{
	Use:   "npm",
	Short: "Parse the output of npm",
	Long:  `Translate the output of npm into the omniversion format.`,
	Run:   shared.WrapCommand(parseNpmOutput),
}
