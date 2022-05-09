package stderr

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"regexp"
	"strings"
)

type Warning struct {
	Identifier string
	Text       string
}

type Error struct {
	Location   string
	Name       string
	Problem    string
	RequiredBy string
	Text       string
	Version    string
}

type Json map[string]interface{}

type Line struct {
	Text string
}

type Output struct {
	Errors               []Error
	Json                 Json
	Lines                []Line
	Warnings             []Warning
	AuditLockfileMissing bool
}

func Strip(input string) (string, Output, error) {
	var result Output
	var errors *multierror.Error
	problemFinderRegex := regexp.MustCompile(`(?m)^npm ERR! ?(?P<text>.*)$`)
	foundProblems := problemFinderRegex.FindAllStringSubmatch(input, -1)
	for _, foundProblem := range foundProblems {
		text := foundProblem[problemFinderRegex.SubexpIndex("text")]
		// some problems have more structure
		problemParserRegex := regexp.MustCompile(`(?m)(?P<problem>[^:]*): (?P<name>\S*)@(?P<version>[^\s,]*)(, required by (?P<requiredBy>\S*))?( (?P<location>.*))?`)
		parsedProblem := problemParserRegex.FindStringSubmatch(text)
		if parsedProblem == nil {
			if strings.Contains(text, "audit This command requires an existing lockfile") {
				result.AuditLockfileMissing = true
			}
			result.Lines = append(result.Lines, Line{Text: text})
		} else {
			problem := parsedProblem[problemParserRegex.SubexpIndex("problem")]
			if problem != "missing" && problem != "extraneous" {
				errors = multierror.Append(errors, fmt.Errorf("unknown npm error kind: %q", problem))
				result.Lines = append(result.Lines, Line{Text: text})
				continue
			}
			newItem := Error{Text: text, Problem: problem}
			for groupIndex, groupName := range problemParserRegex.SubexpNames() {
				if groupIndex != 0 && groupName != "" {
					value := parsedProblem[groupIndex]
					if len(value) > 0 {
						switch groupName {
						case "name":
							newItem.Name = value
						case "version":
							newItem.Version = value
						case "requiredBy":
							newItem.RequiredBy = value
						case "location":
							newItem.Location = value
						}
					}
				}
			}
			result.Errors = append(result.Errors, newItem)
		}
	}

	warningRegex := regexp.MustCompile(`(?m)npm WARN (?P<identifier>\S*) (?P<text>.*)`)
	foundWarnings := warningRegex.FindAllStringSubmatch(input, -1)
	for _, foundWarning := range foundWarnings {
		newWarning := Warning{}
		for groupIndex, groupName := range warningRegex.SubexpNames() {
			if groupIndex != 0 && groupName != "" {
				value := foundWarning[groupIndex]
				if len(value) > 0 {
					switch groupName {
					case "identifier":
						newWarning.Identifier = value
					case "text":
						newWarning.Text = value
					}
				}
			}
		}
		result.Warnings = append(result.Warnings, newWarning)
	}

	strippedInput := problemFinderRegex.ReplaceAllLiteralString(input, "")
	strippedInput = warningRegex.ReplaceAllLiteralString(strippedInput, "")

	// in `json` format, npm will sometimes print a JSON object into stderr
	// we need to remove this before trying to parse the JSON in stdout
	errorJsonRegex := regexp.MustCompile(`(?m)^{\n {2}"error": {\n(.|\n)*`)
	foundErrorJson := errorJsonRegex.FindStringSubmatch(strippedInput)
	if foundErrorJson != nil {
		_ = json.Unmarshal([]byte(strippedInput), &result.Json)
		strippedInput = errorJsonRegex.ReplaceAllLiteralString(strippedInput, "")
	}

	strippedInput = strings.Trim(strippedInput, "\n")

	return strippedInput, result, errors.ErrorOrNil()
}
