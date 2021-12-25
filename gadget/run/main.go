package main

import (
	"fmt"
	"github.com/DagmarC/headfirstgo/gadget"
	"github.com/DagmarC/headfirstgo/gadget/gadgeterror"
	"log"
)

type Player interface {
	Play(song string)
	Stop()
}

// Any error type that satisfies the error can be returned.
func playList(player Player, songs []string) error {
	if len(songs) > 1 {
		return gadgeterror.GadgetError(len(songs) - 6)
	}
	// Assertion types - If type is of a certain type you can call other methods than defined in the interface.
	// ok value signals if the assertion type was successful.
	for _, song := range songs {
		player.Play(song)
	}
	player.Stop()

	// TYPE ASSERTION
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record() // This method is only in the TapeRecorder
	}

	tape, ok := player.(gadget.TapePlayer)
	if ok {
		fmt.Println(tape) // Stringer interface is satisfied.
	}
	return nil
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie J"}
	err := playList(player, mixtape)
	if err != nil {
		log.Fatal(err)
	}
}
