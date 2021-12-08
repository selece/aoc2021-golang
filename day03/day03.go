package day03

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	DAY03_SELECTOR = "day03"
)

var (
	ErrRunUnrecognizedPart = errors.New("unrecognized part selection")
)

func Run(ctx context.Context, part int, input string) error {
	switch part {
	case 1:
		return part1(ctx, input)
	}

	return ErrRunUnrecognizedPart
}

func part1(ctx context.Context, input string) error {
	log := logrus.WithContext(ctx)

	test := "1010"
	bin, err := MakeBinaryFromString(test)
	if err != nil {
		return fmt.Errorf("failed to make binary: %w", err)
	}

	col := MakeBinaryCollection(4)
	col.AddBinary(*bin)

	test2 := "1111"
	bin2, err := MakeBinaryFromString(test2)
	if err != nil {
		return fmt.Errorf("failed to make binary: %w", err)
	}

	col.AddBinary(*bin2)

	s, err := col.StatCol(3)
	if err != nil {
		return fmt.Errorf("error calculating stats for col: %w", err)
	}

	log.Infof("%v", s)

	return nil
}
