package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

func main() {

	level := flag.String("level", "CRITICAL", "Log level to filter for")
	flag.Parse()

	f, err := os.Open("./logs.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()
	// Open the file for reading
	// Check for errors
	// Defer the closing of the file

	bufReader := bufio.NewReader(f)

	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		// Read each line from the file
		if strings.Contains(line, *level) {
			// Check if the line contains the specified log level
			println(line)
			// Print the line to the console
		}
	}

}
