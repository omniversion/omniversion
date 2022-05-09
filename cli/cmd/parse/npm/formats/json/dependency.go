package json

type Dependency struct {
	Dependencies map[string]Dependency
	Extraneous   bool
	From         string
	Invalid      string
	PeerMissing  bool
	Problems     []string
	// either a `string` or a `Requirement` object
	Required interface{}
	Resolved string
	Version  string
}
