package day03

import (
	"errors"
	"strconv"
)

type Binary struct {
	raw string
	arr []int
}

var (
	ErrInvalidRune    = errors.New("invalid rune; valid runes are '0' and '1'")
	ErrAtoiFailed     = errors.New("failed to convert to int")
	ErrLengthMismatch = errors.New("collection length must match binary length")
	ErrOutofRange     = errors.New("col outside of collection length range")
)

func MakeBinaryFromString(input string) (*Binary, error) {
	b := &Binary{}
	b.raw = input

	for _, v := range input {
		if v != '0' && v != '1' {
			return nil, ErrInvalidRune
		}

		vi, err := strconv.Atoi(string(v))
		if err != nil {
			return nil, ErrAtoiFailed
		}

		b.arr = append(b.arr, vi)
	}

	return b, nil
}

func MakeBinaryFromIntArray(input []int) (*Binary, error) {
	// TODO: finish making from arr
	return nil, nil
}

type BinaryCollection struct {
	bins []Binary
	len  int
}

func MakeBinaryCollection(length int) *BinaryCollection {
	return &BinaryCollection{
		bins: []Binary{},
		len:  length,
	}
}

func (c *BinaryCollection) AddBinary(b Binary) error {
	if len(b.arr) != c.len {
		return ErrLengthMismatch
	}

	c.bins = append(c.bins, b)
	return nil
}

type Stat struct {
	count0 int
	count1 int
}

func (s *Stat) getMostCommon() int {
	if s.count0 >= s.count1 {
		return 0
	}

	return 1
}

func (s *Stat) getLeastCommon() int {
	if s.count0 < s.count1 {
		return 0
	}

	return 1
}

func (c *BinaryCollection) statCol(col int) (*Stat, error) {
	if col < 0 || col >= c.len {
		return nil, ErrOutofRange
	}

	s := &Stat{}

	for _, bin := range c.bins {
		v := bin.arr[col]
		switch v {
		case 0:
			s.count0 += 1
		case 1:
			s.count1 += 1
		}
	}

	return s, nil
}

type GammeEpisilon struct {
	gamma    Binary
	episilon Binary
}

func (c *BinaryCollection) GetGammaEpisilon() *GammeEpisilon {
	ge := &GammeEpisilon{}

	// TODO: calc ge
	return ge
}
