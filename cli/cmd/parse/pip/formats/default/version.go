package _default

import (
	"fmt"
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/formats"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
)

func ParseVersionOutput(input string) ([]PackageMetadata, error) {
	compiledRegex := regexp.MustCompile(formats.VersionOutputRegex)
	match := compiledRegex.FindStringSubmatch(input)
	if match == nil {
		return []PackageMetadata{}, fmt.Errorf("invalid version data")
	}
	version := match[compiledRegex.SubexpIndex("version")]
	location := match[compiledRegex.SubexpIndex("location")]
	pythonVersion := match[compiledRegex.SubexpIndex("pythonVersion")]
	return []PackageMetadata{{
		Name:           "pip",
		Current:        version,
		PackageManager: "pip",
		Installations: []InstalledPackage{{
			Version:  version,
			Location: location,
		}},
	}, {
		Name:           "python",
		Current:        pythonVersion,
		PackageManager: "pip",
		Installations: []InstalledPackage{{
			Version: pythonVersion,
		}},
	}}, nil
}
