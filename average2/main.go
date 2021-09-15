// The average calculates the aveage of the command line parameters - floats64.
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/constants"
	"github.com/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFLoats(constants.AverageFloatsText)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numbers)
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average is %0.2f\n", sum/sampleCount)
}
