package main

import (
	"fmt"
	"unicode"
	"os/exec"
)

func isNumber(str string) bool {
	for i := range str {
		if !unicode.IsNumber(rune(str[i])) {
			return false
		}
	}
	return true
}

func isString(str string) bool {
	firstchar := ([]rune(str)[0])
	if firstchar == '"' || firstchar == '\'' {
		return true
	}
	return false
}

func isWord(word string) bool {
	if isString(word) || isNumber(word) {
		return false
	}
	return true
}

func popInt(stack []int) ([]int, int) {
	mslice := stack
	poppedVal := mslice[len(mslice)-1]
	mslice = mslice[:len(mslice)-1]
	return mslice, poppedVal
}

func cmdRunEchoInfo(args string, silent bool) {
	if !silent {
		fmt.Println("[CMD]", args)
	}

	command := exec.Command("/bin/sh", "-c", args)

	err := command.Run()
	if err != nil {
		fmt.Errorf("ERROR: Could not run command\n")
	}
}
