package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/// Locations ///

type Location struct {
	l int
	r int
	c int
}

/////////////////

/// Tokens ///

type TokenKind int
const (
	TOKEN_INT  TokenKind = iota
	TOKEN_WORD TokenKind = iota
)

type Token struct {
	kind     TokenKind
	icontent int
	scontent string
	loc      Location
}

//////////////

func lexfile(filepath string) []Token {
	source, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could not read content of file %s: %s\n", filepath, err)
		os.Exit(3)
	}

	fmt.Print(source)

	tokens := []Token{}
	return tokens
}
