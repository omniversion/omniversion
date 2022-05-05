package shared

import (
	"bytes"
	"fmt"
	"github.com/omniversion/omniversion/cli/test"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleOutput(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	err := runParser(func(input string) ([]PackageMetadata, error) {
		return []PackageMetadata{
			{Name: "test"},
		}, nil
	})(stdin, stdout)

	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "- name: test\n")
}

func TestWrapCommand(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	command := &cobra.Command{Use: "a", Args: cobra.NoArgs, Run: func(cmd *cobra.Command, args []string) {}}
	command.SetIn(stdin)
	command.SetOut(stdout)
	WrapCommand(func(input string) ([]PackageMetadata, error) {
		return []PackageMetadata{
			{Name: "test"},
		}, nil
	})(command, []string{})

	assert.Contains(t, stdout.String(), "- name: test\n")
}

func TestWrapCommandWithError(t *testing.T) {
	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	command := &cobra.Command{Use: "a", Args: cobra.NoArgs, Run: func(cmd *cobra.Command, args []string) {}}
	command.SetIn(stdin)
	command.SetOut(stdout)
	WrapCommand(func(input string) ([]PackageMetadata, error) {
		return []PackageMetadata{}, fmt.Errorf("test error")
	})(command, []string{})

	assert.True(t, fatal)
}

func TestRunParserWithInvalidStdIn(t *testing.T) {
	stdin := test.ErrorReader{}
	stdout := new(bytes.Buffer)
	err := runParser(func(input string) ([]PackageMetadata, error) { return []PackageMetadata{}, nil })(stdin, stdout)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "test error")
}

func TestRunParserWithInvalidStdOut(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := test.ErrorWriter{}
	err := runParser(func(input string) ([]PackageMetadata, error) { return []PackageMetadata{}, nil })(stdin, stdout)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "test error")
}

func TestRunParserWithError(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	err := runParser(func(input string) ([]PackageMetadata, error) { return []PackageMetadata{}, fmt.Errorf("test error") })(stdin, stdout)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "test error")
}

func TestInvalidOutputFormat(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	previousOutputFormat := OutputFormat
	OutputFormat = "invalid"
	err := runParser(func(input string) ([]PackageMetadata, error) { return []PackageMetadata{}, nil })(stdin, stdout)
	OutputFormat = previousOutputFormat

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unknown output format")
}

func TestInvalidJsonOutput(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	previousOutputFormat := OutputFormat
	OutputFormat = "json"
	err := runParser(func(input string) ([]PackageMetadata, error) {
		return []PackageMetadata{
			{Name: "test"},
		}, nil
	})(stdin, stdout)
	OutputFormat = previousOutputFormat

	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "test")
}

func TestInvalidTomlOutput(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	previousOutputFormat := OutputFormat
	OutputFormat = "toml"
	err := runParser(func(input string) ([]PackageMetadata, error) {
		return []PackageMetadata{
			{Name: "test"},
		}, nil
	})(stdin, stdout)
	OutputFormat = previousOutputFormat

	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "test")
}
