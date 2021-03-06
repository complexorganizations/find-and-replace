package main

import (
	"io/ioutil"
	"log"
	"strings"
)

const (
	// the location of the file.
	filePath = "./example.json"
)

var (
	// the string your trying to change.
	oldString = "THIS IS THE OLD STRING."
	// the new string you want to replace it with.
	newString = "THIS IS THE NEW STRING."
)

func main() {
	// lets read the file
	read, err := ioutil.ReadFile(filePath)
	// replace the current old string with the new string
	newContents := strings.Replace(string(read), (oldString), (newString), -1)
	// write the new string
	err = ioutil.WriteFile(filePath, []byte(newContents), 0)
	// if there is a error print it
	if err != nil {
		log.Println(err)
	}
}
