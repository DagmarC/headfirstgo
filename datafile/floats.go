// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// GetFloats reads a float values from each line of a file.
func GetFLoats(filename string) ([]float64, error) {

	var numbers []float64

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	// Look for Scanner Error.
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}
