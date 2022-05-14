package types

type PackagesSource struct {
	Identifier string   `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Url        string   `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Versions   []string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
}
