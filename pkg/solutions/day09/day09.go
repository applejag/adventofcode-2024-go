package day09

import (
	"bytes"
	"fmt"
	"io"
	"slices"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
)

type Day struct{}

var _ solutions.Day = Day{}

const EmptyChunk int = -1

func (Day) Part1(file io.Reader) (any, error) {
	digits, err := ParseDigits(file)
	if err != nil {
		return nil, err
	}
	chunks := DigitsToChunks(digits)
	nextEmptyIndex := slices.Index(chunks, EmptyChunk)
	lastValueIndex := FindLastValueIndex(chunks)

	for nextEmptyIndex != -1 && lastValueIndex != -1 && lastValueIndex > nextEmptyIndex {
		chunks[nextEmptyIndex] = chunks[lastValueIndex]
		chunks = chunks[:lastValueIndex]

		nextEmptyIndex = slices.Index(chunks, EmptyChunk)
		lastValueIndex = FindLastValueIndex(chunks)
	}
	return Checksum(chunks), nil
}

func (Day) Part2(file io.Reader) (any, error) {
	return nil, solutions.ErrNotImplemented
}

func Checksum(chunks []int) int {
	var sum int
	for i, fileID := range chunks {
		if fileID == EmptyChunk {
			continue
		}
		sum += i * fileID
	}
	return sum
}

func FindLastValueIndex(chunks []int) int {
	for i := len(chunks) - 1; i > 0; i-- {
		if chunks[i] != EmptyChunk {
			return i
		}
	}
	return -1
}

func DigitsToChunks(digits []byte) []int {
	var size int
	for _, d := range digits {
		size += int(d)
	}
	chunks := make([]int, size)
	var next int
	for i, d := range digits {
		for j := range d {
			if i%2 == 0 {
				chunks[next+int(j)] = i / 2
			} else {
				chunks[next+int(j)] = EmptyChunk
			}
		}
		next += int(d)
	}
	return chunks
}

func ParseDigits(file io.Reader) ([]byte, error) {
	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	b = bytes.TrimSpace(b)
	for i, char := range b {
		if char >= '0' && char <= '9' {
			b[i] = char - '0'
		} else {
			return nil, fmt.Errorf("invalid symbol: %q", char)
		}
	}
	return b, nil
}
