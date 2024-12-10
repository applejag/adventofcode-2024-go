package day09

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"strings"

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
	lastValueIndex := FindLastFileChunkIndex(chunks)

	for nextEmptyIndex != -1 && lastValueIndex != -1 && lastValueIndex > nextEmptyIndex {
		chunks[nextEmptyIndex] = chunks[lastValueIndex]
		chunks = chunks[:lastValueIndex]

		nextEmptyIndex = slices.Index(chunks, EmptyChunk)
		lastValueIndex = FindLastFileChunkIndex(chunks)
	}
	return Checksum(chunks), nil
}

func (Day) Part2(file io.Reader) (any, error) {
	digits, err := ParseDigits(file)
	if err != nil {
		return nil, err
	}
	chunks := DigitsToChunks(digits)

	nextFileID := chunks[FindLastFileChunkIndex(chunks)]

	for nextFileID != EmptyChunk {
		fileIndex, size := FindFileAndSize(chunks, nextFileID)
		if fileIndex == -1 {
			slog.Debug("Stop because file not found", "fileID", nextFileID)
			break
		}
		emptyIndex := FindEmptyOfSizeIndex(chunks[:fileIndex], size)
		slog.Debug("Find empty of size", "fileID", nextFileID, "size", size, "fileIndex", fileIndex, "emptyIndex", emptyIndex)
		if emptyIndex == -1 {
			slog.Debug("Skip because no empty slot found", "fileID", nextFileID, "size", size)
			nextFileID--
			continue
		}
		copy(chunks[emptyIndex:emptyIndex+size], chunks[fileIndex:fileIndex+size])
		for i := fileIndex; i < fileIndex+size; i++ {
			chunks[i] = EmptyChunk
		}
		nextFileID--
	}

	if slog.Default().Enabled(context.Background(), slog.LevelDebug) {
		slog.Debug("", "chunks", FormatChunks(chunks))
	}

	return Checksum(chunks), nil
}

func FormatChunks(chunks []int) string {
	var sb strings.Builder
	sb.Grow(len(chunks))
	for _, fileID := range chunks {
		if fileID == EmptyChunk {
			sb.WriteByte('.')
		} else if fileID <= 9 {
			sb.WriteByte('0' + byte(fileID))
		} else {
			fmt.Fprintf(&sb, "{%d}", fileID)
		}
	}
	return sb.String()
}

func FindEmptyOfSizeIndex(chunks []int, minSize int) (index int) {
	for i, char := range chunks {
		if char != EmptyChunk {
			continue
		}
		if CountSameCharPrefix(chunks[i:]) < minSize {
			continue
		}
		return i
	}
	return -1
}

func FindFileAndSize(chunks []int, fileID int) (index int, size int) {
	index = slices.Index(chunks, fileID)
	if index == -1 {
		return -1, 0
	}
	return index, CountSameCharPrefix(chunks[index:])
}

func CountSameCharPrefix(chunks []int) (size int) {
	firstValue := chunks[0]
	for i, v := range chunks[1:] {
		if v != firstValue {
			return i + 1
		}
	}
	return len(chunks)
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

func FindLastFileChunkIndex(chunks []int) int {
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
