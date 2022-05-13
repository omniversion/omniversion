package nvm

import (
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func parseAsListOutput(input string) ([]PackageMetadata, error) {
	installedVersionsRegex := regexp.MustCompile(`(?m)^(?P<arrow>->)?\s+(v?(?P<version>\S*))$`)
	installedVersionMatches := installedVersionsRegex.FindAllStringSubmatch(input, -1)
	packageMetadata := &PackageMetadata{
		Name:           "node",
		PackageManager: "nvm",
	}

	for _, installedVersionMatch := range installedVersionMatches {
		arrow := installedVersionMatch[installedVersionsRegex.SubexpIndex("arrow")]
		version := installedVersionMatch[installedVersionsRegex.SubexpIndex("version")]
		packageMetadata.Installations = append(packageMetadata.Installations, InstalledPackage{
			Version: version,
		})
		if arrow != "" {
			packageMetadata.Current = version
		}
	}

	aliases := make(map[string][]string, len(strings.Split(input, "\n")))

	aliasesRegex := regexp.MustCompile(`(?m)^(?P<alias>\S*) -> (v?(?P<version>\S*))( \(-> (v?(?P<dest>\S*))\))?( (?P<default>\(default\)))?$`)
	aliasMatches := aliasesRegex.FindAllStringSubmatch(input, -1)

	for _, aliasMatch := range aliasMatches {
		dest := aliasMatch[aliasesRegex.SubexpIndex("dest")]
		if dest == "N/A" {
			// it's an alias not pointing to an installed version
			continue
		}
		alias := aliasMatch[aliasesRegex.SubexpIndex("alias")]
		version := aliasMatch[aliasesRegex.SubexpIndex("version")]

		if dest == "" {
			// we go straight from the alias to an installed version
			existingAliases := aliases[dest]
			aliasExists := false
			for _, existingAlias := range existingAliases {
				if existingAlias == alias {
					aliasExists = true
				}
			}
			if !aliasExists {
				aliases[dest] = append(aliases[dest], alias)
			}
		} else {
			// it's an alias pointing to an alias pointing to a `dest` version
			otherAlias := version
			existingAliases := aliases[dest]
			aliasExists := false
			otherAliasExists := false
			for _, existingAlias := range existingAliases {
				if existingAlias == alias {
					aliasExists = true
				}
				if existingAlias == otherAlias {
					otherAliasExists = true
				}
			}
			if !aliasExists {
				aliases[dest] = append(aliases[dest], alias)
			}
			if !otherAliasExists {
				aliases[dest] = append(aliases[dest], otherAlias)
			}
		}
	}

	for installationIndex, installation := range packageMetadata.Installations {
		packageMetadata.Installations[installationIndex].VersionAliases = aliases[installation.Version]
	}

	return []PackageMetadata{*packageMetadata}, nil
}
