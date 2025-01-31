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
	beta = v / C
	if Debug {
		fmt.Printf("Beta = %g\n", beta)
	}
	// Check |beta| <= 1 ?
	return beta
}

// Given a speed in metres per second, return the corresponding gamma
// value, sqrt(1/(1-(v/c)^2)).
func Gamma(v float64) (g float64) {
	beta := Beta(v)
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
	beta := Beta(v)
	k = math.Sqrt((1 + beta) / (1 - beta))
	return k
}

// Given a K value, return the corresponding velocity
func KToV(k float64) (v float64) {
	k2 := k * k
	v = (k2 - 1) / (k2 + 1)
	return v
}
