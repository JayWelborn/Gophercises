package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("csv", "problems.csv", "Name of file to open")
	flag.Parse()

	csvFile, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Error opening file %s.\nClosing program", *fileName))
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Error reading lines")
	}
	problems := parseLines(lines)

	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct++
		}
	}

	fmt.Printf("%d out of %d correct\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			line[0],
			strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(message string) {
	fmt.Printf("%s\n", message)
	os.Exit(1)
}
