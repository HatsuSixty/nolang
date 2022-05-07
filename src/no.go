package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var silent  bool
	var run     bool
	var cfile   string
	var outfile string

	flag.StringVar(&cfile,   "c", "", "compile the specified program")
	flag.StringVar(&outfile, "o", "output", "change the output file name to the specified one")
	flag.BoolVar(&run, "r", false, "run the program after successful compilation")
	flag.BoolVar(&silent, "s", false, "do not show output (except errors)")
	flag.Parse()

	if cfile == "" {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "ERROR: Input file not provided\n")
		os.Exit(4)
	}

	if !flag.Parsed() {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "ERROR: Unknown error while parsing flags\n")
		os.Exit(4)
	}

	program := compileFileIntoOps(cfile)
	outasm := outfile + ".asm"
	outobj := outfile + ".o"
	generateYasmLinux_x86_64(program, outasm)
	cmdRunEchoInfo("yasm -felf64 " + outasm, silent)
	cmdRunEchoInfo("ld -o " + outfile + " " + outobj,   silent)

	if run {
		wd, err := os.Getwd()
		if err != nil {}
		cmdRunEchoInfo(wd + "/" + outfile, false)
	}
}
