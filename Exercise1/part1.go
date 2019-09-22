package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func poseQuestion(questionNo int, question string, answer string) bool {
	fmt.Printf("Problem #%d: %s\n", questionNo, question)
	var submitted string
	fmt.Scan(&submitted)
	submitted = strings.TrimSpace(submitted)
	return submitted == answer
}

func main() {
	fileName := flag.String("csv", "problems.csv", "Name of file to open")
	flag.Parse()

	csvFile, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Error opening file %s.\nClosing program", *fileName)
		os.Exit(1)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	current, correct := 0, 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading line. Moving on.")
			break
		}
		current += 1
		if poseQuestion(current, line[0], line[1]) {
			correct += 1
		}
	}
	fmt.Printf("%d out of %d correct\n", correct, current)
}
