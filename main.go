package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
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
	bufferSize := 4096
	isInsideWord := false
	buffer := make([]byte, bufferSize)
	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}
		isInsideWord = !unicode.IsSpace(rune(buffer[0])) && isInsideWord
		bufferCount := CountWords(buffer[:size])
		if isInsideWord {
			bufferCount -= 1
		}
		wordsCount += bufferCount

		isInsideWord = !unicode.IsSpace(rune(buffer[size-1]))
	}

	return wordsCount
}
