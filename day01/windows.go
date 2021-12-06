package day01

import (
	"fmt"
	"strconv"
)

const (
	WINDOW_CAPACITY = 3
)

type SlidingWindow struct {
	values []int
	label  string
}

func makeSlidingWindow(label string) *SlidingWindow {
	return &SlidingWindow{
		values: []int{},
		label:  label,
	}
}

func (s *SlidingWindow) isFull() bool {
	return len(s.values) == WINDOW_CAPACITY
}

func (s *SlidingWindow) add(next int) (*SlidingWindow, error) {
	if s.isFull() {
		return nil, fmt.Errorf("cannot append; sliding window is full")
	}

	s.values = append(s.values, next)
	return s, nil
}

func (s *SlidingWindow) sum() int {
	sum := 0
	for _, v := range s.values {
		sum += v
	}

	return sum
}

type WindowManager struct {
	windows []SlidingWindow
}

func MakeWindowManager() *WindowManager {
	wm := &WindowManager{
		windows: []SlidingWindow{},
	}

	wm.addWindow()
	return wm
}

func (m *WindowManager) addAt(index int, value int) error {
	if index >= len(m.windows) || index < 0 {
		return fmt.Errorf("index out of range: %d", index)
	}

	if m.windows[index].isFull() {
		return fmt.Errorf("window at index is full: %d", index)
	}

	next, err := m.windows[index].add(value)
	if err != nil {
		return fmt.Errorf("failed to add: %w", err)
	}

	m.windows[index] = *next
	return nil
}

func (m *WindowManager) addWindow() {
	label := strconv.Itoa(len(m.windows) + 1)
	m.windows = append(m.windows, *makeSlidingWindow(label))
}

func (m *WindowManager) AddValue(value int, index int) error {
	if len(m.windows) < index+WINDOW_CAPACITY {
		for i := 0; i <= index+WINDOW_CAPACITY-len(m.windows); i++ {
			m.addWindow()
		}
	}

	for i := index; i <= index+WINDOW_CAPACITY; i++ {
		m.addAt(i, value)
	}

	return nil
}

func (m *WindowManager) TrimInvalidWindows() error {
	trim := []SlidingWindow{}
	for _, w := range m.windows {
		if w.isFull() {
			trim = append(trim, w)
		}
	}

	m.windows = trim
	return nil
}

func (m *WindowManager) Print() string {
	res := "-- WindowManager output --\n"
	for _, w := range m.windows {
		res += fmt.Sprintf("%s: %v\n", w.label, w.values)
	}

	return res
}
