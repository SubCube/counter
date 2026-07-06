package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("./words.txt")

	if err != nil {
		log.Fatalln(err)
	}

	wordCount := CountWordsInFile(file)
	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}

func CountWordsInFile(file *os.File) int {
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
