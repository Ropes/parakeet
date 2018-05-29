package git

import (
	"fmt"
	"net/url"
	"strconv"
	"testing"
)

var (
	logLines = []string{
		"666666e11 Guardrails to cache usage refactored (#1287)",
		"131313131 major refactoring to clean things up. (#1302)",
	}
	tests = []struct {
		raw   string
		issue int
		hash  string
	}{
		{
			raw:   "666666e11 Guardrails to cache usage refactored (#1287)",
			issue: 1287,
			hash:  "666666e11",
		},
		{
			raw:   "131313131 major refactoring to clean things up. (#1302)",
			issue: 1302,
			hash:  "131313131",
		},
	}
)

func TestParseIssue(t *testing.T) {
	for i, test := range tests {
		t.Run(fmt.Sprintf("match-%d", i), func(t *testing.T) {
			if !IssueMatch.Match([]byte(test.raw)) {
				t.Errorf("parse failure: %q", test)
			}
			if !LogParts.Match([]byte(test.raw)) {
				t.Errorf("logparts parse failure: %q", test)
			}
		})
		t.Run(fmt.Sprintf("findall-%d", i), func(t *testing.T) {
			if parts := LogParts.FindAllStringSubmatch(test.raw, -1); len(parts[0]) < 3 {
				t.Errorf("findall did not return all exp parts: %q", parts)
			} else if strslice := parts[0]; len(strslice) > 3 {
				if test.raw != strslice[0] {
					t.Error("first string should match raw returned value")
				}
				issue, err := strconv.Atoi(strslice[3])
				if err != nil || test.issue != issue {
					t.Errorf("issue mismatch %d != %d", test.issue, issue)
				}
				if test.hash != strslice[1] {
					t.Errorf("issue hash mismatch %q != %q", test.hash, strslice[1])
				}
			}
		})
	}
}

func TestLog(t *testing.T) {
	for i, test := range tests {
		log := NewLogParser()
		t.Run(fmt.Sprintf("parse-%d", i), func(t *testing.T) {
			if err := log.Parse(test.raw); err != nil {
				t.Fatalf("error parsing: %v", err)
			}
		})
		t.Run(fmt.Sprintf("data-%d", i), func(t *testing.T) {
			if test.issue != log.Issue() {
				t.Errorf("issue mismatch %d != %d", test.issue, log.Issue())
			}
			if test.hash != log.Hash() {
				t.Errorf("issue hash mismatch %q != %q", test.hash, log.Hash())
			}
			if log.Message() == "" {
				t.Error("message is empty")
			}
		})
		t.Run(fmt.Sprintf("markdown-%d", i), func(t *testing.T) {
			u, _ := url.Parse("https://github.com/neh/project")
			ret := log.ProjectMarkdown(*u)
			//TODO: assert markdown valid
			t.Log(ret)
		})
	}
}
