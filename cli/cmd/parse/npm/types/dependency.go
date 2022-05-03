package types

type NpmDependency struct {
	Version      string
	From         string
	Resolved     string
	Dependencies map[string]NpmDependency

	Required    NpmRequirement
	PeerMissing bool
}
