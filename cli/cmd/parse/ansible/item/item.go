package item

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
)

func New(name string) *PackageMetadata {
	newItem := &PackageMetadata{
		Name: name,
	}
	if shared.InjectPackageManager {
		newItem.PackageManager = "ansible"
	}
	return newItem
}
