package day02

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

const (
	DAY02_SELECTOR = "day02"
)

var (
	RUN_UNRECOGNIZED_PART = errors.New("unrecognized part selection")
)

func Run(ctx context.Context, part int, input string) error {
	switch part {
	case 1:
		return part1(ctx, input)

	default:
		return RUN_UNRECOGNIZED_PART
	}
}

func part1(ctx context.Context, input string) error {
	log := logrus.WithContext(ctx)

	mov, err := MakeMovement("up 7")
	if err != nil {
		return err
	}

	log.Infof("%v", mov)

	return nil
}
