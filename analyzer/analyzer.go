package analyzer

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Run an analysis of Solidity contracts in `path`.
// Argument `issues` encodes the Issues to search for.
func Run(issues []Issue, paths []string) (*Report, error) {
	
	report := &Report{
		Issues:           issues,
		FilesAnalyzed:    []string{},
		FindingsPerIssue: make(map[string][]Finding),
	}

	for _, path := range paths {
		err := run(report, path)
		if err != nil {
			return &Report{}, nil
		}
	}

	return report, nil
}

func run(report *Report, path string) error {
	
	pathInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if pathInfo.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}

		for _, file := range files {
			err = run(report, filepath.Join(path, file.Name()))
			if err != nil {
				return err
			}
		}
	} else {
		file := path

		// Only analyze Solidity files
		if !strings.HasSuffix(file, ".sol") {
			return nil
		}

		findingsPerIssue, err := analyzeFile(report.Issues, file)
		if err != nil {
			return err
		}

		// Add file and findings to report.
		report.FilesAnalyzed = append(report.FilesAnalyzed, file)
		for _, issue := range report.Issues {
			report.FindingsPerIssue[issue.Identifier] = append(report.FindingsPerIssue[issue.Identifier],
				findingsPerIssue[issue.Identifier]...,
			)
		}
	}

	return nil
}
// update to "RepoLink/blob/main/"
// e.g. "https://github.com/PartyDAO/party-contracts-c4/blob/main/" for the repo at https://github.com/PartyDAO/party-contracts-c4
const baseURL = "/blob/main/"



func analyzeFile(issues []Issue, file string) (map[string][]Finding, error) {
	findings := make(map[string][]Finding)

	readFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	lineNumber := 0
	inMultilineComment := false
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		// Check for multi-line comments
		if strings.Contains(line, "/*") {
			inMultilineComment = true
		}
		if strings.Contains(line, "*/") {
			inMultilineComment = false
			continue
		}
		
		// Ignore lines that are part of a single-line or multi-line comment
		if strings.HasPrefix(strings.TrimSpace(line), "//") || inMultilineComment {
			continue
		}

		for _, issue := range issues {
			matched, _ := regexp.MatchString(issue.Pattern, line)
			if matched {
				// add base directory to ""
				// e.g for https://github.com/PartyDAO/party-contracts-c4 it will be "party-contracts-c4"
				baseDir := ""
				//fmt.Println("Absolute File Path:", file)
				relativeFilePath, err := filepath.Rel(baseDir, file)
				if err != nil {
					return nil, err
				}
				//fmt.Println("relativeFilePath:", relativeFilePath)
				CodeLocationURL := baseURL + relativeFilePath + "#L" + strconv.Itoa(lineNumber)
				//fmt.Println("CodeLocationURL:", CodeLocationURL)
				findings[issue.Identifier] = append(findings[issue.Identifier], Finding{
					IssueIdentifier: issue.Identifier,
					File:            file,
					LineNumber:      lineNumber,
					LineContent:     strings.TrimSpace(line),
                    CodeLocationURL: CodeLocationURL,
				})
			
			}
		}
	}

	return findings, nil
}
