package main

import (
	"os"
	"path"
)

func CreateSheet(sheetName string) {
	_, err := os.Create(path.Join(RootPath, sheetName))
	if err != nil {
		panic(err)
	}
}
