// Package brackets provides a function to match various opening and closing brackets
package brackets

const testVersion = 3

// this stack has its top at the start.
type Stack string

func push(s Stack, b byte) Stack {
	return Stack(b) + s
}

func pop(s Stack) (Stack, byte) {
	return s[1:], s[0]
}

// Bracket accepts a string and reports whether it contains matching braces, brackets and parentheses.
func Bracket(input string) (matched bool, err error) {
	var stack Stack
	var top byte

	// read each character from the input
	for i := 0; i < len(input); i++ {
		switch in := input[i]; in {
		case '{', '[', '(': // This is an opening symbol - push it on to the stack
			stack = push(stack, in)
		case '}', ']', ')': // This is a closing symbol - is there a matching opening symbol on the stack?
			if len(stack) == 0 {
				return false, nil
			}
			stack, top = pop(stack)
			switch pair := string(top) + string(in); pair {
			case "{}", "[]", "()":
				continue
			default:
				return false, nil
			}
		}
	}

	return len(stack) == 0, nil
}
