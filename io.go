package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

func ReadFile(filePath string) string {
	buffer, err := os.ReadFile(filePath)
	check(err)
	return string(buffer)
}

func WriteFile(sentence, filepath string) error {
	file, e := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(644))
	check(e)
	_, err := file.WriteString(sentence + "\n")
	check(err)
	return file.Close()
}

func getHomeDir() string {
	usr, err := user.Current()
	check(err)
	return usr.HomeDir
}

func PrintRed(out io.Writer, message string) {
	fmt.Fprintln(out, string(colorRed), message, string(colorReset))
}

func PrintGreen(out io.Writer, message string) {
	fmt.Fprintln(out, string(colorGreen), message, string(colorReset))
}
