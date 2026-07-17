package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

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
