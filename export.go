package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

func ExportSheet(sheetName string) {
	sheetPath := path.Join(RootPath, sheetName)

	raw, err := os.ReadFile(sheetPath)
	if err != nil {
		panic(err)
	}
	if len(raw) == 0 {
		panic("Empty File")
	}

	lines := strings.Split(string(raw), "\n")
	if len(lines)%2 != 0 {
		panic("Unterminated timesheet.\nTerminate by running " + os.Args[0] + " " + sheetPath + " end")
	}

	t, err := time.Parse(TimeFormat, lines[0][:len(lines[0])-2])
	if err != nil {
		panic(err)
	}

	total := time.Duration(0)

	println("\\Large\\textbf{Time Sheet for " + t.Format("January") + " 2022}\\normalsize \\\\\n" +
		"\\begin{tabular}{|l|l|l|}\n" +
		"\\hline\n" +
		"\t&  & \\\\\n" +
		"\t\\textbf{Date} & \\textbf{Start} & \\textbf{End} \\\\\n" +
		"\t& & \\\\\n" +
		"\\hline")
	for i := 0; i < len(lines); i += 2 {
		l1, l2 := lines[i], lines[i+1]
		l1, l2 = l1[:len(l1)-2], l2[:len(l2)-2]

		t1, err := time.Parse(TimeFormat, l1)
		if err != nil {
			panic(err)
		}
		t2, err := time.Parse(TimeFormat, l2)
		if err != nil {
			panic(err)
		}
		d := t2.Sub(t1)
		total = total + d
		printLatexRow(t1, t2)
	}
	printTotalTime(total)
	println("\\end{tabular}")
}

func printLatexRow(start time.Time, end time.Time) {
	println("\t& & \\\\")
	fmt.Printf("\t%s & %s & %s \\\\\n", start.Format("2006-01-02"), start.Format("15:04"), end.Format("15:04"))
	println("\t& & \\\\")
	println("\\hline")
}

func printTotalTime(total time.Duration) {
	fmt.Printf(" & \\textbf{Total} & %s \\\\ \\hline\n", total.String())
}
