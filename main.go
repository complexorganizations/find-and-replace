package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var (
	filePath  = "/home/example/example.json"
	oldString = "old"
	newString = "new"
)

func main() {
	// lets read the file
	read, err := ioutil.ReadFile(filePath)

	// if there is a error reading the file than we show the error using the log package
	if err != nil {
		log.Fatalln(err)
	}

	// replace the current old string with the new string
	newContents := strings.Replace(string(read), (oldString), (newString), -1)

	// write the new string
	err = ioutil.WriteFile(filePath, []byte(newContents), 0)

	// if there is a error print it
	if err != nil {
		log.Fatalln(err)
	}
}
