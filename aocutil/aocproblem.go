package aocutil

import (
	"fmt"
)

const (
	PROBLEM_INPUT_SAMPLE = "sample"
	PROBLEM_INPUT_FULL   = "full"
)

type AdventOfCodeInterface interface {
	Run(string) string
	Id() string
}

type AdventOfCodeProblem struct {
	sample *ProblemInput
	full   *ProblemInput
	id     string

	AdventOfCodeInterface
}

func NewProblem(day int, samplePath string, fullPath string) (*AdventOfCodeProblem, error) {
	sip, err := BuildFileScanner(samplePath)
	if err != nil {
		return nil, fmt.Errorf("failed to build sample input: %w", err)
	}

	fip, err := BuildInputFromSession(day)
	if err != nil {
		return nil, fmt.Errorf("failed to build full input: %w", err)
	}

	return &AdventOfCodeProblem{
		id:     BuildSelector(day),
		sample: sip,
		full:   fip,
	}, nil
}
