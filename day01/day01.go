package day01

import (
	"context"
	"fmt"
	"strconv"

	"github.com/selece/aoc2021-golang/aocutil"

	"github.com/sirupsen/logrus"
)

const (
	DAY01_SELECTOR = "day01"
)

func Run(ctx context.Context, part int, input string) error {
	switch part {
	case 1:
		return part1(ctx, input)
	case 2:
		return part2(ctx, input)
	default:
		return fmt.Errorf("unrecognized mode: %v", part)
	}
}

func part1(ctx context.Context, input string) error {
	log := logrus.WithContext(ctx)
	scan, closer, err := aocutil.BuildFileScanner(input)
	if err != nil {
		return fmt.Errorf("failed to build file scanner: %w", err)
	}
	defer closer()

	prev := -1
	count := 0
	for scan.Scan() {
		current, err := strconv.Atoi(scan.Text())
		if err != nil {
			return fmt.Errorf("failed to convert entry: %s", scan.Text())
		}

		msg := ""
		if prev == -1 {
			msg = "(N/A - no previous measurement)"
		} else if current > prev {
			msg = "(increased)"
			count += 1
		} else {
			msg = "(decreased)"
		}
		prev = current

		log.Infof("%v %s\n", current, msg)
	}

	if err := scan.Err(); err != nil {
		return fmt.Errorf("error in scaning file: %w", scan.Err())
	}

	log.Infof("total increases counted: %v\n", count)
	return nil
}

func part2(ctx context.Context, input string) error {
	log := logrus.WithContext(ctx)

	scan, closer, err := aocutil.BuildFileScanner(input)
	if err != nil {
		return fmt.Errorf("failed to build file scanner: %w", err)
	}
	defer closer()

	wm := MakeWindowManager()
	index := 0

	for scan.Scan() {
		value, err := strconv.Atoi(scan.Text())
		if err != nil {
			log.Fatalf("failed to parse to int: %s", scan.Text())
		}

		wm.AddValue(value, index)
		index += 1
	}

	wm.TrimInvalidWindows()

	prev := -1
	count := 0
	for _, w := range wm.windows {
		current := w.sum()
		if current > prev && prev != -1 {
			count += 1
		}
		prev = current
	}

	// print
	log.Info(count)

	return nil
}
