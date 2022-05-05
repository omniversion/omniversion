package types

// Advisory is a security notice reporting a known vulnerability.
// E.g. on item of the output of `npm audit`.
type Advisory struct {
	// CVSSScore is the Common Vulnerability Scoring System assigned to this vulnerability.
	// It should be a number between 0 and 10.
	CVSSScore float64 `json:",omitempty" toml:",omitempty" yaml:"cvss_score,omitempty"`
	// Identifier is a string used to identify this vulnerability.
	// E.g. the dictionary key used by `npm audit --json`.
	// Not currently standardized across package managers.
	Identifier string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Overview is a human-readable high-level description of this vulnerability.
	Overview string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// PatchedVersions is a version range or list of ranges for which this vulnerability has been patched.
	PatchedVersions string `json:",omitempty" toml:",omitempty" yaml:"patched_versions,omitempty"`
	// Recommendation is human-readable advice on how to deal with this vulnerability.
	Recommendation string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Severity is the level of impact the vulnerability may have.
	// E.g. `critical`, `medium`, `high`.
	// This value is not currently standardized across package managers.
	Severity string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Title is the title of this vulnerability as reported by the package manager.
	Title string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Url is a URL containing further information.
	Url string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// VulnerableVersions is a version range or list of ranges affected by this vulnerability.
	VulnerableVersions string `json:",omitempty" toml:",omitempty" yaml:"vulnerable_versions,omitempty"`
}
