package gadgeterror

import "fmt"

type GadgetError float64

// Can be passed or returned as an error -> satisfied.
// Customized errors eg display value.
func (g GadgetError) Error() string {
	return fmt.Sprintf("Limit of songs exceeded by %0.2f", g)
}
