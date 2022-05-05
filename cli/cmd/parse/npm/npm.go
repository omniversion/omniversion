package npm

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-multierror"
	. "github.com/omniversion/omniversion/cli/cmd/parse/npm/types"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
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

	// we might have a list of strings in npm `--parseable` format
	// or valid JSON - so we try to unmarshall it
	dependenciesAsJson := &NpmJson{}
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &dependenciesAsJson)

	if jsonUnmarshallErr == nil {
		err = parseAsJson(input, *dependenciesAsJson, &result)
		return result, multierror.Append(allErrors, err).ErrorOrNil()
	} else {
		err = parseAsList(input, &result)
		return result, multierror.Append(allErrors, err).ErrorOrNil()
	}
}

func parseAsList(input string, result *[]PackageMetadata) *multierror.Error {
	listRegex := regexp.MustCompile(`(?m)^(?P<location>[^\n:]*):(?P<wantedPackage>[^\n:]*)@(?P<wantedVersion>[^\n:]*):((?P<currentPackage>[^\n:]*)@(?P<currentVersion>[^\n:]*)|MISSING):(?P<latestPackage>[^\n:]*)@(?P<latestVersion>[^\n:]*)(:(?P<dir>.*))?$`)
	items := listRegex.FindAllStringSubmatch(input, -1)
	var allErrors *multierror.Error = nil
	for _, foundItem := range items {
		newItem := PackageMetadata{}
		if shared.InjectPackageManager {
			newItem.PackageManager = "npm"
		}
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
				}
			}
		}
		*result = append(*result, newItem)
	}
	return allErrors
}

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

func parseAsJson(input string, dependenciesAsJson NpmJson, result *[]PackageMetadata) *multierror.Error {
	if len(dependenciesAsJson.Dependencies) > 0 {
		return parseJsonDependencies(dependenciesAsJson.Dependencies, result)
	}
	if len(dependenciesAsJson.Advisories) > 0 {
		return parseJsonAdvisories(dependenciesAsJson.Advisories, result)
	}
	npmVersionData := &NpmVersionJson{}
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &npmVersionData)
	if jsonUnmarshallErr == nil {
		if _, ok := (*npmVersionData)["npm"]; ok {
			return parseAsNpmJson(*npmVersionData, result)
		}
	}

	flatJsonData := &NpmFlatJson{}
	jsonUnmarshallErr = json.Unmarshal([]byte(input), &flatJsonData)
	if jsonUnmarshallErr == nil {
		return parseAsFlatJson(*flatJsonData, result)
	}
	var multiError *multierror.Error = nil
	return multierror.Append(multiError, fmt.Errorf("unable to interpret this input: %q", input))
}

func parseJsonDependencies(dependencyData map[string]NpmDependency, result *[]PackageMetadata) *multierror.Error {
	var allErrors *multierror.Error = nil
	for name, dependency := range dependencyData {
		version := dependency.Version
		if version == "" {
			allErrors = multierror.Append(allErrors, fmt.Errorf("no version found: %q", name))
			continue
		}
		newResult := PackageMetadata{
			Name:    name,
			Current: version,
		}
		if shared.InjectPackageManager {
			newResult.PackageManager = "npm"
		}
		*result = append(*result, newResult)
	}
	return allErrors
}

func parseJsonAdvisories(advisoryData map[string]NpmAdvisory, result *[]PackageMetadata) *multierror.Error {
	var allErrors *multierror.Error = nil
	for _, advisory := range advisoryData {
		newDependency := PackageMetadata{
			Name: advisory.ModuleName,
			Advisories: []Advisory{{
				CVSSScore:          advisory.CVSS.Score,
				Identifier:         fmt.Sprintf("%v", advisory.Id),
				Overview:           advisory.Overview,
				PatchedVersions:    advisory.PatchedVersions,
				Recommendation:     advisory.Recommendation,
				Severity:           advisory.Severity,
				Title:              advisory.Title,
				Url:                advisory.Url,
				VulnerableVersions: advisory.VulnerableVersions,
			}},
		}
		if shared.InjectPackageManager {
			newDependency.PackageManager = "npm"
		}
		if len(advisory.Findings) > 0 {
			newDependency.Current = advisory.Findings[0].Version
		}
		*result = append(*result, newDependency)
	}
	return allErrors
}

func parseAsNpmJson(dependenciesAsNpmVersionJson NpmVersionJson, result *[]PackageMetadata) *multierror.Error {
	var allErrors *multierror.Error = nil
	for packageName, version := range dependenciesAsNpmVersionJson {
		newResult := PackageMetadata{
			Name:    packageName,
			Current: version,
		}
		if shared.InjectPackageManager {
			newResult.PackageManager = "npm"
		}
		*result = append(*result, newResult)
	}
	return allErrors
}

func parseAsFlatJson(dependenciesAsJson NpmFlatJson, result *[]PackageMetadata) *multierror.Error {
	var allErrors *multierror.Error = nil
	for packageName, details := range dependenciesAsJson {
		newResult := PackageMetadata{
			Name:    packageName,
			Current: details.Current,
			Wanted:  details.Wanted,
			Latest:  details.Latest,
		}
		if shared.InjectPackageManager {
			newResult.PackageManager = "npm"
		}
		*result = append(*result, newResult)
	}
	return allErrors
}

var ParseCommand = &cobra.Command{
	Use:   "npm",
	Short: "Parse the output of npm",
	Long:  `Translate the output of npm into the omniversion format.`,
	Run:   shared.WrapCommand(parseNpmOutput),
}
