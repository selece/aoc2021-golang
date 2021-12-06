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

type WindowManager struct {
	windows []SlidingWindow
}

func MakeWindowManager() *WindowManager {
	return &WindowManager{
		windows: []SlidingWindow{},
	}
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

func (m *WindowManager) addRecurse(index int, value int) error {
	if index >= len(m.windows) || index < 0 {
		return fmt.Errorf("index out of range: %d", index)
	}

	for i, w := range m.windows {
		if i >= index {
			break
		}

		if !w.isFull() {
			m.addAt(i, value)
		}
	}

	return nil
}

func (m *WindowManager) addWindow(label string) {
	m.windows = append(m.windows, *makeSlidingWindow(label))
}

func (m *WindowManager) getLatestWindow() (int, *SlidingWindow) {
	latestIndex := len(m.windows) - 1

	if latestIndex == -1 {
		return -1, nil
	}

	latestWindow := m.windows[latestIndex]

	return latestIndex, &latestWindow
}

func (m *WindowManager) AddValue(value int) error {
	if len(m.windows) == 0 {
		m.windows = append(m.windows, *makeSlidingWindow("0"))
	}

	latestIndex, latestWindow := m.getLatestWindow()

	if latestWindow.isFull() {
		m.addWindow(strconv.Itoa(latestIndex + 1))
		latestIndex, _ = m.getLatestWindow()
	}

	err := m.addAt(latestIndex, value)
	if err != nil {
		return fmt.Errorf("error adding to window: %w", err)
	}

	err = m.addRecurse(latestIndex, value)
	if err != nil {
		return fmt.Errorf("error recurse adding to windows: %w", err)
	}

	return nil
}

func (m *WindowManager) Print() string {
	res := ""
	for _, w := range m.windows {
		res += fmt.Sprintf("%s: %v\n", w.label, w.values)
	}

	return res
}
