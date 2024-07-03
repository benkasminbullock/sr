// Special relativity
package sr

import (
	"fmt"
	"math"
)

// The speed of light in metres per second
const C = 299792458.0

var Debug = false

// Given a speed in metres per second, return the beta value (speed
// divided by speed of light) corresponding to it.
func Beta(v float64) (beta float64) {

}

// Given a speed in metres per second, return the corresponding gamma
// value, sqrt(1/(1-(v/c)^2)).
func Gamma(v float64) (g float64) {
	beta := v / C
	if Debug {
		fmt.Printf("Beta = %g\n", beta)
	}
	g = math.Sqrt(1 / (1 - beta*beta))
	if Debug {
		fmt.Printf("Gamma = %g\n", g)
	}
	return g
}

// Given a velocity in metres per second, return the corresponding K
// factor. Negative velocities are approaching, and positive
// velocities are receding.
func K(v float64) (k float64) {

}
