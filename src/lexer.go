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
	TOKEN_STR   TokenKind = iota
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

	if !(TOKEN_COUNT == 3) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of Tokens in lexfile\n")
		os.Exit(1)
	}

	var finalstring string

	c := 0
	for c < len(line) {
		curchar := rune(line[c])

		if unicode.IsSpace(curchar) && (finalstring != "") {

			switch {
			case finalstring[0] == '/':
				c -= len(finalstring)
				if ((len(line)-1) > c+1) {
					if (line[c] == '/' && line[c+1] == '/') {
						goto yeahiquit
					}
				}
				// else
				if !isQuote(rune(finalstring[0])) { // treat as word
					tokens = append(tokens,
						Token{kind: TOKEN_WORD,
							scontent: "/",
							loc: Location{loc.f, loc.r, c + 1}})
				}


			case isNumber(finalstring):
				i, err := strconv.Atoi(finalstring)
				if err != nil {}
				tokens = append(tokens,
					Token{kind: TOKEN_INT,
						icontent: i,
						loc: Location{loc.f, loc.r, c - len(finalstring) + 1}})

			case !isQuote(rune(finalstring[0])):
				for ch := range finalstring {
					if isQuote(rune(finalstring[ch])) {
						fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: String prefixes are not allowed\n",
							loc.f, loc.r, c - len(finalstring))
						os.Exit(1)
					}
				}
				tokens = append(tokens,
					Token{kind: TOKEN_WORD,
						scontent: finalstring,
						loc: Location{loc.f, loc.r, c - len(finalstring) + 1}})

			case isQuote(rune(finalstring[0])):
				c -= len(finalstring)
				c += 1
				str := ""
				strclosed := false
				for c < len(line) {
					if isQuote(rune(line[c])) {
						strclosed = true
						break
					}

					if line[c] == '\\' {
						if !((len(line)-1) > c+1) {
							fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Expected escape character but got nothing\n",
								loc.f, loc.r, c + 1)
							os.Exit(1)
						}

						c += 1
						escapechar := line[c]

						switch escapechar {
						case 'n':
							str += "\n"
						case 't':
							str += "\t"
						case 'r':
							str += "\r"
						case '\'':
							str += "'"
						case '"':
							str += "\""
						case '\\':
							str += "\\"
						default:
							fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unknown escape character: %c\n",
								loc.f, loc.r, c + 1, escapechar)
							os.Exit(1)
						}

						c += 1
						if isQuote(rune(line[c])) {
							strclosed = true
							break
						}
					}
					str += string(line[c])
					c += 1
				}

				if !strclosed {
					fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: String not closed\n",
						loc.f, loc.r, c - len(str) + 1)
					os.Exit(1)
				}

				c += 1
				postfix := ""
				for c < len(line) {
					if unicode.IsSpace(rune(line[c])) {
						break
					} else {
						postfix += string(line[c])
					}
					c += 1
				}

				ischar := false

				switch {
				case postfix == "ch":
					ischar = true
				default:
					if postfix != "" {
						fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unknown postfix: %s\n",
							loc.f, loc.r, c - len(postfix) + 1, postfix)
						os.Exit(1)
					}
				}

				switch {
				case ischar:
					if !(len(str) == 1) {
						fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Character literals can only contain 1 character\n",
							loc.f, loc.r, c - len(postfix) + 1)
						os.Exit(1)
					}

					tokens = append(tokens,
						Token{kind: TOKEN_INT,
							icontent: int(str[0]),
							loc: Location{loc.f, loc.r, c - len(str) + 1}})

				default:
					tokens = append(tokens,
						Token{kind: TOKEN_STR,
							scontent: str,
							loc: Location{loc.f, loc.r, c - len(str) + 1}})
				}
			}

			finalstring = ""
		} else {
			if !(unicode.IsSpace(curchar)) {
				finalstring += string(curchar)
			}
		}
		c += 1
	}

yeahiquit:
	return tokens
}

func lexfile(filepath string) []Token {
	tokens := []Token{}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could not open file `%s`: %s\n", filepath, err)
		os.Exit(3)
	}

	scanner := bufio.NewScanner(file)

	f := filepath
	r := 1

	for scanner.Scan() {
		tokens = append(tokens, lexline(scanner.Text(), Location{f: f, r: r})...)
		r += 1
	}

	if err = scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could read file `%s`: %s", filepath, err)
	}

	file.Close()
	return tokens
}
