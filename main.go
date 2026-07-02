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

	fmt.Println("Data:", string(data))
	fmt.Println("Data length is", len(data))
}
