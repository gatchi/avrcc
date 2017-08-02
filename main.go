package main

import (
	"bufio"
	"fmt"
	util "github.com/gatchi/utilgo"
	"os"
)

/// Exit codes
const (
	properExit int = 0
	noFileName int = 1
)

func main() {
	// Check if name of file is supplied.
	// If not, error and ask for one.
	// If one is, use it to load the file.
	if len(os.Args) < 2 {
		fmt.Println("Please supply name of file to assemble.")
		os.Exit(noFileName)
	}
	inputFileName := os.Args[1]

	// Open assembly file
	file, err := os.Open(inputFileName)
	defer file.Close()
	if os.IsNotExist(err) {
		fmt.Println("File does not exist by that name.")
		os.Exit(1)
	} else if err != nil {
		fmt.Println("Could not open file: " + err.Error())
		os.Exit(1)
	} else {
		fmt.Printf("Opened %s for reading.\n", inputFileName)
	}

	// Process input file

	/* // Using reader
	fileReader := bufio.NewReader(file)
	token, err := fileReader.ReadString(' ')
	fmt.Println("First token: " + token)
	if err != nil {
		fmt.Print("Something happened: ")
		fmt.Println(err)
	}
	*/

	// Using scanner:
	scanner := bufio.NewScanner(file)
	scanner.Split(util.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading input:", err)
	}

	fmt.Printf("Closed %s.\n", inputFileName)
}
