package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	data, err := os.ReadFile("./words.txt")

	if err != nil {
		log.Fatal(err)
	}

	wordsCount := CountWords(data)

	fmt.Println(wordsCount)
}

func CountWords(data []byte) int {
	wordsCount := 0

	if len(data) > 0 {
		wordsCount++
	}

	for _, value := range data {
		if value == ' ' {
			wordsCount++
		}
	}

	return wordsCount
}
