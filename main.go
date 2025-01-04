package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"quartilesolver/solver"
)

const defaultWordRelFP = "./assets/usa2.txt"

type fragmentArgs []string

func (f *fragmentArgs) String() string {
	return strings.Join(*f, ", ")
}

func (f *fragmentArgs) Set(value string) error {
	*f = strings.Split(value, " ")

	return nil
}

func main() {
	var (
		inFile   = flag.String("in", "", "input file containing line-delineated word fragments")
		wordFile = flag.String("words", defaultWordRelFP, "file containing line-delineated dictionary of words")

		fragments fragmentArgs
		err       error
	)

	flag.Var(&fragments, "fragments", "space-delineated list of fragments")
	flag.Parse()

	fmt.Println("Fragments: ", fragments)

	if *inFile != "" {
		if fragments, err = readInFile(*inFile); err != nil {
			fmt.Print(fmt.Errorf("read in-file: %w", err))
			os.Exit(1)
		}
	}

	solution, score, err := solver.New(solver.WithWordFP(*wordFile)).Solve(fragments)
	if err != nil {
		fmt.Print(fmt.Errorf("solve puzzle: %w", err))
		os.Exit(1)
	}

	fmt.Println("Total words found: ", len(solution))
	fmt.Println("Expected score: ", score)
	pprint(solution)
}

func pprint(ss []string) {
	for _, s := range ss {
		fmt.Println(s)
	}
}

func readInFile(path string) ([]string, error) {
	fp, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("get filepath: %w", err)
	}

	wordStream, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var fragmentStrs []string

	fragmentBytes := bytes.Split(wordStream, []byte{'\n'})
	for _, fragment := range fragmentBytes {
		// clean the fragments
		if f := strings.TrimSpace(string(fragment)); f != "" {
			fragmentStrs = append(fragmentStrs, string(fragment))
		}
	}

	return fragmentStrs, nil
}
