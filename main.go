package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	fileDir   = "/home/example/"
	filePath  = "/home/example/example.json"
	oldString = "old"
	newString = "new"
)

func visit(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	matched, err := os.Open(fileDir)

	if err != nil {
		panic(err)
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		//fmt.Println(string(read))
		fmt.Println(path)
		// change it here.
		newContents := strings.Replace(string(read), (oldString), (newString), -1)

		// fmt.Println(newContents)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

	}

	return nil
}

func main() {
	err := filepath.Walk(".", visit)
	if err != nil {
		panic(err)
	}
}
