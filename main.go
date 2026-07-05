package main

import (
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

	/*
		data, err := io.ReadAll(file)

		if err != nil {
			log.Fatalln(err)
		}

		wordsCount := CountWords(data)

		fmt.Println(wordsCount)
	*/

	PrintFileContent(file)
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}

func PrintFileContent(file *os.File) {
	bufferSize := 4096
	buffer := make([]byte, bufferSize)
	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}
		_ = size
		fmt.Print(string(buffer[:size]))
	}
}
