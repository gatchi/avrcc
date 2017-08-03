package main

import (
	"bufio"
	"fmt"
	//"io"
	"os"
	"strings"
)

/// Exit codes
const (
	properExit int = 0
	noFileName int = 1
	mnemonicError int = 2
)

var (
	mnemonic string
	ops [5]string
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
	defer alertAndClose(file)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist by that name.")
		alertAndClose(file)
		os.Exit(1)
	} else if err != nil {
		fmt.Println("Could not open file: " + err.Error())
		alertAndClose(file)
		os.Exit(1)
	} else {
		fmt.Printf("Opened %s for reading.\n", inputFileName)
	}

	/*------------ Process input file ------------------*/

	// Using scanner cause we goin by lines
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	line := fileScanner.Text()

	/* Mnemonic */

	// Need a string reader for this
	lineReader := strings.NewReader(line)

	// Read first token, which should be a mnemonic followed by a space
	for {
		letter, _, _ := lineReader.ReadRune()
		if letter == ' ' {
			break
		}
		mnemonic = mnemonic + string(letter)
	}
	fmt.Println(mnemonic)

	// Check if mnemonic is valid; if it isn't, error and exit
	if mnemonicList[mnemonic] == "" {
		fmt.Printf("Error on line %v: Line doesn't start with valid mnemonic\n", 1)
		logTokenError(mnemonic)
		alertAndClose(file)
		os.Exit(mnemonicError)
	}

	/* Operands */

	/*
	// Find the number of operands we're expecting
	opNum := mnemonicOpList[mnemonic]
	fmt.Printf("Numeber of ops: %d\n", opNum)

	// Extract operands until hitting newline or a comment
	for i := 0; i < opNum; i++ {
		// If this is the last op, expect a semicolon or newline at the end
		token, err := fileReader.ReadString(',')

		// Catch unexpected error
		if err != nil && err != io.EOF {
			fmt.Print("Something happened: ")
			fmt.Println(err)
			break
		}

		// Print token
		fmt.Printf("OP #%d: %s\n", i, strings.Trim(token, ", "))

		// In case less text than expected
		if err == io.EOF {
			break
		}
	}
	*/
}

func alertAndClose(file *os.File) {
	file.Close()
	fmt.Println("File closed.")
}

func logTokenError(token string) {
		fmt.Println("Token read: " + token)
		fmt.Printf("Length of token: %v\n", len(token))
		fmt.Printf("Last token value: %v\n", token[len(token)-1])
		fmt.Println("Mapped value: " + mnemonicList[token])
}
