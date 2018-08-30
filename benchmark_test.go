package parser_test

import (
	"testing"

	"github.com/gomillas/parser"
)

func Benchmark(b *testing.B) {
	const src = "125 + 2 * (sqrt 9 - 1) - 3"

	for i := 0; i < b.N; i++ {
		m := parser.New(src)
		if _, err := mathExp(m, "="); err != nil {
			b.FailNow()
		}
	}
}
