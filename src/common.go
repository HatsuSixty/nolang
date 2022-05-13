package main

import (
	"fmt"
	"unicode"
	"os/exec"
	"os"
	"bytes"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func isNumber(str string) bool {
	for i := range str {
		if !unicode.IsNumber(rune(str[i])) {
			return false
		}
	}
	return true
}

func isQuote(char rune) bool {
	if char == '\'' || char == '"' {
		return true
	}
	return false
}

func popInt(stack []int) ([]int, int) {
	mslice := stack
	poppedVal := mslice[len(mslice)-1]
	mslice = mslice[:len(mslice)-1]
	return mslice, poppedVal
}

func in(str string, slc []string) bool {
	for i := range slc {
		if slc[i] == str {
			return true
		}
	}
	return false
}

func fileExists(path string) bool {
	returnval := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		returnval = false
	}
	return returnval
}

func cmdRunEchoInfo(args string, silent bool) {
	if !silent {
		fmt.Println("[CMD]", args)
	}

	var buf bytes.Buffer

	cmd := exec.Command("/bin/sh", "-c", args)
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	err := cmd.Run()
	if !silent && !(buf.String() == "") {
		fmt.Print(buf.String())
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Shell command error: %s\n", err)
		os.Exit(1)
	}
}
