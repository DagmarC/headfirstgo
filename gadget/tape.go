package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Paying, ", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped.")
}

// Satisfies the Stringer interface - TapePlayer can be passed to fmt.Println(), ... and will be printed.
func (t TapePlayer) String() string {
	return fmt.Sprint("Tape player string method")
}

type TapeRecorder struct {
	Microphones string
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Paying, ", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped.")
}

func (t TapeRecorder) Record() {
	fmt.Println("Record.")
}
