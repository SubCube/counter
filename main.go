package main

import (
	"fmt"
	"os"
)

func main() {

	totalWords := 0
	filenames := os.Args[1:]

	isErrorOccurred := false

	for _, filename := range filenames {
		wordCount, err := CountWordsInFile(filename)
		if err != nil {
			isErrorOccurred = true
			fmt.Fprintln(os.Stderr, "counter:", err)
			continue
		}
		fmt.Println(wordCount, filename)
		totalWords += wordCount
	}

	if len(filenames) == 0 {
		// handle scd input
		words := CountWords(os.Stdin)
		fmt.Println(words, "stdin")
	}

	if len(filenames) > 1 {
		fmt.Println("Total:", totalWords)
	}

	if isErrorOccurred {
		os.Exit(1)
	}
}
