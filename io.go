package main

import (
	"os"
	"os/user"
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
