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
	case 2:
		return solve(ctx, input, part)
	}
	return ErrRunUnrecognizedPart
}

func solve(ctx context.Context, input string, part int) error {
	log := logrus.WithContext(ctx)
	pos1 := &Vector{}
	pos2 := &AimVector{}

	scan, closer, err := aocutil.BuildFileScanner(input)
	if err != nil {
		return fmt.Errorf("failed to build file scanner: %w", err)
	}
	defer closer()

	for scan.Scan() {
		if part == 1 {
			v, err := StringToVector(scan.Text())
			if err != nil {
				return fmt.Errorf("failed to parse to vector: %v", scan.Text())
			}

			log.Debugf("move with vec: %v", v)
			pos1 = pos1.Move(*v)
		}

		if part == 2 {
			v, err := StringToAimVector(scan.Text())
			if err != nil {
				return fmt.Errorf("failwed to parse to aimvec: %v", scan.Text())
			}

			log.Debugf("move with aimvec: %v", v)
			pos2 = pos2.Move(*v)
		}
	}

	if part == 1 {
		log.Infof("final position: %v", pos1)
	}

	if part == 2 {
		log.Infof("final position: %v", pos2)
	}
	return nil
}
