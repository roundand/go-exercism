// Package wordy provides a function to parse and solve prose arithmetic expression.
package wordy

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type stateMap map[string]eventMap
type eventMap map[string]move
type move struct {
	action    func(n, i int) int
	nextState string
}

var sm = stateMap{
	"start": {
		"WhatIs": {func(n, i int) int { return i }, "op"},
	},
	"op": {
		"plus":         {func(n, i int) int { return n + i }, "op"},
		"minus":        {func(n, i int) int { return n - i }, "op"},
		"multipliedBy": {func(n, i int) int { return n * i }, "op"},
		"dividedBy":    {func(n, i int) int { return n / i }, "op"},
		"?":            {func(n, i int) int { return n }, "end"},
	},
}

// Answer accepts a prose arithmetic expression and returns an integer solution, or failure.
func Answer(e string) (int, bool) {
	e = strings.Replace(e, "?", " ? 0", -1) // question mark is final token, with dummy value.
	e = strings.Replace(e, " is", "Is", -1) // Simplify "What is" text
	e = strings.Replace(e, " by", "By", -1) // Simplify "multiplied by" and "divided by" text
	t := strings.Split(e, " ")              // tokenise input.
	opVals := opValStreamer(t)              // initialise event feeder with tokens
	return sm.run("start", 0, opVals)
}

// fsm engine
func (sm stateMap) run(st string, n int, opVals func() (string, int, error)) (int, bool) {
	if st == "end" {
		return n, true
	}
	op, val, err := opVals()
	if err != nil {
		return failed(fmt.Sprintf("couldn't parse val %q for op %q", val, op))
	}
	m, found := sm[st][op]
	if !found {
		return failed(fmt.Sprintf("Couldn't find event %q for state %q", op, st))
	}

	return sm.run(m.nextState, m.action(n, val), opVals)
}

// returns operator and optional value, or parse Int error
func opValStreamer(tokens []string) func() (string, int, error) {
	t := tokens // capture tokens as closure
	return func() (string, int, error) {
		opTok, t1 := nextToken(t)
		t = t1
		if opTok == "" {
			return "", 0, nil
		}

		valTok, t2 := nextToken(t)
		t = t2
		if valTok == "" {
			return opTok, 0, nil
		}

		val, err := strconv.Atoi(valTok)
		if err != nil {
			return "", 0, err
		}

		return opTok, val, nil
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
