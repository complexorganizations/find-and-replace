package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var (
	// the location of the file.
	filePath  = "./example.json"
	// the string your trying to change.
	oldString = "This is the old string we are trying to change."
	// the new string you want to replace it with.
	newString = "This is the string the old one is going to replace."
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
