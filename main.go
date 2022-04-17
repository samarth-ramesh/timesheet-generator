package main

import (
	"os"
	"path"
)

const RootPath = "/home/srinivas/Documents/work/timesheet"
const Start = false
const End = true

func main() {
	args := os.Args[1:]

	if (len(args) != 1) && (len(args) != 2) {
		panic("Expected  1 or 2 arguments \nUsage: " + os.Args[0] + " [add|export] <sheet_name>")
	}

	if len(args) == 2 {
		if args[0] == "add" {
			// need to create the sheet
			CreateSheet(args[1])
			return
		} else if args[0] == "export" {
			ExportSheet(args[1])
			return
		} else {
			panic("Unknown parameter " + args[0])
		}
	}

	stat, err := os.Stat(path.Join(RootPath, args[0]))
	if os.IsNotExist(err) {
		panic("sheet does not exist.\nCreate sheet with " + os.Args[0] + " add <sheet_name>")
	}
	s := stat.Size()
	MakeEntry(args[0], s)
}
