package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/omniversion/omniversion/cli/cmd/version"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io"
)

var log = logrus.New()

// InjectPackageManager causes the package manager's name be injected into each package metadata item.
// This is usually not needed, as it increases the size of output unnecessarily and
// items for different package managers are stored in separate files.
var InjectPackageManager = false

// OutputFormat determines how the command's output is formatted
var OutputFormat = "yaml"

func WrapCommand(parser func(input string) ([]PackageMetadata, error)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := runParser(parser)(cmd.InOrStdin(), cmd.OutOrStdout())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func runParser(parser func(input string) ([]PackageMetadata, error)) func(stdin io.Reader, stdout io.Writer) error {
	return func(stdin io.Reader, stdout io.Writer) error {
		input, err := io.ReadAll(stdin)
		if err != nil {
			return err
		}
		result, parseErr := parser(string(input))
		if parseErr != nil {
			return parseErr
		}
		data := map[string]interface{}{
			"version": version.Version,
			"items":   result,
		}
		buf := new(bytes.Buffer)
		switch OutputFormat {
		case "yaml":
			err = yaml.NewEncoder(buf).Encode(data)
		case "toml":
			err = toml.NewEncoder(buf).Encode(data)
		case "json":
			err = json.NewEncoder(buf).Encode(data)
		default:
			err = fmt.Errorf("unknown output format: %q. valid values are `yaml` (default), `toml` and `json`", OutputFormat)
		}
		if err != nil {
			return err
		}
		_, err = stdout.Write(buf.Bytes())
		if err != nil {
			return err
		}
		return nil
	}
}
