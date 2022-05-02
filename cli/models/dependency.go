package models

type Dependency struct {
	Name       string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Identifier string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Host       string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`

	Version   string                `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Default   string                `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Latest    string                `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Wanted    string                `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Installed []InstalledDependency `json:",omitempty" toml:",omitempty" yaml:",omitempty"`

	Architecture string     `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Author       string     `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Description  string     `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Homepage     string     `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	License      string     `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Pm           string     `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Sources      []string   `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Advisories   []Advisory `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
}
