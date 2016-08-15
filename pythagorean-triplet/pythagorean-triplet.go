// Package pythagorean hosts two functions for
// producing pythagorean triplets, Range and Sum
package pythagorean

// Triplet contains the three integer members of a Pythagorean triplet.
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
  var triplets []Triplet
  for x := min; x <= max; x++ {
    x2 := x * x
    for y := x; y <= max; y++ {
      y2 := y * y
      for z := y; z <= max; z++ {
        z2 := z * z

        // gotcha - capture the triplet and move on
        if x2 + y2 == z2 {
          triplets = append(triplets, Triplet{x, y, z})
          break
        }

        // no triplet exists for this x and y
        if z2 > x2 + y2 {
          break
        }
      }
    }
  }
  return triplets
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
  var triplets []Triplet

  // no triangle can have a side which is longer than half its perimeter
  for _, t := range(Range(1, p / 2)) {
    if (t[0] + t[1] + t[2]) == p {
      triplets = append(triplets, t)
    }
  }
  return triplets
}
