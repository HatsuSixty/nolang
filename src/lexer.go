package main

import (
	"fmt"
	"unicode"
	"strconv"
	"bufio"
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

func lexline(line string, loc Location) []Token {
	tokens := []Token{}

	line += " "

	if !(TOKEN_COUNT == 2) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of Tokens in lexfile()\n")
		os.Exit(1)
	}

	c := 1

	var finalstring string

	for c := range line {
		curchar := rune(line[c])

		if unicode.IsSpace(curchar) && (finalstring != "") {

			switch {
			case isNumber(finalstring):
				i, err := strconv.Atoi(finalstring)
				if err != nil {}
				tokens = append(tokens,
					Token{kind: TOKEN_INT,
						icontent: i,
						loc: Location{loc.f, loc.r, c - len(finalstring)}})

			case isWord(finalstring):
				tokens = append(tokens,
					Token{kind: TOKEN_WORD,
						scontent: finalstring,
						loc: Location{loc.f, loc.r, c - len(finalstring)}})
			}
			finalstring = ""
		} else {
			if !(unicode.IsSpace(curchar)) {
				finalstring += string(curchar)
			}
		}
	}
	c += 1

	return tokens
}

func lexfile(filepath string) []Token {
	tokens := []Token{}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could not open file `%s`: %s", filepath, err)
		os.Exit(3)
	}

	scanner := bufio.NewScanner(file)

	f := filepath
	r := 1

	for scanner.Scan() {
		tokens = append(tokens, lexline(scanner.Text(), Location{f: f, r: 1})...)
		r += 1
	}

	if err = scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could read file `%s`: %s", filepath, err)
	}

	file.Close()
	return tokens
}
