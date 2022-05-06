package types

type NpmPackageJson struct {
	Name             string            `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Version          string            `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	License          string            `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Dependencies     map[string]string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	DevDependencies  map[string]string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	PeerDependencies map[string]string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
}
