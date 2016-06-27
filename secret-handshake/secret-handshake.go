// secret converts a number to a sequence of body actions.
package secret

// Handshake converts a number to an array of strings
func Handshake (secret int) []string {
  var signal []string                           // declare and initialise signal as a slice of strings

  if (secret < 0) {
    return signal
  }

  if (secret & (1 << 0) != 0) {              // 1 = wink
    signal = append(signal, "wink")
  }

  if (secret & (1 << 1) != 0) {                 // 10 = double blink
    signal = append(signal, "double blink")
  }

  if (secret & (1 << 2) != 0) {                 // 100 = close your eyes
    signal = append(signal, "close your eyes")
  }

  if (secret & (1 << 3) != 0) {                 // 1000 = jump
    signal = append(signal, "jump")
  }

  if (secret & (1 << 4) != 0) {                 // 10000 = Reverse the order of the operations in the secret handshake.
    l := len(signal) - 1
    for i := 0; i < (len(signal) / 2); i++ {
      signal[i], signal[l-i] = signal[l-i], signal[i]
    }
  }

  return signal
}
