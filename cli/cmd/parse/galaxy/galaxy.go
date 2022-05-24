package galaxy

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func ParseOutput(input string) ([]PackageMetadata, error) {
	var testYaml []map[string]string
	err := yaml.Unmarshal([]byte(input), &testYaml)
	if err == nil {
		return ParseRequirementsYamlFile(input)
	} else {
		return ParseListOutput(input)
	}
}

var ParseCommand = &cobra.Command{
	Use:     "galaxy",
	Short:   "Parse the output of `ansible-galaxy list` or a `requirements.yaml` file",
	Long:    "Translate `ansible-galaxy` output into the omniversion format.",
	Run:     shared.WrapCommand(ParseOutput),
	Aliases: []string{"ansible", "ansible-galaxy"},
}
