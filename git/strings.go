package git

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// IssueMatch parses the issue ID from a Git log line.
var IssueMatch *regexp.Regexp

// LogParts parses all parts of a git --oneline log
var LogParts *regexp.Regexp

var ErrLogParse = errors.New("error slice parsing log")
var ErrMatching = errors.New("error matching string")

func init() {
	IssueMatch = regexp.MustCompile("(#([0-9]+))")
	LogParts = regexp.MustCompile(`([a-z0-9]+)(.+)\(#([0-9]+)\)`)
}

// Log define api..
//
//
type Log struct {
	Matcher *regexp.Regexp
	hash    string
	issue   int
	message string
}

// NewLogParser returns a Log configured for processing
// git log --oneline's output.
func NewLogParser() *Log {
	return &Log{
		Matcher: LogParts,
	}
}

// Parse breaks appart a single git log line
func (l *Log) Parse(raw string) error {
	slice := l.Matcher.FindAllStringSubmatch(raw, -1)
	if len(slice) < 1 {
		return ErrMatching
	}
	parsed := slice[0]
	if len(parsed) < 4 {
		return ErrLogParse
	}

	l.hash = parsed[1]
	l.message = strings.TrimSpace(parsed[2])
	iss, err := strconv.Atoi(parsed[3])
	if err != nil {
		return err
	}
	l.issue = iss
	return nil
}

// Hash returns git commit Hash
func (l Log) Hash() string {
	return l.hash
}

// Message returns the git one-line commit
func (l Log) Message() string {
	return l.message
}

// Issue returns the parsed Git issue appended on the end of the git one-line
func (l Log) Issue() int {
	return l.issue
}

// ProjectMarkdown returns an updated git log --oneline
func (l Log) ProjectMarkdown(u url.URL) string {
	basePath := u.Path
	commitURL := u
	issuePath := fmt.Sprintf("%s/issues/%d", basePath, l.Issue())
	u.Path = issuePath
	commitURL.Path = fmt.Sprintf("%s/commit/%s", basePath, l.Hash())
	return fmt.Sprintf("[%s](%s) %s [%d](%s)", l.hash, commitURL.String(), l.message, l.issue, u.String())
}
