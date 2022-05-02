package npm

type NpmPeerMissingRequirement struct {
	RequiredBy string
	Requires   string
}
