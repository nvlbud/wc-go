// Count lines in file like "wc -l /path/to/file"

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// Count lines in file
// If file does not contain "\n" character returns 0 as original wc command
func countLines(filename string) int {
	var line_separator []byte = []byte("\n")
	// Total lines in file
	var lines int = 0

	// Open file in read mode
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Read file byte-by-byte
	data := make([]byte, 1)
	for {
		_, err := file.Read(data)

		if err != nil {
			// EOF is not an error, stop counting
			if err == io.EOF {
				break
			}
			// Something went wrong, print log and exit
			log.Fatal(err)
		} else {
			// Inc lines count if line separator found
			if bytes.Equal(data, line_separator) {
				lines++
			}
		}
	}

	return lines
}

func main() {
	// If no arguments passsed print usage info
	if len(os.Args) == 1 {
		fmt.Printf("Count number of lines\nUsage: %s filename\n", os.Args[0])
		return
	}

	// Allow only one file
	if len(os.Args) != 2 {
		log.Fatal("Incorrect usage")
	}

	fmt.Printf("%d %s\n", countLines(os.Args[1]), os.Args[1])
}
