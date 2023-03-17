package analyzer

import (
	"fmt"
	"strings"
	"time"
)

// Report is the end result of an analysis containing the files analyzed,
// the issues searched for and a map of findings per issue.
type Report struct {
	Issues        []Issue
	FilesAnalyzed []string
	// Key is Issue Identifier
	FindingsPerIssue map[string][]Finding
	NumLines         int
	NumFunctions     int
	DateTime         time.Time
	TotalFindings     int
	TotalGasFindings  int
	TotalLowFindings  int
}

// Issue represents an Issue to search for in the codebase.
// The pattern field is a RegEx string which must compile.
type Issue struct {
	Identifier string
	Severity   Severity
	Title      string
	Link       string
	Pattern    string
}

// Finding represents a possible Issue found in the codebase.
type Finding struct {
	IssueIdentifier string
	File            string
	LineNumber      int
	LineContent     string
	CodeLocationURL string
}

// Severity type defining the severity level for an Issue.
type Severity int

// The Severity Enum.
const (
	GASOP Severity = iota
	NC
	LOW
)

// Markdown returns the report as string in markdown style.
func (r Report) Markdown() string {
	const c4uditRepoLink = "https://github.com/byterocket/c4udit"
	// Issue output in Code4Rena format:
	// ### {{ issue.Title }}
	//
	// #### Impact
	// Issue information: [{{ issue.Identifier }}]({{ issue.Link }})
	//
	// #### Findings
	// {{ _, finding := range findings: finding.String() }}
	//
	// #### Tools used
	// [c4udit]({{ c4uditRepoLink }})
	//

		// Add date/time stamp
	date := time.Now().Format("January 2, 2006")
	dateString := fmt.Sprintf("**Audit Date:** %s\n\n", date)

	buf := strings.Builder{}
	buf.WriteString(dateString) // Add date/time stamp

	// Add findings summary
	totalFindings := 0
	totalGasFindings := 0
	totalLowFindings := 0
	for _, issue := range r.Issues {
		findings := r.FindingsPerIssue[issue.Identifier]
		count := len(findings)
		totalFindings += count
		if issue.Severity == GASOP {
			totalGasFindings += count
		} else if issue.Severity == LOW {
			totalLowFindings += count
		}
	}
	r.TotalFindings = totalFindings
	r.TotalGasFindings = totalGasFindings
	r.TotalLowFindings = totalLowFindings

	// Add summary line
	buf.WriteString(fmt.Sprintf("**Total Findings:** %d, **Total Gas Findings:** %d, **Total Low Findings:** %d\n\n", totalFindings, totalGasFindings, totalLowFindings))

	
	buf.WriteString("---\n\n")
	
	buf.WriteString("# Audit Report\n")
	buf.WriteString("## Notice: This report may contain some false positives, check all findings manually to verify them\n")
	buf.WriteString("\n")

	buf.WriteString("## Files analyzed\n")
	for _, f := range r.FilesAnalyzed {
		buf.WriteString("- " + f + "\n")
	}
	buf.WriteString("\n")

	buf.WriteString("## **Findings Summary**\n")
	buf.WriteString("\n")
	buf.WriteString("| Issue Identifier | Issue | Count |\n")
	buf.WriteString("| ---------------- | ----- | ----- |\n")
	for _, issue := range r.Issues {
		findings := r.FindingsPerIssue[issue.Identifier]
		count := len(findings)
		if count == 0 {
			continue
		}
		buf.WriteString(fmt.Sprintf("| %s | %s | %d |\n", issue.Identifier, issue.Title, count))
	}
	buf.WriteString("\n")

	// CHANGE: Add a horizontal line to visually separate the findings section
	buf.WriteString("---\n\n")

	buf.WriteString("## Findings\n")
	buf.WriteString("\n")
	for _, issue := range r.Issues {
		findings := r.FindingsPerIssue[issue.Identifier]
		if len(findings) == 0 {
			continue
		}
		numFindings := len(findings)
		buf.WriteString(strings.Repeat("_", len(issue.Identifier)+len(issue.Title)+7) + "\n") // add a line of underscores
		buf.WriteString("### **[" + issue.Identifier + "] " + issue.Title + "**\n")
		buf.WriteString(strings.Repeat("_", len(issue.Identifier)+len(issue.Title)+7) + "\n") // add a line of underscores
		buf.WriteString("\n")

		// Impact
		buf.WriteString("#### **Impact**\n")
		buf.WriteString(issue.Link + "\n")
		buf.WriteString("\n")

		// Number of Findings
		buf.WriteString(fmt.Sprintf("#### **Number of Findings:** %d\n", numFindings))
		buf.WriteString("\n")

    	// Findings
    	buf.WriteString("#### **Details:**\n")
		buf.WriteString("| File | Line Number | Line Content | Code Location |\n")
		buf.WriteString("| ---- | ----------- | ------------ | ------------- |\n")
    	for _, finding := range findings {
        	buf.WriteString(finding.TableRow())
    	}
    	buf.WriteString("\n")

		// Tools used
		buf.WriteString("#### **Tools used**\n")
		buf.WriteString("[c4udit](" + c4uditRepoLink + ")\n")

		buf.WriteString("\n")
	}

	return buf.String()
}

func (r Report) String() string {
	// Build files string.
	files := "**Files analyzed:**\n"
	for _, f := range r.FilesAnalyzed {
		files += fmt.Sprintf("- %s\n", f)
	}
	files += "\n"

	// Build issues string.
	issues := "**Issues found:**\n"
	for i, issue := range r.Issues {
		// Get findings for issue
		findings := r.FindingsPerIssue[issue.Identifier]

		// Skip if no findings
		if len(findings) == 0 {
			continue
		}

		// Add findings per issue
		issues += " " + issue.Identifier + ":\n"
		for _, finding := range findings {
			issues += "  " + finding.String()
		}

		// Add newline if not last issue
		if i+1 != len(r.Issues) {
			issues += "\n"
		}
	}

	return files + issues
}

func (i Issue) String() string {
	return i.Identifier
}

func (f Finding) String() string {
	return fmt.Sprintf("%s::%d => %s\n", f.File, f.LineNumber, f.LineContent)
}

func (s Severity) String() string {
	return []string{
		"Gas Optimization",
		"Non-Critical",
		"Low Risk",
	}[s]
}

func (f Finding) TableRow() string {
    return fmt.Sprintf("| %s | %d | %s | [%s](%s) |\n", f.File, f.LineNumber, f.LineContent, "View Code", f.CodeLocationURL)
}
