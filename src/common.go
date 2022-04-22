package main

import (
	"os/exec"
	"fmt"
)

func cmdRunEchoed(args string, silent bool) {
	if !silent {
		fmt.Println("[CMD]", args)
	}

	err := exec.Command(args)

	if err != nil {
		fmt.Errorf("ERROR: Could not run command: %s", err.String())
	}
}
