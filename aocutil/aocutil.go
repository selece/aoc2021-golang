package aocutil

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

type ProblemRunner = func(context.Context, int, string) error

func BuildSelector(day int) string {
	return fmt.Sprintf("day%02d", day)
}

func BuildFileScanner(input string) (*bufio.Scanner, func() error, error) {
	file, err := os.Open(fmt.Sprintf("./%s", input))
	if err != nil {
		return nil, nil, fmt.Errorf("error in opening file handle: %w", err)
	}

	return bufio.NewScanner(file), file.Close, nil
}
