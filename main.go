package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}
	defer file.Close()

	count := CountWords(file)

	return count, nil
}

func CountWords(file io.Reader) int {
	wordsCount := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordsCount++
	}

	if err := scanner.Err(); err != nil {
		log.Printf("error scanning: %v", err)
	}

	return wordsCount
}
