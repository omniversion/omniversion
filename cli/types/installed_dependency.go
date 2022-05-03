package types

type InstalledDependency struct {
	Version  string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Location string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
}
