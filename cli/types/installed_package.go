package types

// InstalledPackage is an installation of a software package
// in a particular location on a particular host.
type InstalledPackage struct {
	// Location is the absolute path to the installation.
	Location string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Version is the package version that has been installed.
	Version string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// VersionAliases are alternative names for Version (e.g. `latest`, `16` etc.)
	VersionAliases []string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	Hash           string   `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
}
