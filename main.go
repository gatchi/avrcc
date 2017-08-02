package main

import (
	"bufio"
	"fmt"
	"os"
)

/// Exit codes
const (
	properExit int = 0
	noFileName int = 1
	mnemonicError int = 2
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

	// Process input file

	// Using reader
	fileReader := bufio.NewReader(file)
	var mnemonic string

	// Read first token, which should be a mnemonic followed by a space
	token, err := fileReader.ReadString(' ')
	//fmt.Println("First token: " + token)
	if err != nil {
		fmt.Print("Something happened: ")
		fmt.Println(err)
	}
	if mnemonicList[token[:len(token)-1]] != "" {
		mnemonic = token
		fmt.Println("Mnemonic: " + token)
	} else {
		fmt.Printf("Error on line %v: Line doesn't start with valid mnemonic\n", 1)
		logTokenError(token)
		alertAndClose(file)
		os.Exit(mnemonicError)
	}
	_ = mnemonic

	/*// Using scanner:
	scanner := bufio.NewScanner(file)
	scanner.Split(util.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading input:", err)
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
