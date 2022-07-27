package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		usage()
	}

	wordsTrackerFile := getHomeDir() + "/english_words/words.txt"
	trash := getHomeDir() + "/english_words/trash.txt"
	sentence := os.Args[1]

	CheckIfSentenceExists(sentence, wordsTrackerFile, trash)
	writeSentenceOnFile(sentence, wordsTrackerFile)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func usage() {
	var message string = fmt.Sprintf("usage: %s \"sentence to add\"\n", os.Args[0])
	color.Red(message)
	flag.PrintDefaults()
	os.Exit(2)
}

func writeSentenceOnFile(sentence, filePath string) {
	err := WriteFile(sentence, filePath)
	check(err)
	message := fmt.Sprintf("Sentence '%s' added successfully\n", sentence)
	color.Green(message)
}
