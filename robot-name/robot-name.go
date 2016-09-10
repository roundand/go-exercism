// Package robotname provides a Robot type with Name and Reset methods.
package robotname

import "math/rand"

var used = map[string]bool{}

// Robot is a mechanical servant with a name.
type Robot struct {
	name string
}

// Name returns the name of a Robot, generating it first if necessary.
func (rob *Robot) Name() string {
	if rob.name == "" {
    // generate an unused new name
    for rob.name = newName(); used[rob.name] == true; rob.name = newName() {}
    used[rob.name] = true
	}
	return rob.name
}

func newName() string {
  b := make([]byte, 5)
  b[0] = byte('A' + rand.Intn(26))
  b[1] = byte('A' + rand.Intn(26))
  b[2] = byte('0' + rand.Intn(10))
  b[3] = byte('0' + rand.Intn(10))
  b[4] = byte('0' + rand.Intn(10))
  return string(b)
}

// Reset blanks the name of a Robot, so that the next call to Name will generate a new name.
func (rob *Robot) Reset() {
	rob.name = ""
}
