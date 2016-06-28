// Package secret encodes a number as a sequence of body actions.
package secret

// Handshake converts a number to an array of strings.
func Handshake(secret int) []string {

	var signal []string // declare and initialise signal as a slice of strings

	// wrap builtin append() in a regular function for ease of reference
	var appender = func(ss []string, s string) []string {
		return append(ss, s)
	}

	// Reverse the order of the operations in the secret handshake
	var reverser = func(ss []string, _ string) []string {
		for i, l := 0, len(signal)-1; i < ((l + 1) / 2); i++ {
			ss[i], ss[l-i] = ss[l-i], ss[i]
		}
		return ss
	}

	// The program should consider strings specifying an invalid binary as the value 0.
	if secret < 0 {
		return signal
	}

	// here's how we encode each bit of a valid binary as an action
	var steps = []struct {
		bit    int
		encode func([]string, string) []string
		action string
	}{
		{1 << 0, appender, "wink"},
		{1 << 1, appender, "double blink"},
		{1 << 2, appender, "close your eyes"},
		{1 << 3, appender, "jump"},
		{1 << 4, reverser, ""},
	}

	for _, step := range steps {
		if (secret & step.bit) != 0 {
			signal = step.encode(signal, step.action)
		}
	}

	return signal
}
