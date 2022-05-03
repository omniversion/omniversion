package types

type NpmOutdatedDependency struct {
	Current   string
	Wanted    string
	Latest    string
	Dependent string
	Location  string
}
