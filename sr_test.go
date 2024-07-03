package sr

import "testing"

func TestGamma(t *testing.T) {
	gamma := Gamma(C * 0.6)
	if gamma != 1/0.8 {
		t.Errorf("%g should be %g", gamma, 0.8)
	}
}
