// Package parser uses regular expressions to split a string into smaller 'tokens'.
//
// The advantage of using this package is that you can parse any arbitrary
// string, as long as you can split it into smaller regular expressions.
package parser

import (
	"regexp"
	"strings"
)

// Common regular expressions.
const (
	numberRegexp             = `^\s*([+-]?(0|[1-9][0-9]*)(\.[0-9]+)?([eE][+-]?[0-9]+)?)`
	idRegexp                 = `^\s*([_a-z][_a-z0-9]*)`
	singleQuotedStringRegexp = `^\s*`
)

// Parser contains information about the current state.
type Parser struct {
	Source string
	Offset int
}

// New instance.
func New(src string) *Parser {
	return &Parser{Source: src, Offset: 0}
}

// Find returns the offset and the string matching a regular expression,
// starting from the current Offset position.
func (m *Parser) Find(expr string) (result string, offset int) {
	re := regexp.MustCompile(expr)

	substr := m.Source[m.Offset:]
	loc := re.FindStringSubmatchIndex(substr)
	if loc != nil {
		i := 0
		if len(loc) > 2 {
			i = 2
		}

		result = substr[loc[i]:loc[i+1]]
		offset = m.Offset + loc[i]
		m.Offset += loc[1]
	}

	return result, offset
}

// FindNumber returns the next number.
func (m *Parser) FindNumber() (string, int) {
	return m.Find(numberRegexp)
}

// FindID returns the next identifier.
func (m *Parser) FindID() (string, int) {
	return m.Find(idRegexp)
}

// Line of the current offset.
func (m *Parser) Line() int {
	return strings.Count(m.Source[:m.Offset], "\n") + 1
}

// Column of the current offset.
func (m *Parser) Column(params ...int) int {
	substr := m.Source[:m.Offset]
	pos := strings.LastIndex(substr, "\n")

	if pos > -1 {
		substr = substr[pos+1:]
	}

	return len(substr) + 1
}
