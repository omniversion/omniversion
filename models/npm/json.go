package npm

type NpmJson struct {
	Advisories   map[string]NpmAdvisory
	Problems     []string
	Dependencies map[string]NpmDependency
}
