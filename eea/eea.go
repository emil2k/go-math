// Package eea is an implementation of the Extended Euclidean algorithm.
package eea

type state struct {
	i  int // iteration number
	q  int // reprsents quotient of r0/r1
	r0 int // represents r from i-2
	r1 int // represents r from i-1
	s0 int // represents s from i-2
	s1 int // represents s from i-1
	t0 int // repretents t from i-2
	t1 int // repretents t from i-1
}

// newState returns the initial state for computing the running the algorithm on the passed a & b.
func newState(a, b int) *state {
	return &state{2, a / b, a, b, 1, 0, 0, 1}
}

// compute computes one iteration of the algororithm updating the state.
func (s *state) compute() {
	s.r1, s.r0 = s.r0-s.q*s.r1, s.r1
	s.s1, s.s0 = s.s0-s.q*s.s1, s.s1
	s.t1, s.t0 = s.t0-s.q*s.t1, s.t1
	s.i++
	if s.r1 != 0 { // otherwise done computing
		s.q = s.r0 / s.r1
	}
}

// GCD uses the Extended Euclidean algorithm to compute the gcd of the passed a & b.
// Along with the gcd it returns the their Bezout coefficient and absolute value of the
// quotients with gcd.
func GCD(a, b int) (gcd, ca, cb, qa, qb int) {
	s := newState(a, b)
	for s.r1 != 0 {
		s.compute()
	}
	return s.r0, s.s0, s.t0, abs(s.t1), abs(s.s1)
}

// abs returns the absolute value for an int
func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
