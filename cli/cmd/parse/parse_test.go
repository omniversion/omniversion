package parse

import (
	"bytes"
	"github.com/omniversion/omniversion-cli/cli/models"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyOutput(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	runParser(func(input string) ([]models.Dependency, error) {
		return []models.Dependency{}, nil
	})(stdin, stdout, stderr, []string{})

	assert.Equal(t, "", string(stdout.Bytes()))
	assert.Equal(t, "", string(stderr.Bytes()))
}

func TestSimpleOutput(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	runParser(func(input string) ([]models.Dependency, error) {
		return []models.Dependency{
			{Name: "test"},
		}, nil
	})(stdin, stdout, stderr, []string{})

	assert.Equal(t, "- name: test\n", string(stdout.Bytes()))
	assert.Equal(t, "", string(stderr.Bytes()))
}

func TestWrapCommand(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	command := &cobra.Command{Use: "a", Args: cobra.NoArgs, Run: func(cmd *cobra.Command, args []string) {}}
	command.SetIn(stdin)
	command.SetOut(stdout)
	command.SetErr(stderr)
	wrapCommand(func(input string) ([]models.Dependency, error) {
		return []models.Dependency{
			{Name: "test"},
		}, nil
	})(command, []string{})

	assert.Equal(t, "- name: test\n", stdout.String())
}
