package shared

import "strings"

func ShortModuleName(fullName string) string {
	components := strings.Split(fullName, "/")
	switch len(components) {
	case 2:
		return components[1]
	case 1:
		return components[0]
	default:
		return strings.Join(components[2:], "/")
	}
}

func CleanVersion(version string) string {
	return strings.Trim(strings.TrimLeft(version, "v"), "\n")
}

func CleanVersions(versions []string) []string {
	var cleanedVersions []string
	for _, version := range versions {
		cleanedVersion := CleanVersion(version)
		if cleanedVersion != "" {
			cleanedVersions = append(cleanedVersions, cleanedVersion)
		}
	}
	return cleanedVersions
}

func LastVersion(versions []string) string {
	if len(versions) == 0 {
		return ""
	}
	return CleanVersion(versions[len(versions)-1])
}
