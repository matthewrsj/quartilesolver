package solver

import (
	"fmt"
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

	isInteractive bool
}

// New returns a new *Solver configured by the options functions.
func New(options ...func(*Solver)) *Solver {
	solver := &Solver{
		wordFP:        _WordRelFP,
		isInteractive: false,
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

// WithInteractivityToggle sets the --interactive toggle.
func WithInteractivityToggle(interactive bool) func(*Solver) {
	return func(s *Solver) {
		s.isInteractive = interactive
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
		if err := s.checkAndRecord(fLen1, numFragments); err != nil {
			return nil, 0, fmt.Errorf("check and record: %w", err)
		}

		// two-fragment words
		for j, fragment2 := range fragments {
			if j == i {
				continue
			}

			numFragments = 2
			if err := s.checkAndRecord(fLen1+fragment2, numFragments); err != nil {
				return nil, 0, fmt.Errorf("check and record: %w", err)
			}

			// three-fragment words
			for k, fragment3 := range fragments {
				if k == i || k == j {
					continue
				}

				numFragments = 3
				if err := s.checkAndRecord(fLen1+fragment2+fragment3, numFragments); err != nil {
					return nil, 0, fmt.Errorf("check and record: %w", err)
				}

				// four-fragment words
				for l, fragment4 := range fragments {
					if l == i || l == j || i == k {
						continue
					}

					numFragments = 4
					if err := s.checkAndRecord(fLen1+fragment2+fragment3+fragment4, numFragments); err != nil {
						return nil, 0, fmt.Errorf("check and record: %w", err)
					}
				}
			}
		}
	}

	// apply quartile bonus if we have found all of them
	s.solution.addQuartileBonusIfApplicable()

	return s.solution.words, s.solution.score, nil
}

func (s *Solver) checkAndRecord(base string, numFragments int) error {
	if !wordInDict(base, s.problem.dictionary) {
		return nil
	}

	var (
		ok, del bool
		err     error
	)

	if s.isInteractive {
		if ok, err = isWordSolution(base); err != nil {
			return fmt.Errorf("is word solution: %w", err)
		}

		if !ok {
			if del, err = confirmDelete(); err != nil {
				return fmt.Errorf("confirm delete: %w", err)
			}

			if del {
				if err = s.RemoveWordFromDict(base); err != nil {
					return fmt.Errorf("remove word from dict: %w", err)
				}
			}

			return nil
		}
	}

	s.solution.addWord(base)
	s.solution.updateScore(numFragments)

	return nil
}

func wordInDict(word string, dict map[string]struct{}) bool {
	_, ok := dict[word]

	return ok
}
