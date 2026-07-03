package main

import (
	"bytes"
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
	words := bytes.Fields(data)
	return len(words)
}
