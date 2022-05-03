package types

type Advisory struct {
	Access             string  `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	CVSSScore          float64 `json:",omitempty" toml:",omitempty" yaml:"cvss_score,omitempty"`
	Id                 int     `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Overview           string  `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	PatchedVersions    string  `json:",omitempty" toml:",omitempty" yaml:"patched_versions,omitempty"`
	Recommendation     string  `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	References         string  `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Severity           string  `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Title              string  `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Url                string  `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	VulnerableVersions string  `json:",omitempty" toml:",omitempty" yaml:"vulnerable_versions,omitempty"`
}
