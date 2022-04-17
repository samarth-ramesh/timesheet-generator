package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

const RootPath = "/tmp/timesheet"
const Start = false
const End = true

func main() {
	args := os.Args[1:]

	if (len(args) != 1) && (len(args) != 2) {
		panic("Expected  1 or 2 arguments \nUsage: " + os.Args[0] + " [add] <sheet_name>")
	}

	if len(args) == 2 {
		if args[0] == "add" {
			// need to create the sheet
			CreateSheet(args[1])
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

func CreateSheet(sheetName string) {
	_, err := os.Create(path.Join(RootPath, sheetName))
	if err != nil {
		panic(err)
	}
}

func MakeEntry(sheetName string, fileSize int64) {
	sheetPath := path.Join(RootPath, sheetName)
	lastEntry := GetLastEntryType(sheetPath)
	timestamp := GetDateTimeString()
	line := GenerateEntryString(timestamp, lastEntry, fileSize == 0)
	file, err := os.OpenFile(sheetPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(line)
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
}

func GetLastEntryType(sheetPath string) bool {
	raw, err := os.ReadFile(sheetPath)
	if err != nil {
		panic(err)
	}
	data := string(raw)
	lines := strings.Split(data, "\n")
	lastEntry := Start
	if lines[0] != "" {
		if len(lines)%2 == 0 { // i.e there are an even number of lines;
			// => there are equal number of ends for starts;
			// => the prev one was an end
			lastEntry = End
		} else {
			lastEntry = Start
		}
	} else {
		lastEntry = End // the first line must be a start. Hence the -1'th line can be treated as end.
	}
	return lastEntry
}

func GetDateTimeString() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04")
}

func GenerateEntryString(timeString string, lastEntryType bool, isFirstItem bool) string {
	endItem := ""
	if lastEntryType {
		endItem = "S"
	} else {
		endItem = "E"
	}
	startItem := "\n"
	if isFirstItem {
		startItem = ""
	}
	return fmt.Sprintf("%s %s %s", startItem, timeString, endItem)
}
