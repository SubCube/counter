package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("Error: no filename provided")
	}

	totalWords := 0
	filenames := os.Args[1:]

	for _, filename := range filenames {
		wordCount := CountWordsInFile(filename)
		fmt.Println(wordCount, filename)
		totalWords += wordCount
	}

	if len(filenames) > 1 {
		fmt.Println("Total:", totalWords)
	}
}

func CountWordsInFile(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln(err)
	}

	return CountWords(file)
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
