package day07

import (
	"bufio"
	"fmt"
	"io"
	"iter"
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
	var sum int
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
	return nil, solutions.ErrNotImplemented
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
	result, err := strconv.Atoi(before)
	if err != nil {
		return Equation{}, err
	}
	split := strings.Split(after, " ")
	operands := make([]int, len(split))
	for i, s := range split {
		num, err := strconv.Atoi(s)
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
	Result   int
	Operands []int
}

type Operator byte

const (
	OpAdd Operator = 1 + iota
	OpMul
)

func (eq Equation) Eval(operators []Operator) int {
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
		}
	}
	return result
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
