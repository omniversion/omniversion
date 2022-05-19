package types

// PackageMetadata contains all available information about a particular software package,
// including existing installations, security advisories and all sorts of other metadata.
// If multiple versions of the same package need to be tracked
// for a single combination of package manager and host,
// this should be done using a single PackageMetadata structure.
// By contrast, otherwise identical packages made available by different package managers
// are considered different, as their metadata and/or content *may* indeed differ.
// We also distinguish between package metadata on different hosts,
// as installations are different and available versions
// will depend on the configured repositories, firewall settings etc.
type PackageMetadata struct {
	// Name is the identifier used to install the package.
	Name string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Aliases contain alternative names by which the package may be known or have been known
	Aliases []string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`

	// PackageManager is the package manager through which the package can be installed.
	// If a package is available through multiple package managers, we should use separate
	// PackageMetadata objects to track them.
	PackageManager string `json:"packageManager,omitempty" toml:"package_manager,omitempty" yaml:"package_manager,omitempty"`
	// Host is the name of the machine on which this package has been installed, requested
	// or otherwise tracked by a package manager.
	// This may be `localhost`, a hostname (with or without schema) or the name of a docker container.
	// While the naming is flexible, care should be taken to keep these identifiers unique.
	// If a package is installed on multiple hosts, we should use separate
	// PackageMetadata objects to track them.
	Host string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Installations contains metadata on versions of the package installed on the current machine.
	Installations []InstalledPackage `json:",omitempty" toml:",omitempty" yaml:",omitempty"`

	// DependencyType is the kind of dependency (`prod`, `dev`, `peer`)
	Type DependencyType `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Optional is true if the dependency need not be installed
	Optional *bool `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Direct specifies whether the dependency has been required directly or transitively
	Direct *bool `json:",omitempty" toml:",omitempty" yaml:",omitempty"`

	// Current is the currently installed version of the package as reported by the package manager.
	// E.g. what is reported by `npm ls <package>`, `rvm info` `pip show <package>`.
	Current string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Default is the version the package manager would select by default.
	// E.g. the version designated as `default` by `rvm ls`.
	Default string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Latest is the most recent version of the package known to be available.
	Latest string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Wanted is the version of the package that would be installed
	// based on the relevant constraints defined for this package manager.
	// E.g. this is the highest version matching the range defined in `package.json` for `npm`.
	Wanted string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Locked is the version specified in the relevant lock file for this package manager.
	// E.g. the version defined in `package-lock.json` or `npm-shrinkwrap.json`.
	Locked string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Extraneous is true if the package is not required, but installed.
	Extraneous *bool `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Missing is true if the package is required, but not installed.
	// Installed optional packages do not count as extraneous.
	Missing *bool `json:",omitempty" toml:",omitempty" yaml:",omitempty"`

	// Dependencies are packages directly required by this package at runtime.
	Dependencies []string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// DevDependencies are packages directly required by this package during development.
	DevDependencies []string `json:"devDependencies,omitempty" toml:"dev_dependencies,omitempty" yaml:"dev_dependencies,omitempty"`
	// PeerDependencies are packages directly required, but not managed by this package.
	// E.g. a plugin might add functionality to another package
	// which is assumed to be installed independently of the plugin.
	PeerDependencies []string `json:"peerDependencies,omitempty" toml:"peer_dependencies,omitempty" yaml:"peer_dependencies,omitempty"`

	// Architecture is the architecture for which the package was compiled.
	// E.g. the architecture field reported by `rvm ls` or `apt list`.
	Architecture string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Author is the author or list of authors reported by the package manager.
	Author string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Description is the package description reported by the package manager.
	Description string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Homepage is the package's homepage reported by the package manager.
	Homepage string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// License is a license identifier as reported by the package manager.
	// This is not currently standardized across package managers.
	License string `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Sources through which the package is available, as reported by the package manager.
	// E.g. the `sources` field in `apt list`.
	Sources []PackagesSource `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
	// Advisories contains security notices on known vulnerabilities.
	// E.g. the information contained in the output of `npm audit`.
	Advisories []Advisory `json:",omitempty" toml:",omitempty" yaml:",omitempty"`
}

type DependencyType string

const (
	ProdDependency DependencyType = "prod"
	DevDependency  DependencyType = "dev"
	PeerDependency DependencyType = "peer"
)
