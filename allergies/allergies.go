// package allergies provides functions for allergy lists and allergy scores.
package allergies

// Allergies master list
var allergies = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func isFlagged(score, flag int) bool {
	f := score & (1 << uint(flag))
	return f != 0
}

// Allergies accepts an int allergies score and interprets it as a set of flags, returning the corresponding allergies as []string.
func Allergies(score int) []string {
	var r []string
	for i := 0; i < len(allergies); i++ {
		if isFlagged(score, i) {
			r = append(r, allergies[i])
		}
	}
	return r
}

// AllergicTo determines if someone with a supplied allergies score is allergic to a supplied allergy.
func AllergicTo(score int, allergy string) bool {

	for i := 0; i < len(allergies); i++ {
		if isFlagged(score, i) && (allergies[i] == allergy) {
			return true
		}
	}

	return false
}
