package day07

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
)

type Day struct{}

var _ solutions.Day = Day{}

func (Day) Part1(file io.Reader) (any, error) {
	equations, err := Parse(file)
	if err != nil {
		return nil, err
	}
	var sum int64
	for _, eq := range equations {
		for ops := range eq.PermutationsIter() {
			if eq.Eval(ops) == eq.Result {
				sum += eq.Result
				break
			}
		}
	}
	return sum, nil
}

func (Day) Part2(file io.Reader) (any, error) {
	equations, err := Parse(file)
	if err != nil {
		return nil, err
	}
	var sum int64
	for _, eq := range equations {
		for ops := range eq.PermutationsIterPart2() {
			if eq.Eval(ops) == eq.Result {
				sum += eq.Result
				slog.Debug("works", "operands", eq.Operands, "operators", ops, "result", eq.Result)
				break
			}
		}
	}
	return sum, nil
}

func Parse(file io.Reader) ([]Equation, error) {
	var equations []Equation
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		eq, err := ParseEquation(scanner.Text())
		if err != nil {
			return nil, err
		}
		equations = append(equations, eq)
	}
	return equations, nil
}

func ParseEquation(line string) (Equation, error) {
	before, after, ok := strings.Cut(line, ": ")
	if !ok {
		return Equation{}, fmt.Errorf("invalid format on line: %q", line)
	}
	result, err := strconv.ParseInt(before, 10, 64)
	if err != nil {
		return Equation{}, err
	}
	split := strings.Split(after, " ")
	operands := make([]int64, len(split))
	for i, s := range split {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return Equation{}, err
		}
		operands[i] = num
	}
	return Equation{
		Result:   result,
		Operands: operands,
	}, nil
}

type Equation struct {
	Result   int64
	Operands []int64
}

type Operator byte

const (
	OpAdd Operator = 1 + iota
	OpMul
	OpConcat
)

func (op Operator) String() string {
	switch op {
	case OpAdd:
		return "+"
	case OpMul:
		return "*"
	case OpConcat:
		return "||"
	default:
		panic(fmt.Sprintf("unexpected day07.Operator: %#v", op))
	}
}

func (eq Equation) Eval(operators []Operator) int64 {
	if len(operators) != len(eq.Operands)-1 {
		panic(fmt.Errorf("wrong number of operators, want %d but got %d", len(eq.Operands)-1, len(operators)))
	}
	result := eq.Operands[0]
	for i, num := range eq.Operands[1:] {
		switch operators[i] {
		case OpAdd:
			result += num
		case OpMul:
			result *= num
		case OpConcat:
			result = concat(result, num)
		}
	}
	return result
}

func concat(a, b int64) int64 {
	return int64(math.Pow(10, float64(intLen(b))))*a + b
}

func intLen(n int64) int64 {
	switch {
	case n < 10:
		return 1
	case n < 100:
		return 2
	case n < 1_000:
		return 3
	case n < 10_000:
		return 4
	case n < 100_000:
		return 5
	case n < 1_000_000:
		return 6
	case n < 10_000_000:
		return 7
	case n < 100_000_000:
		return 8
	case n < 1_000_000_000:
		return 9
	case n < 10_000_000_000:
		return 10
	case n < 100_000_000_000:
		return 11
	case n < 1_000_000_000_000:
		return 12
	case n < 10_000_000_000_000:
		return 13
	case n < 100_000_000_000_000:
		return 14
	case n < 1_000_000_000_000_000:
		return 15
	case n < 10_000_000_000_000_000:
		return 16
	case n < 100_000_000_000_000_000:
		return 17
	case n < 1_000_000_000_000_000_000:
		return 18
	default:
		panic(fmt.Errorf("too big number: %d", n))
	}
}

func (eq Equation) PermutationsIter() iter.Seq[[]Operator] {
	return func(yield func([]Operator) bool) {
		operators := make([]Operator, len(eq.Operands)-1)
		for i := range len(operators) {
			operators[i] = OpAdd
		}

		for range 1 << len(operators) {
			if !yield(operators) {
				return
			}

			if !addOneToOperatorSlice(operators, 0) {
				return
			}
		}
	}
}

func addOneToOperatorSlice(operators []Operator, index int) bool {
	if index >= len(operators) {
		return false
	}
	switch operators[index] {
	case OpAdd:
		operators[index] = OpMul
		return true
	case OpMul:
		operators[index] = OpAdd
		return addOneToOperatorSlice(operators, index+1)
	default:
		return false
	}
}

func (eq Equation) PermutationsIterPart2() iter.Seq[[]Operator] {
	return func(yield func([]Operator) bool) {
		operators := make([]Operator, len(eq.Operands)-1)
		for i := range len(operators) {
			operators[i] = OpAdd
		}

		for range 10000000 {
			if !yield(operators) {
				return
			}

			if !addOneToOperatorSlicePart2(operators, 0) {
				return
			}
		}
		panic("too many iterations, possible infinite loop")
	}
}

func addOneToOperatorSlicePart2(operators []Operator, index int) bool {
	if index >= len(operators) {
		return false
	}
	switch operators[index] {
	case OpAdd:
		operators[index] = OpMul
		return true
	case OpMul:
		operators[index] = OpConcat
		return true
	case OpConcat:
		operators[index] = OpAdd
		return addOneToOperatorSlicePart2(operators, index+1)
	default:
		panic(fmt.Sprintf("unexpected day07.Operator: %#v", operators[index]))
	}
}
