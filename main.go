package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("./words.txt")

	if err != nil {
		log.Fatalln(err)
	}

	wordCount := CountWords(file)
	fmt.Println(wordCount)
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
