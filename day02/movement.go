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

	DirectionToVectorMap = map[Direction]*Vector{
		Up:       &Vector{0, -1},
		Down:     &Vector{0, 1},
		Forward:  &Vector{1, 0},
		Backward: &Vector{-1, 0},
	}
)

type Movement struct {
	dir  Direction
	dist int
}

type Vector struct {
	x int
	y int
}

var (
	ErrNoRegexpMatch = errors.New("no match for movement regexp; check formatting")
	ErrFailedToMap   = errors.New("no entry on movement map")
	ErrFailedAtoi    = errors.New("failed to parse string to int")
)

func MakeMovement(input string) (*Movement, error) {
	r, _ := regexp.Compile("(forward|backwards|up|down) ([0-9]*)")
	results := r.FindStringSubmatch(input)

	if results == nil {
		return nil, ErrNoRegexpMatch
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

func (m *Movement) ToVector() *Vector {
	v := DirectionToVectorMap[m.dir]
	switch m.dir {
	case Up:
	case Down:
		v.y = v.y * m.dist
	case Forward:
	case Backward:
		v.x = v.x * m.dist
	}

	return v
}

func (v *Vector) Add(v2 Vector) *Vector {
	return &Vector{
		x: v.x + v2.x,
		y: v.y + v2.y,
	}
}

func (v *Vector) Move(m Movement) *Vector {
	v2 := m.ToVector()
	return v.Add(*v2)
}
