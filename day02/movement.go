package day02

import (
	"errors"
	"regexp"
	"strconv"
)

type Direction int

const (
	Up = iota
	Down
	Forward
	Backward
)

var (
	DirectionMap = map[string]Direction{
		"up":       Up,
		"down":     Down,
		"forward":  Forward,
		"backward": Backward,
	}

	DirectionModifierMap = map[Direction]int{
		Up:       1,
		Down:     -1,
		Forward:  1,
		Backward: -1,
	}
)

type Movement struct {
	dir  Direction
	dist int
}

type Position struct {
	h int
	v int
}

var (
	ErrMovementNoRegexpMatch = errors.New("no match for movement regexp; check formatting")
	ErrFailedToMap           = errors.New("no entry on movement map")
	ErrFailedAtoi            = errors.New("failed to parse string to int")
)

func MakeMovement(input string) (*Movement, error) {
	r, _ := regexp.Compile("(forward|backwards|up|down) ([0-9]*)")
	results := r.FindStringSubmatch(input)

	if results == nil {
		return nil, ErrMovementNoRegexpMatch
	}

	if _, ok := DirectionMap[results[1]]; !ok {
		return nil, ErrFailedToMap
	}

	dir := DirectionMap[results[1]]
	dist, err := strconv.Atoi(results[2])

	if err != nil {
		return nil, ErrFailedAtoi
	}

	return &Movement{dir, dist}, nil
}
