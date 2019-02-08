package k8slate

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Write - writes a rendered document to disk at dirPath
func Write(doc RenderedDocument, path string) {
	err := ioutil.WriteFile(path, []byte(doc.Result), 0644)

	if err != nil {
		fmt.Println("Error writing file at ", path, " ", err)
		os.Exit(-1)
	}
}

// Mkdirp - creates a nested directory structure ending at path
func Mkdirp(path string) {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		fmt.Println("Error creating output folder at ", path, " ", err)
		os.Exit(-1)
	}
}
