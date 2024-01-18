package utils

import (
	"fmt"
	"io/ioutil"
)

func ListFiles() {
	// Example function to list files
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
