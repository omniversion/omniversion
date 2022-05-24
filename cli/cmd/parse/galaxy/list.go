package galaxy

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/galaxy/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

// ParseListOutput parses the output of `ansible-galaxy list`, `ansible-galaxy role list`,
// `ansible-galaxy list -vvv` or `ansible-galaxy --version`
func ParseListOutput(input string) ([]PackageMetadata, error) {

	result := make([]PackageMetadata, 0, strings.Count(input, "\n")+5)

	ansibleCoreMatch := regexp.MustCompile(`(?m)^ansible-galaxy \[core (?P<version>\S+)]$`).FindStringSubmatch(input)
	if ansibleCoreMatch != nil {
		newItem := item.New("ansible-core")
		version := ansibleCoreMatch[1]
		newItem.Current = shared.CleanVersion(version)
		newItem.Aliases = []string{"ansible-galaxy"}
		newItem.Installations = []InstalledPackage{{
			Version: shared.CleanVersion(version),
		}}
		result = append(result, *newItem)
	}

	pythonMatch := regexp.MustCompile(`(?m)^\s*python version = (?P<version>\S+)`).FindStringSubmatch(input)
	if pythonMatch != nil {
		newItem := item.New("python")
		version := pythonMatch[1]
		newItem.Current = shared.CleanVersion(version)
		newItem.Installations = []InstalledPackage{{
			Version: shared.CleanVersion(version),
		}}
		result = append(result, *newItem)
	}

	clangMatch := regexp.MustCompile(`(?m)^\s*python version = \S+ \(.*\) \[Clang (?P<version>\S+) .*]`).FindStringSubmatch(input)
	if clangMatch != nil {
		newItem := item.New("clang")
		version := clangMatch[1]
		newItem.Current = shared.CleanVersion(version)
		newItem.Installations = []InstalledPackage{{
			Version: shared.CleanVersion(version),
		}}
		result = append(result, *newItem)
	}

	jinjaMatch := regexp.MustCompile(`(?m)^\s*jinja version = (?P<version>\S+)$`).FindStringSubmatch(input)
	if jinjaMatch != nil {
		newItem := item.New("jinja")
		version := jinjaMatch[1]
		newItem.Current = shared.CleanVersion(version)
		newItem.Installations = []InstalledPackage{{
			Version: shared.CleanVersion(version),
		}}
		result = append(result, *newItem)
	}

	compiledRegex := regexp.MustCompile(`(?m)^- (?P<name>\S*), v?(?P<version>\S+)$`)
	matches := compiledRegex.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		name := match[compiledRegex.SubexpIndex("name")]
		version := match[compiledRegex.SubexpIndex("version")]
		newItem := item.New(name)
		nameComponents := strings.Split(name, ".")
		shortName := nameComponents[len(nameComponents)-1]
		newItem.Aliases = []string{shortName}
		newItem.Current = shared.CleanVersion(version)
		newItem.Installations = []InstalledPackage{{
			Version: shared.CleanVersion(version),
		}}
		result = append(result, *newItem)
	}
	return result, nil
}
