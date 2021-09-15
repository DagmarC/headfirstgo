// GetMaps reads from file and stores it into the map where key is the name and value is the number of occurrences in the file.
// Example -> key: Richard Cervenka, value: 3
package datafile

import (
	"bufio"
	"os"
)

func GetMaps(filename string) (map[string]int, error) {
	var votes map[string]int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	votes = make(map[string]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		votes[line]++
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return votes, nil
}
