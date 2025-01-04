package solver

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

const (
	_WordRelFP = "./assets/usa2.txt"

	_QuartileBonus          = 40
	_QuartileBonusThreshold = 5
	_FragmentsInQuartile    = 4
)

type Solver struct {
	solution solutionState
	problem  problemState
	wordFP   string
}

// New returns a new *Solver configured by the options functions.
func New(options ...func(*Solver)) *Solver {
	solver := &Solver{
		wordFP: _WordRelFP,
		solution: solutionState{
			words:         []string{},
			score:         0,
			quartileCount: 0,
		},
		problem: problemState{
			dictionary: wordDict{},
			fragments:  []string{},
		},
	}

	for _, o := range options {
		o(solver)
	}

	return solver
}

// WithWordFP sets the word dictionary filepath for the solver.
func WithWordFP(path string) func(*Solver) {
	return func(s *Solver) {
		s.wordFP = path
	}
}

func getFragmentScores() []int {
	return []int{1, 2, 4, 8}
}

type wordDict map[string]struct{}

// Solve receives a list of strings (quartile fragments) and returns a list of words that can be created
// from those fragments, up to a fragment length of four.
func (s *Solver) Solve(fragments []string) ([]string, int, error) {
	if err := s.getDict(); err != nil {
		return nil, 0, fmt.Errorf("get dictionary: %w", err)
	}

	s.problem.fragments = fragments

	// single-fragment words
	for i, fLen1 := range fragments {
		numFragments := 1
		s.checkAndRecord(fLen1, numFragments)

		// two-fragment words
		for j, fragment2 := range fragments {
			if j == i {
				continue
			}

			numFragments = 2
			s.checkAndRecord(fLen1+fragment2, numFragments)

			// three-fragment words
			for k, fragment3 := range fragments {
				if k == i || k == j {
					continue
				}

				numFragments = 3
				s.checkAndRecord(fLen1+fragment2+fragment3, numFragments)

				// four-fragment words
				for l, fragment4 := range fragments {
					if l == i || l == j || i == k {
						continue
					}

					numFragments = 4
					s.checkAndRecord(fLen1+fragment2+fragment3+fragment4, numFragments)
				}
			}
		}
	}

	// apply quartile bonus if we have found all of them
	s.solution.addQuartileBonusIfApplicable()

	return s.solution.words, s.solution.score, nil
}

// getDict reads a dictionary from disk.
func (s *Solver) getDict() error {
	fp, err := filepath.Abs(s.wordFP)
	if err != nil {
		return fmt.Errorf("get filepath: %w", err)
	}

	wordStream, err := os.ReadFile(fp)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	words := bytes.Split(wordStream, []byte{'\n'})
	wordMap := make(map[string]struct{}, len(words))

	for i := 0; i < len(words); i++ {
		wordMap[string(words[i])] = struct{}{}
	}

	s.problem.dictionary = wordMap

	return nil
}

func (s *Solver) checkAndRecord(base string, numFragments int) {
	if wordInDict(base, s.problem.dictionary) {
		s.solution.addWord(base)
		s.solution.updateScore(numFragments)
	}
}

func wordInDict(word string, dict map[string]struct{}) bool {
	_, ok := dict[word]

	return ok
}
