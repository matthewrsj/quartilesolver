package solver

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// isWordSolution asks the user if the word was a solution in the puzzle.
func isWordSolution(word string) (bool, error) {
	s, err := getInputFromReader(
		*bufio.NewReader(os.Stdin),
		fmt.Sprintf("Is %s a solution? (Y/n)", word),
	)

	return isYesDefaultResponse(s), err
}

func confirmDelete() (bool, error) {
	s, err := getInputFromReader(
		*bufio.NewReader(os.Stdin),
		"Are you sure you want to delete this word from the dictionary? (y/N)",
	)

	return isNoDefaultResponse(s), err
}

func isYesDefaultResponse(s string) bool {
	return isBoolResponse(s, true)
}

func isNoDefaultResponse(s string) bool {
	return isBoolResponse(s, false)
}

func isBoolResponse(input string, def bool) bool {
	cleanInput := strings.TrimSpace(strings.ToLower(input))

	if cleanInput == "" {
		return def
	}

	return strings.HasPrefix(cleanInput, "y")
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

	return s, nil
}
