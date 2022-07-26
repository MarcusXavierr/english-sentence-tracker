package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		usage()
	}

	wordsTrackerFile := getHomeDir() + "/english_words/words.txt"
	trash := getHomeDir() + "/english_words/trash.txt"
	sentence := os.Args[1]
	fmt.Printf("before")
	CheckIfSentenceExists(sentence, wordsTrackerFile, trash)
	writeSentenceOnFile(sentence, wordsTrackerFile)
	fmt.Printf("after")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s \"sentence to add\"\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func writeSentenceOnFile(sentence, filePath string) {
	err := WriteFile(sentence, filePath)
	check(err)
	fmt.Printf("Sentence '%s' added successfully\n", sentence)
}
