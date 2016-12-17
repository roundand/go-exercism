// Package school provides functions to operate a school grades DB.
package school

import "sort"

// School contains student names, listed by grade.
type School map[int][]string

// New returns an empty School.
func New() *School {
	s := School(map[int][]string{})
	return &s
}

// Grade is maps a grade level to a collection of student names.
type Grade struct {
	grade int
	names []string
}

// Enrollment returns all grades for the school.
func (s *School) Enrollment() (grades []Grade) {

	// get ordered copy of school grades, since map keys are guaranteed unordered.
	keys := []int{}
	for key := range *s {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// also sort the names within each grade, since they were added in random order
	for _, key := range keys {
		sort.Strings((*s)[key])
		grades = append(grades, Grade{key, (*s)[key]})
	}

	return grades
}

// Grade returns student names for the given grade.
func (s *School) Grade(grade int) (names []string) {
	return (*s)[grade]
}

// Add adds a grade and name to the school record.
func (s *School) Add(name string, grade int) *School {
	(*s)[grade] = append((*s)[grade], name)
	return s
}
