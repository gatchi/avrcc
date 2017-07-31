package main

import "fmt"
import "os"

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
	_ = file
	fmt.Printf("Closed %s.\n", inputFileName)
}
