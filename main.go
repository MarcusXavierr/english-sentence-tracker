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

	wordsTrackerFile, trash := getFiles()
	sentence := os.Args[1]

	saveSentence(sentence, wordsTrackerFile, trash)
}

func saveSentence(sentence string, wordsTrackerFile string, trash string) {
	if !CheckIfSentenceExists(os.Stdout, sentence, wordsTrackerFile, trash) {
		writeSentenceOnFile(sentence, wordsTrackerFile)
	}
}

func getFiles() (string, string) {
	wordsTrackerFile := getHomeDir() + "/english_words/words.txt"
	trash := getHomeDir() + "/english_words/trash.txt"
	return wordsTrackerFile, trash
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func usage() {
	var message string = fmt.Sprintf("usage: %s \"sentence to add\"\n", os.Args[0])
	PrintRed(os.Stdout, message)
	flag.PrintDefaults()
	os.Exit(2)
}

func writeSentenceOnFile(sentence, filePath string) {
	err := WriteFile(sentence, filePath)
	check(err)
	message := fmt.Sprintf("Sentence %q added successfully\n", sentence)
	PrintGreen(os.Stdout, message)
}
