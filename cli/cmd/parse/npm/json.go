package npm

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-multierror"
	. "github.com/omniversion/omniversion/cli/cmd/parse/npm/types"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
)

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
		newResult := PackageMetadata{
			Name:   name,
			Wanted: version,
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

func parseAsPackageJson(packageJson *NpmPackageJson, result *[]PackageMetadata) *multierror.Error {
	var err *multierror.Error
	if len(packageJson.Dependencies) > 0 {
		err = multierror.Append(err, parsePackageJsonDependencies(packageJson.Dependencies, result))
	}
	if len(packageJson.DevDependencies) > 0 {
		err = multierror.Append(err, parsePackageJsonDependencies(packageJson.DevDependencies, result))
	}
	if len(packageJson.PeerDependencies) > 0 {
		err = multierror.Append(err, parsePackageJsonDependencies(packageJson.PeerDependencies, result))
	}
	return err
}

func parsePackageJsonDependencies(dependencies map[string]string, result *[]PackageMetadata) error {
	for name, version := range dependencies {
		newItem := PackageMetadata{
			Name:   name,
			Wanted: version,
		}
		if shared.InjectPackageManager {
			newItem.PackageManager = "npm"
		}
		*result = append(*result, newItem)
	}
	return nil
}
