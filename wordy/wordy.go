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

// an experiment in FSM using golang.
// First pass, all hard-coded.

// start
// :what -> is
func start(t []string) (int, bool) {
	switch nextToken(t) {
	case "WhatIs":
		return whatIs(t[1:])
	default:
		return failed(fmt.Sprintf("Expected 'What' but got %q.\n", nextToken(t)))
	}
}

// whatIs
// :[n] (x := n)-> op
func whatIs(t []string) (int, bool) {
	s := nextToken(t)
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Could't parse %q as int.\n", s)
		return failed(err.Error())
	}
	return op(i, t[1:])
}

// op
// :(nil) -> end
// :plus -> plus
// :minus -> minus
// :multiplied -> multipliedBy
// :divided -> dividedBy
func op(n int, t []string) (int, bool) {
	switch nextToken(t) {
	case "":
		return n, true
	case "plus":
		return plus(n, t[1:])
	case "minus":
		return minus(n, t[1:])
	case "multipliedBy":
		return multipliedBy(n, t[1:])
	case "dividedBy":
		return dividedBy(n, t[1:])
	default:
		return failed(fmt.Sprintf("Expected <operator> but got %q.\n", nextToken(t)))
	}
}

// plus
// :[n] (x := x+n) -> op
func plus(n int, t []string) (int, bool) {
	s := nextToken(t)
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Could't parse %q as int.\n", s)
		return failed(err.Error())
	}
	return op(n+i, t[1:])
}

// minus
// :[n] (x := x-n) -> op
func minus(n int, t []string) (int, bool) {
	s := nextToken(t)
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Could't parse %q as int.\n", s)
		return failed(err.Error())
	}
	return op(n-i, t[1:])
}

// multipliedBy
// :[n] -> (x := x*n) -> op
func multipliedBy(n int, t []string) (int, bool) {
	s := nextToken(t)
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Could't parse %q as int.\n", s)
		return failed(err.Error())
	}
	return op(n*i, t[1:])
}

// dividedBy
// :n (x := x /n) -> op
func dividedBy(n int, t []string) (int, bool) {
	s := nextToken(t)
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Could't parse %q as int.\n", s)
		return failed(err.Error())
	}
	return op(n/i, t[1:])
}

// returns next token, or empty string on end of input
func nextToken(t []string) string {
	if len(t) == 0 {
		return ""
	}
	return t[0]
}

// central spot for logging errors
func failed(msg string) (int, bool) {
	log.Printf("Error message: %q\n", msg)
	return 0, false
}
