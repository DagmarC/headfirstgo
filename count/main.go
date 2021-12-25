// Counts the number of votes per person from a specified file.
package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/DagmarC/headfirstgo/constants"
	"github.com/DagmarC/headfirstgo/datafile"
)

func main() {
	var votes map[string]int

	votes, err := datafile.GetMaps(constants.VotesText)
	if err != nil {
		log.Fatal(err)
	}
	// Map is unordered collection, so to have it ordered follow:
	// 1. Create slice for holding the key values.
	// 2. Loop through tha map and store all keys in the slice.
	// 3. Order the slice.
	// 3. Loop throughout the slice of ordered keys and obtain the value via the key.

	// 1.
	var orderedVoteKeys []string
	// 2.
	for key := range votes {
		orderedVoteKeys = append(orderedVoteKeys, key)
	}
	// 3.
	sort.Strings(orderedVoteKeys)
	// 4.
	for _, orderedKey := range orderedVoteKeys {
		fmt.Printf("%s has %d votes.\n", orderedKey, votes[orderedKey])
	}
}
