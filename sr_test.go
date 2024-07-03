package sr

import (
	"math"
	"testing"
)

func TestGamma(t *testing.T) {
	gamma := Gamma(C * 0.6)
	if !close(gamma, 1/0.8) {
		t.Errorf("%g should be %g", gamma, 1/0.8)
	}
}

// Allowable floating point error.
var eps = 1e-10

func close(a, b float64) bool {
	gap := 2 * math.Abs((a-b)/(a+b))
	if gap > eps {
		return false
	}
	return true
}

func TestK(t *testing.T) {
	k := K(C * 0.5)
	if !close(k, math.Sqrt(3)) {
		t.Errorf("K value %g should be %g", k, math.Sqrt(3))
	}
}
func TestClose(t *testing.T) {
	if close(10, 10*(1+2*eps)) {
		t.Errorf("Failed close test 1")
	}
	if !close(10, 10*(1+eps/2)) {
		t.Errorf("Failed close test 2")
	}
}
