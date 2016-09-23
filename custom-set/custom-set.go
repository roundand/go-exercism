// Package stringset implements Set as a collection of unique string values.
package stringset

import (
	"fmt"
)

const testVersion = 3

// Set implements set operations
type Set map[string]bool

// New returns a Set.
func New() Set {
	ms := Set{}
	return ms
}

// NewFromSlice returns a Set.
func NewFromSlice(sl []string) Set {
	set := New()
	for _, s := range sl {
		set.Add(s)
	}
	return set
}

// Add modifies the target Set.
func (set Set) Add(a string) {
	m := map[string]bool(set)
	m[a] = true
}

// Delete modifies the target Set.
func (set Set) Delete(a string) {
	m := map[string]bool(set)
	delete(m, a)
}

// Has reports if a Set has a value.
func (set Set) Has(a string) bool {
	m := map[string]bool(set)
	return (m[a])
}

// IsEmpty reports if a set is empty.
func (set Set) IsEmpty() bool {
	return (set.Len() == 0)
}

// Len reports the length of a Set.
func (set Set) Len() int {
	m := map[string]bool(set)
	return (len(m))
}

// Slice returns he contents of a Set as a slice of strings.
func (set Set) Slice() []string {
	ss := []string{}
	for s := range set {
		ss = append(ss, s)
	}
	return (ss)
}

// String returns the content of the Set as an approprately formatted string.
func (set Set) String() string {
	res := ""
	for s := range set {
		if res == "" {
			res = fmt.Sprintf("%q", s)
		} else {
			res += ", " + fmt.Sprintf("%q", s)
		}
	}
	return "{" + res + "}"
}

// Equal reports whether two Sets have identical members.
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

// Subset returns s1 ⊆ s2.
func Subset(s1, s2 Set) bool {
	for s := range s1 {
		if !s2.Has(s) {
			return false
		}
	}
	return true
}

// Disjoint reports whether s1 and s2 have no common members.
func Disjoint(s1, s2 Set) bool {
	// let's be more efficient than return !Intersection(s1, s2).IsEmpty()
	for s := range s1 {
		if s2.Has(s) {
			return false
		}
	}
	return true
}

// Intersection returns the Set of members common to s1 and s2.
func Intersection(s1, s2 Set) Set {
	s3 := New()
	for s := range s1 {
		if s2.Has(s) {
			s3.Add(s)
		}
	}
	return s3
}

// Union returns the Set of members of s1 and / or s2.
func Union(s1, s2 Set) Set {
	s3 := New()
	for s := range s1 {
		s3.Add(s)
	}
	for s := range s2 {
		s3.Add(s)
	}
	return s3
}

// Difference returns all members of s1 which are not in s2, ie s1 \ s2.
func Difference(s1, s2 Set) Set {
	s3 := New()
	for s := range s1 {
		if !s2.Has(s) {
			s3.Add(s)
		}
	}
	return s3
}

// SymmetricDifference returns the Set of all members of either s1 or s2 but not both.
func SymmetricDifference(s1, s2 Set) Set {
	return Union(Difference(s1, s2), Difference(s2, s1))
}
