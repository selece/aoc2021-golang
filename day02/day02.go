package day02

import (
	"context"
	"errors"
	"fmt"

	"github.com/selece/aoc2021-golang/aocutil"
	"github.com/sirupsen/logrus"
)

const (
	DAY02_SELECTOR = "day02"
)

var (
	ErrRunUnrecognizedPart = errors.New("unrecognized part selection")
)

func Run(ctx context.Context, part int, input string) error {
	switch part {
	case 1:
		return part1(ctx, input)

	default:
		return ErrRunUnrecognizedPart
	}
}

func part1(ctx context.Context, input string) error {
	log := logrus.WithContext(ctx)
	pos := &Vector{}

	scan, closer, err := aocutil.BuildFileScanner(input)
	if err != nil {
		return fmt.Errorf("failed to build file scanner: %w", err)
	}
	defer closer()

	for scan.Scan() {
		v, err := StringToVector(scan.Text())
		if err != nil {
			return fmt.Errorf("failed to parse to movement: %v", scan.Text())
		}

		log.Debugf("move with vec: %v", v)
		pos = pos.Add(*v)
	}

	log.Infof("final position: %v", pos)
	return nil
}
