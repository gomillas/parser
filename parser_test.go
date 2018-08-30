package parser_test

import (
	"fmt"
	"strconv"

	"github.com/gomillas/parser"
)

func ExampleParser_Find_example1() {
	const src = "3.14159"

	m := parser.New(src)
	if token, _ := m.FindNumber(); len(token) > 0 {
		fmt.Printf("%s is a numeric expression", src)
	}

	// Output: 3.14159 is a numeric expression
}

func ExampleParser_Find_example2() {
	const src = `lorem
ipsum dolor`

	m := parser.New(src)
	for token, offset := m.FindID(); len(token) > 0; token, offset = m.FindID() {
		fmt.Printf("%s (offset: %d)\n", token, offset)
	}

	// Unordered output:
	// lorem (offset: 0)
	// ipsum (offset: 6)
	// dolor (offset: 12)
}

func ExampleParser_Find_example3() {
	const regexp = `^\s*(\w+)`
	const src = `1 2 3
	yellow 5`

	m := parser.New(src)
	for token, _ := m.Find(regexp); len(token) > 0; token, _ = m.Find(regexp) {
		if _, err := strconv.Atoi(token); err != nil {
			fmt.Printf(
				"%s is not a number (line: %d, col: %d)\n",
				token,
				m.Line(),
				m.Column(),
			)
		}
	}

	// Output: yellow is not a number (line: 2, col: 8)
}
