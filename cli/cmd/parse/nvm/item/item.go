package item

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
)

func New() *PackageMetadata {
	newItem := &PackageMetadata{
		Name: "node",
	}
	if shared.InjectPackageManager {
		newItem.PackageManager = "nvm"
	}
	return newItem
}
