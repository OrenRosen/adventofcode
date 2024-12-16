package textutil

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseLines(file string) ([]string, error) {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	return lines, nil
}

func GetMatrix(file string) ([][]string, error) {
	lines, err := ParseLines(file)
	if err != nil {
		return nil, err
	}

	matrix := [][]string{}
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}

	return matrix, nil
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}
