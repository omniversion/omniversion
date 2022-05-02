package npm

type NpmRequirement struct {
	Id          string `json:"_id"`
	Name        string
	Version     string
	PeerMissing []NpmPeerMissingRequirement
}
