package nvm

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/nvm/item"
	. "github.com/omniversion/omniversion/cli/types"
	"strings"
)

func isNvmrcFile(input string) bool {
	return strings.Count(input, "\n") <= 1 && !strings.Contains(input, " ")
}

func parseNvmrcFile(input string) ([]PackageMetadata, error) {
	newItem := item.New()
	newItem.Wanted = input
	return []PackageMetadata{*newItem}, nil
}
