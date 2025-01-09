package solver

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (s *Solver) writeWordsToDict(words []string) error {
	fp, err := filepath.Abs(s.wordFP)
	if err != nil {
		return fmt.Errorf("get filepath: %w", err)
	}

	if err := os.WriteFile(fp, []byte(strings.Join(words, "\n")), os.ModePerm); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func (s *Solver) RemoveWordFromDict(word string) error {
	words, err := s.getDictContents()
	if err != nil {
		return fmt.Errorf("get dictionary contents: %w", err)
	}

	var i int

	for _, current := range words {
		if current != word {
			words[i] = current
			i++
		}
	}

	if i != len(words) {
		words = words[:i]
		if err := s.writeWordsToDict(words); err != nil {
			return fmt.Errorf("write words to dict: %w", err)
		}
	}

	// not found
	return nil
}

func (s *Solver) getDictContents() ([]string, error) {
	fp, err := filepath.Abs(s.wordFP)
	if err != nil {
		return nil, fmt.Errorf("get filepath: %w", err)
	}

	wordStream, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	words := bytes.Split(wordStream, []byte{'\n'})
	wordList := make([]string, len(words))
	for i, word := range words {
		wordList[i] = string(word)
	}

	return wordList, nil
}

// getDict reads a dictionary from disk.
func (s *Solver) getDict() error {
	words, err := s.getDictContents()
	if err != nil {
		return fmt.Errorf("get dictionary contents: %w", err)
	}

	wordMap := make(map[string]struct{}, len(words))

	for i := 0; i < len(words); i++ {
		wordMap[words[i]] = struct{}{}
	}

	s.problem.dictionary = wordMap

	return nil
}
