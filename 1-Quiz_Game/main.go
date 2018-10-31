package main

// declare as executable

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// import library packages

func main() { // main function of the exectuable
	csvFilename := flag.String("csv", "problems.csv", "a csv file: [question], [answer]")
	// "short" declaration implementing command-line flags
	flag.Parse()

	file, err := os.Open(*csvFilename)
	// "short" declaration opening a file using OS-like functionality
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)  // reading from a CSV file
	lines, err := r.ReadAll() // reads each record is a slice of fields
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines) // call function written below

	correct := 0                 // variable for tracking correct answers
	for i, p := range problems { // iterates over elements, first return is index
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q) // output problem, i is incremented up to be human readable,
		var answer string                            // variable to hold the incoming answer
		fmt.Scanf("%s\n", &answer)                   // ask for user input and save to answer variable
		if answer == p.a {
			// correct answers increment the correct variable
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
	// output score
}

func parseLines(lines [][]string) []problem { // parameter: lines slice of type string, output: problem struct
	ret := make([]problem, len(lines)) // allocates and initializes an object, type is problem struct, size is # of lines
	for i, line := range lines {       // iterates over elements, first return is index
		ret[i] = problem{ // set return
			q: line[0],                    // because of our format this is the question
			a: strings.TrimSpace(line[1]), // because of our format this is the answer
		}
	}
	return ret
}

// declare problem struct
type problem struct {
	q string
	a string
}

// exit the execution with a message
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
