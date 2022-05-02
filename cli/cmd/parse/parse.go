package parse

import (
	"bytes"
	"github.com/omniversion/omniversion/cli/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
)

var log = logrus.New()

func wrapCommand(parser func(input string) ([]models.Dependency, error)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		runParser(parser)(cmd.InOrStdin(), cmd.OutOrStdout(), cmd.ErrOrStderr(), args)
	}
}

func runParser(parser func(input string) ([]models.Dependency, error)) func(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) {
	return func(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) {
		input, err := io.ReadAll(stdin)
		if err != nil {
			log.Fatal(err)
		}
		result, parseErr := parser(string(input))
		if parseErr != nil {
			_, err = io.WriteString(stderr, parseErr.Error())
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(result) == 0 {
			// don't output empty arrays as "[]"
			// it's inconvenient for concatenation
			return
		}
		buf := new(bytes.Buffer)
		err = yaml.NewEncoder(buf).Encode(result)
		if err != nil {
			log.Fatal(err)
		}
		_, err = stdout.Write(buf.Bytes())
		if err != nil {
			log.Fatal(err)
		}
	}
}
