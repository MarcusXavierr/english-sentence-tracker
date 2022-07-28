package main

import (
	"fmt"
	"os"
	"strings"
)

func CheckIfSentenceExists(sentence string, filePathList ...string) {
	for index := range filePathList {
		if hasSentence(sentence, filePathList[index]) {
			fmt.Fprintf(os.Stderr, "sentence '%s' already exists on file %s\n", sentence, filePathList[index])
			os.Exit(3)
		}
	}
}

func hasSentence(sentence, filePath string) bool {
	buffer := ReadFile(filePath)
	return verifyIsSentenceExists(sentence, buffer)
}

func verifyIsSentenceExists(sentence, buffer string) bool {
	var allLines []string = strings.Split(buffer, "\n")
	for line := range allLines {
		if compareStrings(sentence, allLines[line]) {
			return true
		}
	}
	return false
}

//Get one line from my file and compare to a sentence
func compareStrings(sentence, line string) bool {
	sentence = strings.ToLower(sentence)
	line = strings.ToLower(line)
	line = strings.Trim(line, " ")
	return strings.Compare(line, sentence) == 0
}
