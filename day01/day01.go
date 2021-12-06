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

func Run(ctx context.Context, log *logrus.Logger, part int, input string) error {
	switch part {
	case 1:
		return part1(ctx, log, input)
	case 2:
		return part2(ctx, log, input)
	default:
		return fmt.Errorf("unrecognized mode: %v", part)
	}
}

func part1(ctx context.Context, log *logrus.Logger, input string) error {
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

type sw struct {
	values []int
	label  string
}

func makeSW(label string) *sw {
	return &sw{
		values: make([]int, 0),
		label:  label,
	}
}

func (s *sw) isFull() bool {
	return len(s.values) == 3
}

func (s *sw) append(next int) error {
	if s.isFull() {
		return fmt.Errorf("cannot append; sliding window is full")
	}

	s.values = append(s.values, next)
	return nil
}

func part2(ctx context.Context, log *logrus.Logger, input string) error {
	scan, closer, err := aocutil.BuildFileScanner(input)
	if err != nil {
		return fmt.Errorf("failed to build file scanner: %w", err)
	}
	defer closer()

	windows := make([]sw, 0)
	for scan.Scan() {
		latest := windows[len(windows)-1]
		if latest.isFull() {
			windows = append(windows, *makeSW(strconv.Itoa(len(windows))))
		}

		value, err := strconv.Atoi(scan.Text())
		if err != nil {
			log.Fatalf("failed to parse to int: %s", scan.Text())
		}
		latest.append(value)
	}

	for _, window := range windows {
		log.Infof("%d: %v", window.label, window.values)
	}

	return nil
}
