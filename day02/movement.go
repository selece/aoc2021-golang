package day02

import (
	"errors"
	"regexp"
	"strconv"
)

type Vector struct {
	x int
	y int
}

var (
	UnitVectorMap = map[string]Vector{
		"up":       {0, -1},
		"down":     {0, 1},
		"forward":  {1, 0},
		"backward": {-1, 0},
	}
)

var (
	ErrNoRegexpMatch = errors.New("no match for movement regexp; check formatting")
	ErrFailedToMap   = errors.New("no entry on movement map")
	ErrFailedAtoi    = errors.New("failed to parse string to int")
)

func StringToVector(input string) (*Vector, error) {
	r, _ := regexp.Compile("(forward|backwards|up|down) ([0-9]*)")
	results := r.FindStringSubmatch(input)

	if results == nil {
		return nil, ErrNoRegexpMatch
	}

	if _, ok := UnitVectorMap[results[1]]; !ok {
		return nil, ErrFailedToMap
	}

	v := Vector{}
	v = UnitVectorMap[results[1]]

	mag, err := strconv.Atoi(results[2])

	if err != nil {
		return nil, ErrFailedAtoi
	}

	v.x = v.x * mag
	v.y = v.y * mag

	return &v, nil
}

func (v *Vector) Add(v2 Vector) *Vector {
	return &Vector{
		x: v.x + v2.x,
		y: v.y + v2.y,
	}
}
