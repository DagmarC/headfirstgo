// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"log"
	"os"
)

// GetsStrings reads a string from each line of a file.
func GetStrings(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, line)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
