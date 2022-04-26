package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
	"strconv"
	"os"
)

/// Locations ///

type Location struct {
	f string
	r    int
	c    int
}

/////////////////

/// Tokens ///

type TokenKind int
const (
	TOKEN_INT   TokenKind = iota
	TOKEN_WORD  TokenKind = iota
	TOKEN_COUNT TokenKind = iota
)

type Token struct {
	kind     TokenKind
	icontent int
	scontent string
	loc      Location
}

//////////////

func lexfile(filepath string) []Token {
	tokens := []Token{}

	source, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could not read content of file %s: %s\n", filepath, err)
		os.Exit(3)
	}

	f := filepath
	r := 1
	c := 1

	if !(TOKEN_COUNT == 2) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of Tokens in lexfile()\n")
		os.Exit(1)
	}

	var finalstring string

	for i := range string(source) {
		curchar := string(source[i])

		// TODO: if there is no space at the end of the file,
		// it does not recognize "finalstring" as a word
		if (unicode.IsSpace(rune(source[i]))) && (finalstring != "") {

			switch {
			case isNumber(finalstring):
				i, err = strconv.Atoi(finalstring)
				tokens = append(tokens,
					Token{kind: TOKEN_INT,
						icontent: i,
						loc: Location{f, r, c - len(finalstring)}})

			case isWord(finalstring):
				tokens = append(tokens,
					Token{kind: TOKEN_WORD,
						scontent: finalstring,
						loc: Location{f, r, c - len(finalstring)}})
			}

			finalstring = ""
		} else {
			finalstring += curchar
		}

		if curchar == "\n" || curchar == "\r" {
			r += 1
			c  = 0
		}
		c += 1
	}

	return tokens
}
