package eea

import (
	"testing"
)

func TestCompute(t *testing.T) {
	expected := []state{
		state{3, 4, 46, 10, 0, 1, 1, -5},
		state{4, 1, 10, 6, 1, -4, -5, 21},
		state{5, 1, 6, 4, -4, 5, 21, -26},
	}
	s := newState(240, 46)
	for i, x := range expected {
		if s.compute(); x != *s {
			t.Errorf("Computing failed on iterion with %d with %v, expected %v", i+2, *s, x)
		}
	}
}

func TestGCD(t *testing.T) {
	a, b := 240, 46
	gcd, ca, cb, qa, qb := GCD(a, b)
	switch {
	case gcd != 2:
		t.Errorf("Computing gcd failed with %d", gcd)
	case gcd != ca*a+cb*b:
		t.Errorf("Computing Bezout coeffiecients failed with %d and %d", ca, cb)
	case qa != abs(a/gcd):
		t.Errorf("Computing quotient of a failed with %d", qa)
	case qb != abs(b/gcd):
		t.Errorf("Computing quotient of b failed with %d", qb)
	}
}
