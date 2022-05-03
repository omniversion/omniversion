package types

type NpmAdvisory struct {
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
	}
	Recommendation string
	Id             int
	References     string
	Title          string
	Overview       string
	Url            string
}
