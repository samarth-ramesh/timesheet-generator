package main

import (
	"os"
	"path"
)

func CreateSheet(sheetName string) {
	pathname := path.Join(RootPath, sheetName)
	err := os.MkdirAll(RootPath, os.ModeDir)
	if err != nil {
		panic(err)
	}
	_, err = os.Create(pathname)
	if err != nil {
		panic(err)
	}
}
