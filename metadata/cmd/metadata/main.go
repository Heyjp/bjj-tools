package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must enter in a file location")
		return
	}

	if len(os.Args) != 3 {
		log.Fatal("Please enter a metadata file and a chapters file")
	}
	metadata := os.Args[1]
	chaptersFile := os.Args[2]

	TestMetaData(metadata)
	ClearVideoMetaData(metadata)
	chapters := CreateChapters(chaptersFile)
	WriteChaptersToMetaData(metadata, chapters)
}
