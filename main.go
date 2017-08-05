package main

import (
	"bufio"
	"fmt"
	"io"
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
	line string
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

	/*---------------- Open input file -----------------*/

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

	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)

	// Read until no more file to read
	j := 1
	for reader.Scan() {
		line = reader.Text()
		println("Line ", j)
		j++
		println("-------")
		println(line)

		sReader := strings.NewReader(line)

		// Crop comment, if any
		newEnd := 0
		for c, _ := sReader.ReadByte(); c != ';'; newEnd++ {
			c, err = sReader.ReadByte()
			if err == io.EOF {
				break
			}
		}
		cLine := line[:newEnd]
		sReader.Reset(cLine)

		// Get mnemonic
		mnemonic = ""
		for c, _ := sReader.ReadByte(); c != ' '; {
			mnemonic = mnemonic + string(c)
			c, err = sReader.ReadByte()
			if err == io.EOF {
				break
			}
		}
		println(mnemonic)

		// Check if mnemonic is valid; if it isn't, error and exit
		if mnemonicList[mnemonic] == "" {
			fmt.Printf("Error on line %v: Line doesn't start with valid mnemonic\n", 1)
			logTokenError(mnemonic)
			alertAndClose(file)
			os.Exit(mnemonicError)
		}

		// Find the number of operands we're expecting
		opNum := mnemonicOpList[mnemonic]
		println(opNum)

		// Extract operands until hitting end of line
		opIndex := 0
		for c, err := sReader.ReadByte(); err != io.EOF && opIndex < opNum; {
			for ; c != ','; {
				ops[opIndex] = ops[opIndex] + string(c)
				c, err = sReader.ReadByte()
			}
			opIndex++
			println(ops[opIndex-1])
			c, err = sReader.ReadByte()
		}
	}
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
