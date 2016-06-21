// package Triangle provides a function to classify triangles.
package triangle

import (
  "math"
)

const testVersion = 2

// Function KindFromSides classifies triangles from the lengths of their sides, reporting results of type Kind
func KindFromSides(a, b, c float64) Kind {

  // sort sides in order of ascending length
  x, y, z := sortSides(a, b, c)

  // sides are invalid if any length is zero or negative length,
  // or if two shortest sides are less than longest side
  if (math.IsNaN(0 * x * y * z) || x <= 0 || (x + y) < z) {
    return NaT
  }

	// triangle is equilateral if all sides are same length
	if a == b && b == c {
    return Equ
  }

  // triange is isosceles if two shortest or longest sides are same length
	if x == y || y == z {
		return Iso
	}

  // triangle is scalene if all sides are different lengths
	return Sca
}

// public type Kind is based on private type kind, so callers cannot create their own values
type Kind kind

type kind int

// callers can reference these Kind values because they're public and pre-defined
const (
	NaT Kind = iota // not a triangle
	Equ Kind = iota // equilateral
	Iso Kind = iota // isosceles
	Sca Kind = iota // scalene
)

// sorts side lengths in ascending order
func sortSides(a, b, c float64) (x, y, z float64) {

	// if they're in correct order
	if a <= b && b <= c {
		return a, b, c
	}

	//  if they're in simple reverse order
	if a >= b && b >= c {
		return c, b, a
	}

	// if b is first
	if b < a {
		if a < c {
			return b, a, c
		}
		return b, c, a
	}

	// so, b is last
	if a < c {
		return a, c, b
	}

	// only one option left...
	return c, a, b
}
