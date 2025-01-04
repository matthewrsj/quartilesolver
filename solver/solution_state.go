package solver

import (
	"sort"
)

type solutionState struct {
	words                []string
	score, quartileCount int
}

func (ss *solutionState) addWord(word string) {
	i := sort.SearchStrings(ss.words, word)
	ss.words = append(ss.words, "")
	copy(ss.words[i+1:], ss.words[i:])
	ss.words[i] = word
}

func (ss *solutionState) updateScore(numFragments int) {
	fss := getFragmentScores()
	if numFragments-1 > len(fss) || numFragments-1 < 0 {
		return
	}

	if numFragments == _FragmentsInQuartile {
		ss.quartileCount++
	}

	ss.score += getFragmentScores()[numFragments-1]
}

func (ss *solutionState) addQuartileBonusIfApplicable() {
	if ss.quartileCount >= _QuartileBonusThreshold {
		ss.score += _QuartileBonus
	}
}
