package aocutil

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
)

type ProblemRunner = func(context.Context, int, string) error

type ProblemInput struct {
	s *bufio.Scanner
	c func() error
}

func BuildSelector(day int) string {
	return fmt.Sprintf("day%02d", day)
}

func buildURLForDay(day int) string {
	return fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day)
}

func BuildFileScanner(input string) (*ProblemInput, error) {
	file, err := os.Open(fmt.Sprintf("./%s", input))
	if err != nil {
		return nil, fmt.Errorf("error in opening file handle: %w", err)
	}

	return &ProblemInput{bufio.NewScanner(file), file.Close}, nil
}

func BuildInputFromSession(day int) (*ProblemInput, error) {
	req, err := http.NewRequest("GET", buildURLForDay(day), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build GET request: %w", err)
	}

	si, err := BuildFileScanner("./SESSION")
	if err != nil {
		return nil, fmt.Errorf("failed to get session secret: %w", err)
	}
	defer si.c()

	sessionValue := si.s.Text()

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionValue})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do client request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get 200 OK: %v", resp.StatusCode)
	}

	return &ProblemInput{bufio.NewScanner(resp.Body), resp.Body.Close}, nil
}
