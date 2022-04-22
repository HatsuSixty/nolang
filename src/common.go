package main

import (
	"fmt"
	"os/exec"
)

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
