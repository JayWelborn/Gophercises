package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("csv", "problems.csv", "Name of file to open")
	timeOut := flag.Int("limit", 30, "Seconds per answer before timeout")
	flag.Parse()

	csvFile, err := os.Open(*fileName)
	if err != nil {
		exit(
			fmt.Sprintf("Error opening file %s.\nClosing program", *fileName),
			1)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Error reading lines", 1)
	}

	fmt.Println("Press enter when ready")
	var ready string
	fmt.Scanln(&ready)

	problems := parseLines(lines)
	correct := 0
	answerChannel := make(chan bool, 1)

	for i, problem := range problems {
		go getAnswer(answerChannel, i, problem)
		select {
		case res := <-answerChannel:
			if res {
				correct++
			}
		case <-time.After(time.Duration(*timeOut) * time.Second):
			exit(fmt.Sprintf(
				"Time limit exceeded.\n%d out of %d correct",
				correct,
				len(problems)),
				0)
		}
	}

	exit(fmt.Sprintf("%d out of %d correct", correct, len(problems)), 0)
}

func getAnswer(ret chan<- bool, i int, p problem) {
	fmt.Printf("Problem #%d: %s = ", i+1, p.question)
	var answer string
	fmt.Scanf("%s\n", &answer)
	ret <- answer == p.answer
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

func exit(message string, code int) {
	fmt.Printf("%s\n", message)
	os.Exit(code)
}
