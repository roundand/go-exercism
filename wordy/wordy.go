// Package wordy provides a function to parse and solve prose arithmetic expression.
package wordy

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Answer accepts a prose arithmetic expression and returns an integer solution, or failure.
func Answer(e string) (int, bool) {
	e = strings.Replace(e, "?", "", -1)                         // question marks can be removed.
	e = strings.Replace(e, "What is", "WhatIs", -1)             // Simplify What is prose
	e = strings.Replace(e, "multiplied by", "multipliedBy", -1) // Simplify multiplication prose
	e = strings.Replace(e, "divided by", "dividedBy", -1)       // Simplify division prose
	t := strings.Split(e, " ")                                  // tokenise input.
	return start(t)
}

// start
// :what -> is
func start(t []string) (int, bool) {
	tok, t := nextToken(t)
	if tok == "WhatIs" {
		return whatIs(t)
	}
	return failed(fmt.Sprintf("Expected 'What' but got %q.\n", tok))
}

// whatIs
// :[n] (x := n)-> op
func whatIs(t []string) (int, bool) {
	s, t := nextToken(t)
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Could't parse %q as int.\n", s)
		return failed(err.Error())
	}
	return op(i, t)
}

// op
func op(n int, t []string) (int, bool) {
	opTok, t := nextToken(t)
	if opTok == "" { // we're all done here!
		return n, true
	}

	// we expect each of our operators to be folloewd by a value
	valTok, t := nextToken(t)
	val, err := strconv.Atoi(valTok)
	if err != nil {
		return failed(err.Error())
	}

	switch opTok {
	case "plus":
		return op(n+val, t) //plus(n, t[1:])
	case "minus":
		return op(n-val, t)
	case "multipliedBy":
		return op(n*val, t)
	case "dividedBy":
		return op(n/val, t)
	default:
		return failed(fmt.Sprintf("Expected <operator> but got %q.\n", opTok))
	}
}

// returns next token, or empty string on end of input
func nextToken(t []string) (string, []string) {
	if len(t) == 0 {
		return "", t
	}
	return t[0], t[1:]
}

// central spot for logging errors
func failed(msg string) (int, bool) {
	log.Printf("Error message: %q\n", msg)
	return 0, false
}
