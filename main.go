package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/selece/aoc2021-golang/aocutil"
	"github.com/selece/aoc2021-golang/day01"

	"github.com/sirupsen/logrus"
)

type cargs struct {
	day   int
	part  int
	input string
}

func extractArgs(args []string) (*cargs, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("incorrect number of args, expected 3 but received: %v", args)
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("unrecognized format for day: %w", err)
	}

	part, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, fmt.Errorf("unrecognized format for part: %w", err)
	}

	input := args[2]

	return &cargs{day, part, input}, nil
}

func main() {
	// setup logging
	ctx := context.Background()
	log := logrus.StandardLogger()

	// map input args to modules
	problemsMap := make(map[string]aocutil.ProblemRunner)
	problemsMap[day01.DAY01_SELECTOR] = day01.Run

	// nab the args, minus the program name
	args := os.Args[1:]
	pargs, err := extractArgs(args)

	if err != nil {
		log.Fatalf("failed to parse args: %w", err)
	}

	// search the map for our specified day
	if runner, found := problemsMap[aocutil.BuildSelector(pargs.day)]; found {
		if err = runner(ctx, log, pargs.part, pargs.input); err != nil {
			log.Fatalf("error in runner: %w", err)
		}
	} else {
		log.Fatalf("unrecognized day: %d", pargs.day)
	}
}
