package rubygems

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

func parseRubygemsOutput(input string) ([]Dependency, error) {
	extractionRegex := regexp.MustCompile(`(?m)(?P<name>.*) \((?P<versions>.*)\)(\n(?P<content>(^ +.*\n?|^\n)*))?`)
	items := extractionRegex.FindAllStringSubmatch(input, -1)

	result := make([]Dependency, 0, len(items))
	var allErrors *multierror.Error
	for _, item := range items {
		name := item[extractionRegex.SubexpIndex("name")]
		versions := item[extractionRegex.SubexpIndex("versions")]
		content := item[extractionRegex.SubexpIndex("content")]

		if content == "" {
			err := parseListItem(name, versions, &result)
			if err != nil {
				allErrors = multierror.Append(allErrors, *err)
			}
		} else {
			err := parseDetails(name, versions, content, &result)
			if err != nil {
				allErrors = multierror.Append(allErrors, *err)
			}
		}
	}
	return result, allErrors.ErrorOrNil()
}

func parseListItem(name string, versions string, dependencies *[]Dependency) *error {
	versionComponents := strings.Split(versions, " < ")
	if len(versionComponents) != 2 {
		err := fmt.Errorf("unable to parse package description: %q", name)
		return &err
	}

	currentVersion := versionComponents[0]
	latestVersion := versionComponents[1]

	newResult := Dependency{
		Pm:        "rubygems",
		Name:      name,
		Version:   currentVersion,
		Latest:    latestVersion,
		Installed: []InstalledDependency{{Version: currentVersion}},
	}
	*dependencies = append(*dependencies, newResult)
	return nil
}

func parseDetails(name string, versions string, content string, dependencies *[]Dependency) *error {
	parseRegex := regexp.MustCompile(`(?m)\s+Authors?: (?P<authors>(.+\n)+)\s+Homepage: (?P<homepage>.+)\n\s+Licenses?: (?P<license>.+)\n\s+Installed at ?(?P<locations>(.+\n)+)\n(?P<description>(\n?.+)+)`)
	groupNames := parseRegex.SubexpNames()

	parsedContent := parseRegex.FindStringSubmatch(content)

	newResult := Dependency{
		Pm:   "rubygems",
		Name: name,
	}
	if !strings.Contains(versions, ",") {
		// if there is only a single, this must be the current installation
		newResult.Version = versions
	}

	if len(parsedContent) < len(groupNames) {
		err := fmt.Errorf("unable to parse package description: %q", name)
		return &err
	}

	for j, group := range groupNames {
		if j != 0 && group != "" && len(parsedContent[j]) > 0 {
			switch group {
			case "authors":
				newResult.Author = strings.Trim(parsedContent[j], "\n")
			case "homepage":
				newResult.Homepage = strings.Trim(parsedContent[j], "\n")
			case "license":
				newResult.License = parsedContent[j]
			case "locations":
				parseLocations(parsedContent[j], &newResult)
			case "description":
				newResult.Description = parseDescription(parsedContent[j])
			}
		}
	}
	*dependencies = append(*dependencies, newResult)
	return nil
}

func parseLocations(locationsData string, dependency *Dependency) {
	locationsRegex := regexp.MustCompile(`(?m)^\s*(\((?P<version>.*)\))?: (?P<location>.*)`)
	installedLocationData := locationsRegex.FindAllStringSubmatch(locationsData, -1)
	var installations []InstalledDependency
	for _, installedLocation := range installedLocationData {
		var newInstallation InstalledDependency
		parseVersion(installedLocation[locationsRegex.SubexpIndex("version")], dependency, &newInstallation)

		newInstallation.Location = installedLocation[locationsRegex.SubexpIndex("location")]
		installations = append(installations, newInstallation)
	}
	dependency.Installed = installations
}

func parseVersion(versionData string, dependency *Dependency, installedDependency *InstalledDependency) {
	// could be either empty or a version or "default" or both, separated by a comma
	versionString := versionData
	versionComponents := strings.Split(versionString, ", ")
	installedDependency.Version = versionComponents[0]
	if len(versionComponents) == 2 && versionComponents[1] == "default" {
		dependency.Default = versionComponents[0]
	}
	if len(versionComponents) == 1 && versionComponents[0] == "default" {
		dependency.Default = dependency.Version
	}
}

func parseDescription(descriptionData string) string {
	components := strings.Split(descriptionData, "\n")
	var result = []string{}
	for _, component := range components {
		trimmedComponent := strings.Trim(component, " ")
		if len(trimmedComponent) > 0 {
			result = append(result, trimmedComponent)
		}
	}
	return strings.Join(result, "\n")
}

var ParseCommand = &cobra.Command{
	Use:   "rubygems",
	Short: "Parse the output of rubygems",
	Long:  `Transform the output of rubygems into a common format.`,
	Run:   shared.WrapCommand(parseRubygemsOutput),
}