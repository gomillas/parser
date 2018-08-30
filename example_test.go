package parser_test

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gomillas/parser"
)

func Example() {
	const src = "125 + 2 * (sqrt 9 - 1) - 3"

	m := parser.New(src)
	if result, err := mathExp(m, "="); err == nil {
		fmt.Println(result)
	} else {
		panic(
			fmt.Errorf("%s (line: %d, col: %d)", err.Error(), m.Line(), m.Column()),
		)
	}

	// Output: 126
}

var operators = map[string]struct {
	priority  int
	operation func(x, y float64) float64
}{
	"=": {priority: 0, operation: func(x, y float64) float64 { return y }},
	"+": {priority: 1, operation: func(x, y float64) float64 { return x + y }},
	"-": {priority: 1, operation: func(x, y float64) float64 { return x - y }},
	"*": {priority: 2, operation: func(x, y float64) float64 { return x * y }},
	"/": {priority: 2, operation: func(x, y float64) float64 { return x / y }},
}

var functions = map[string](func(x float64) float64){
	"sin": math.Sin, "cos": math.Cos, "tan": math.Tan, "sqrt": math.Sqrt,
}

func mathExp(p *parser.Parser, op0 string) (result float64, err error) {
	if result, err = valueExp(p); err != nil {
		return result, err
	}

	for op1 := operatorExp(p, op0); len(op1) > 0; op1 = operatorExp(p, op0) {
		val, err := mathExp(p, op1)
		if err != nil {
			return result, err
		}

		result = operators[op1].operation(result, val)
	}

	return
}

func valueExp(p *parser.Parser) (result float64, err error) {
	if token, _ := p.FindNumber(); len(token) > 0 {
		return strconv.ParseFloat(token, 64)
	} else if token, _ := p.Find(`^\s*\(`); len(token) > 0 {
		result, err = mathExp(p, "=")
		if err != nil {
			return
		}

		if token, _ := p.Find(`^\s*\)`); len(token) == 0 {
			err = errors.New(") is expected")
			return
		}
	} else if token, _ := p.Find(`^\s*(sin|cos|tan|sqrt)`); len(token) > 0 {
		result, err = valueExp(p)
		if err != nil {
			return
		}

		result = functions[token](result)
	} else {
		err = errors.New("Value is expected")
	}

	return
}

func operatorExp(p *parser.Parser, op0 string) string {
	offset := p.Offset
	op1, _ := p.Find(`^\s*([+\-*\\])`)

	if !(operators[op0].priority < operators[op1].priority) {
		p.Offset = offset
		return ""
	}

	return op1
}
