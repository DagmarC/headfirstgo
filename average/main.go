// The average calculates the aveage of the command line parameters - floats64.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var numbers []float64

	for _, number := range arguments {
		number, err := strconv.ParseFloat(number, 64)
		if err != nil {
			log.Fatal("Could not convert the number: ", err)
		}
		numbers = append(numbers, number)
	}

	fmt.Printf("Average of all numbers is %0.2f. ", average(numbers...))
}

func average(numbers ...float64) float64 {
	var sum float64 = 0

	for _, n := range numbers {
		sum += n
	}
	return sum / float64(len(numbers))
}
