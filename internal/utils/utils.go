package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ListFiles() {
	dir := "."
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.Mode().IsRegular() && strings.HasSuffix(file.Name(), ".mp4") {
			fmt.Println(filepath.Join(dir, file.Name()))
		}
	}
}
