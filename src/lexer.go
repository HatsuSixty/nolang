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
	l int
	r int
	c int
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

	emptyloc := Location{l: 0, r: 0, c: 0}

	if !(TOKEN_COUNT == 2) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of Tokens in lexfile()\n")
		os.Exit(1)
	}

	var finalstring string
	for i := range string(source) {
		curchar := string(source[i])
		if (unicode.IsSpace(rune(source[i]))) && (finalstring != "") {

			switch {
			case isNumber(finalstring):
				i, err = strconv.Atoi(finalstring)
				tokens = append(tokens, Token{kind: TOKEN_INT,  icontent: i,        loc: emptyloc})
			case finalstring == "+":
				tokens = append(tokens, Token{kind: TOKEN_WORD, scontent: "+",      loc: emptyloc})
			case finalstring == "-":
				tokens = append(tokens, Token{kind: TOKEN_WORD, scontent: "-",      loc: emptyloc})
			case finalstring == "*":
				tokens = append(tokens, Token{kind: TOKEN_WORD, scontent: "*",      loc: emptyloc})
			case finalstring == "divmod":
				tokens = append(tokens, Token{kind: TOKEN_WORD, scontent: "divmod", loc: emptyloc})
			case finalstring == "drop":
				tokens = append(tokens, Token{kind: TOKEN_WORD, scontent: "drop",   loc: emptyloc})
			case finalstring == "print":
				tokens = append(tokens, Token{kind: TOKEN_WORD, scontent: "print",  loc: emptyloc})
			default:
				// ignore and treat as word, then, in the parsing phase, the compiler will
				// do the error reporting
				tokens = append(tokens, Token{kind: TOKEN_WORD, scontent: finalstring,  loc: emptyloc})
			}

			finalstring = ""
		} else {
			finalstring += curchar
		}
	}

	return tokens
}
