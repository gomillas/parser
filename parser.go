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
func (p *Parser) Find(expr string) (result string, offset int) {
	re := regexp.MustCompile(expr)

	substr := p.Source[p.Offset:]
	loc := re.FindStringSubmatchIndex(substr)
	if loc != nil {
		i := 0
		if len(loc) > 2 {
			i = 2
		}

		result = substr[loc[i]:loc[i+1]]
		offset = p.Offset + loc[i]
		p.Offset += loc[1]
	}

	return result, offset
}

// FindNumber returns the next number.
func (p *Parser) FindNumber() (string, int) {
	return p.Find(numberRegexp)
}

// FindID returns the next identifier.
func (p *Parser) FindID() (string, int) {
	return p.Find(idRegexp)
}

// Line of the current offset.
func (p *Parser) Line() int {
	return strings.Count(p.Source[:p.Offset], "\n") + 1
}

// Column of the current offset.
func (p *Parser) Column(params ...int) int {
	substr := p.Source[:p.Offset]
	pos := strings.LastIndex(substr, "\n")

	if pos > -1 {
		substr = substr[pos+1:]
	}

	return len(substr) + 1
}
