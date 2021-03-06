package main

// declare as executable

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// import library packages

func main() { // main function of the exectuable
	csvFilename := flag.String("csv", "problems.csv", "a csv file: [question], [answer]")
	timeLimit := flag.Int("limit", 10, "the time limit in seconds")
	flag.Parse() // parse command-line
	// implementing command-line flags (flag name, default, description)

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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	// create a new timer that emits once

	correct := 0                 // variable for tracking correct answers
	for i, p := range problems { // iterates over elements, first return is index
		fmt.Printf("Problem #%d: %s = ", i+1, p.q) // output problem, i is incremented up to be human readable,
		answerCh := make(chan string)              // initialize channel
		go func() {                                // goroutine - lightweight thread of execution (synchronous)
			var answer string
			fmt.Scanf("%s\n", &answer)
			// ask for user input and save to answer variable
			answerCh <- answer
		}()

		select {
		case <-timer.C: // channel - a pipe to connect conncurent goroutines (asynchronous)
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems)) // output score
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
				// correct answers increment the correct variable
			}
		}
	}

	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
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
