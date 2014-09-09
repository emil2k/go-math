package eea

import (
	"fmt"
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

func TestMMI(t *testing.T) {
	a, n := 21, 5
	if x, err := MMI(a, n); err != nil {
		t.Errorf("Computing mmi failed with error %s", err)
	} else if i := x * a % n; i != 1 {
		t.Errorf("Commputing mmi failed when inversing with %d", i)
	}
}

func TestMMINoInverse(t *testing.T) {
	if _, err := MMI(2, 2); err != NoInverse {
		t.Errorf("Computing mmi with no inverse failed with wrong error %s", err)
	}
}

func ExampleGCD() {
	a, b := 240, 46
	gcd, ca, cb, qa, qb := GCD(a, b)
	fmt.Printf("Greatest common divisor : %d\n", gcd)
	fmt.Printf("Bezout's coeffecients : %d (%d) & %d (%d)\n", ca, a, cb, b)
	fmt.Printf("GCD quotients absolute values : %d (%d) & %d (%d)\n", qa, a, qb, b)
	// Output:
	// Greatest common divisor : 2
	// Bezout's coeffecients : -9 (240) & 47 (46)
	// GCD quotients absolute values : 120 (240) & 23 (46)
}
