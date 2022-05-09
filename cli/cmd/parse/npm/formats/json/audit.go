package json

import (
	"encoding/json"
	"fmt"
	. "github.com/omniversion/omniversion/cli/cmd/parse/npm/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	. "github.com/omniversion/omniversion/cli/types"
	"sort"
)

type V8AuditJsonOutput struct {
	AuditReportVersion int
	Vulnerabilities    map[string]struct {
		Name     string
		Severity string
		IsDirect bool
		// either an array of strings or an array of objects
		Via     []interface{}
		Effects []string
		Range   string
		Nodes   []string
		// either a boolean or an object
		FixAvailable interface{}
	}
	Metadata struct {
		Vulnerabilities struct {
			Info     int
			Low      int
			Moderate int
			High     int
			Critical int
			Total    int
		}
		Dependencies struct {
			Prod         int
			Dev          int
			Optional     int
			Peer         int
			PeerOptional int
			Total        int
		}
	}
}

type V6AuditJsonOutput struct {
	Actions []struct {
		IsMajor  bool
		Action   string
		Resolves []struct {
			Id       int
			Path     string
			Dev      bool
			Optional bool
			Bundled  bool
		}
		Module string
		Target string
		Depth  int
	}
	Advisories map[string]struct {
		Findings []struct {
			Version string
			Paths   []string
		}
		VulnerableVersions string `json:"vulnerable_versions"`
		ModuleName         string `json:"module_name"`
		Severity           string
		Access             string
		PatchedVersions    string `json:"patched_versions"`
		CVSS               struct {
			Score        float64
			VectorString string
		} `json:"cvss"`
		Recommendation string
		Id             int
		References     string
		Title          string
		Overview       string
		Url            string
	}
	Muted    []interface{}
	Metadata struct {
		Vulnerabilities struct {
			Info     int
			Low      int
			Moderate int
			High     int
			Critical int
		}
		Dependencies         int
		DevDependencies      int
		OptionalDependencies int
		TotalDependencies    int
	}
	RunId string
}

func ParseAuditOutput(input string, stderrOutput stderr.Output) ([]PackageMetadata, error) {
	if input == "" && stderrOutput.AuditLockfileMissing {
		return []PackageMetadata{}, fmt.Errorf("this command requires an existing lockfile")
	}

	v8AuditJson := V8AuditJsonOutput{}
	v8JsonUnmarshallErr := json.Unmarshal([]byte(input), &v8AuditJson)
	var result map[string]PackageMetadata
	if v8JsonUnmarshallErr == nil && v8AuditJson.AuditReportVersion != 0 {
		result = make(map[string]PackageMetadata, len(v8AuditJson.Vulnerabilities))
		for _, vulnerability := range v8AuditJson.Vulnerabilities {
			advisory := Advisory{
				Severity:           vulnerability.Severity,
				VulnerableVersions: vulnerability.Range,
			}
			if existingResult, ok := result[vulnerability.Name]; ok {
				existingResult.Advisories = append(existingResult.Advisories, advisory)
			} else {
				newItem := New(vulnerability.Name)
				newItem.Advisories = []Advisory{advisory}
				result[vulnerability.Name] = *newItem
			}
		}
	} else {
		v6AuditJson := V6AuditJsonOutput{}
		v6JsonUnmarshallErr := json.Unmarshal([]byte(input), &v6AuditJson)
		if v6JsonUnmarshallErr == nil {
			result = make(map[string]PackageMetadata, len(v6AuditJson.Advisories))
			for _, advisory := range v6AuditJson.Advisories {
				newAdvisory := Advisory{
					CVSSScore:          advisory.CVSS.Score,
					Identifier:         fmt.Sprintf("%v", advisory.Id),
					Overview:           advisory.Overview,
					PatchedVersions:    advisory.PatchedVersions,
					Recommendation:     advisory.Recommendation,
					Severity:           advisory.Severity,
					Title:              advisory.Title,
					Url:                advisory.Url,
					VulnerableVersions: advisory.VulnerableVersions,
				}

				if existingResult, ok := result[advisory.ModuleName]; ok {
					existingResult.Advisories = append(existingResult.Advisories, newAdvisory)
					result[advisory.ModuleName] = existingResult
				} else {
					newItem := New(advisory.ModuleName)
					newItem.Advisories = []Advisory{newAdvisory}
					result[advisory.ModuleName] = *newItem
				}
			}
		} else {
			// report the unmarshal error for v8, not the outdated format
			return []PackageMetadata{}, v8JsonUnmarshallErr
		}
	}
	// we were dealing with a map, so need to turn it into an array
	// and sort results to get consistent output
	resultsArray := make([]PackageMetadata, 0, len(result))
	for _, packageMetadata := range result {
		resultsArray = append(resultsArray, packageMetadata)
	}
	sort.Slice(resultsArray, func(i, j int) bool {
		return resultsArray[i].Name < resultsArray[j].Name
	})
	// also need to sort the advisories by identifier
	for _, item := range resultsArray {
		sort.Slice(item.Advisories, func(i, j int) bool {
			return item.Advisories[i].Identifier < item.Advisories[j].Identifier
		})
	}
	return resultsArray, nil
}
