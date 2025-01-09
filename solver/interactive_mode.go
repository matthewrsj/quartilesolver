package solver

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// isWordSolution asks the user if the word was a solution in the puzzle.
func isWordSolution(word string) (bool, error) {
	var (
		s   string
		err error
	)

	r := bufio.NewReader(os.Stdin)
	if s, err = getInputFromReader(*r, fmt.Sprintf("Is %s a solution? (y/n)", word)); err != nil {
		return false, fmt.Errorf("get input from reader: %w", err)
	}

	return isYesResponse(s), nil
}

func confirmDelete() (bool, error) {
	s, err := getInputFromReader(
		*bufio.NewReader(os.Stdin),
		"Are you sure you want to delete this word from the dictionary? (y/n)",
	)

	return isYesResponse(s), err
}

func isYesResponse(s string) bool {
	return strings.HasPrefix(strings.ToLower(s), "y")
}

func getInputFromReader(reader bufio.Reader, msg string) (string, error) {
	var (
		s   string
		err error
	)

	if _, err = fmt.Fprint(os.Stderr, msg+" "); err != nil {
		return "", fmt.Errorf("print to stderr: %w", err)
	}

	if s, err = reader.ReadString('\n'); err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	if s == "" {
		return "", errors.New("empty input")
	}

	return s, nil
}
