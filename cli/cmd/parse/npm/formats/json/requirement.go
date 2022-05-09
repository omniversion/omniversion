package json

type Requirement struct {
	Id          string `json:"_id"`
	Name        string
	Version     string
	PeerMissing []struct {
		RequiredBy string
		Requires   string
	}
}
