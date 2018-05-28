package git

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// IssueMatch parses the issue ID from a Git log line.
var IssueMatch *regexp.Regexp

// LogParts parses all parts of a git --oneline log
var LogParts *regexp.Regexp

var ErrLogParse = errors.New("error slice parsing log")

func init() {
	IssueMatch = regexp.MustCompile("(#([0-9]+))")
	LogParts = regexp.MustCompile(`([a-z0-9]+)(.+)\(#([0-9]+)\)`)
}

// Log define api..
//
//
type Log struct {
	Matcher regexp.Regexp
	hash    string
	issue   int
	message string
}

// Parse breaks appart a single git log line
func (l *Log) Parse(raw string) error {
	slice := l.Matcher.FindAllStringSubmatch(raw, -1)
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

func (l Log) Hash() string {
	return l.hash
}

func (l Log) Messsage() string {
	return l.message
}

func (l Log) Issue() int {
	return l.issue
}
