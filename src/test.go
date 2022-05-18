package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"path/filepath"
	"bytes"
	"bufio"
)

var testfailed int = 0
var ignored    int = 0
var nobin      string

type TestCase struct {
	stdout string
	stderr string
}

func fileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func cmdRunReturnTestCase(args string) TestCase {
	fmt.Printf("[CMD] %s\n", args)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("/bin/sh", "-c", args)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Shell command error: %s\n", err)
		fmt.Fprintf(os.Stderr, "Output:\n")
		fmt.Fprintf(os.Stderr, "    stdout:\n")
		fmt.Fprintf(os.Stderr, "        " + stdout.String() + "\n")
		fmt.Fprintf(os.Stderr, "    stderr:\n")
		fmt.Fprintf(os.Stderr, "        " + stderr.String() + "\n")
		testfailed += 1
	}
	return TestCase{stdout: stdout.String(), stderr: stderr.String()}
}

func printTestCase(test TestCase) {
	fmt.Printf("    stdout:\n%s\n\nstderr:\n%s\n\n",
		test.stdout, test.stderr)
}

func testCaseEqual(a TestCase, b TestCase) bool {
	return ((a.stdout == b.stdout) && (a.stderr == b.stderr))
}

func loadTestCaseForFile(file string, folder string) (TestCase, bool) {
	testcpath := folder + "/" + fileNameWithoutExt(filepath.Base(file)) + ".txt"

	var tstdout string
	var tstderr string
	var ignoredfile bool = false

	var instderr bool = false
	var instdout bool = false

	if fileExists(testcpath) {
		f, err := os.Open(testcpath)
		if err != nil {}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			switch {
			case line == "stderr:::":
				instderr = true
				instdout = false
			case line == "stdout:::":
				instdout = true
				instderr = false
			default:
				switch {
				case instdout:
					tstdout += string(line) + "\n"
				case instderr:
					tstderr += string(line) + "\n"
				}
			}
			if !instdout && !instderr {
				fmt.Fprintf(os.Stderr, "ERROR: Invalid format\n")
				os.Exit(1)
			}
		}
		f.Close()
	} else {
		fmt.Fprintf(os.Stderr, "[WARNING] No output file for `%s` encountered, just testing if it compiles.\n",
			file)
		ignoredfile = true
		ignored += 1
	}

	return TestCase{stdout: tstdout, stderr: tstderr}, ignoredfile
}

func runTestForFile(file string, folder string) {
	fmt.Printf("[INFO] Testing %s\n", file)
	wd, err := os.Getwd()
	if err != nil {}
	cmdRunEchoInfo("rm -f " + wd + "/" + folder + "/output*", true)
	cmdRunEchoInfo("rm -f " + wd + "/output*", true)

	tce, ignored := loadTestCaseForFile(file, folder)
	tci          := cmdRunReturnTestCase(nobin + " -c " + file + " -r -s")

	if !testCaseEqual(tce, tci) && !ignored {
		fmt.Fprintf(os.Stderr, "ERROR: Test failed:\n")
		fmt.Fprintf(os.Stderr, "    Expected:\n")
		printTestCase(tce)
		fmt.Fprintf(os.Stderr, "    Got:\n")
		printTestCase(tci)
		testfailed += 1
	}
}

func runTestForFolder(folder string) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could not read files of folder `%s`: %s\n",
			folder, err)
		os.Exit(1)
	}

	for _, f := range files {
		if !f.IsDir() {
			if filepath.Ext(f.Name()) == ".no" {
				runTestForFile(folder + "/" + f.Name(), folder)
			}
		}
	}

	fmt.Printf("\nFailed %d, Ignored %d\n", testfailed, ignored)
}

func saveTestCase(tc TestCase, file string) {
	f, err := os.OpenFile(file, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0644)
	if isError(err) {
		os.Exit(3)
	}

	f.WriteString("stdout:::\n")
	f.WriteString(tc.stdout)
	f.WriteString("stderr:::\n")
	f.WriteString(tc.stderr)

	f.Close()
}

func updateOutputForFile(file string, folder string) {
	fmt.Printf("[INFO] Updating output of file %s\n", file)
	wd, err := os.Getwd()
	if err != nil {}
	cmdRunEchoInfo("rm -f " + wd + "/" + folder + "/output*", true)
	cmdRunEchoInfo("rm -f " + wd + "/output*", true)

	tc := cmdRunReturnTestCase(nobin + " -c " + file + " -r -s")
	saveTestCase(tc, fileNameWithoutExt(file) + ".txt")
}

func updateOutputForFolder(folder string) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could not read files of folder `%s`: %s\n",
			folder, err)
		os.Exit(1)
	}

	for _, f := range files {
		if !f.IsDir() {
			if filepath.Ext(f.Name()) == ".no" {
				updateOutputForFile(folder + "/" + f.Name(), folder)
			}
		}
	}
}

func usage() {
	fmt.Println("Usage:"                                                                                   )
	fmt.Println("    ./test <SUBCOMMAND>"                                                                  )
	fmt.Println("Subcommands:"                                                                             )
	fmt.Println("    update   [FOLDER]         Update output of files in [FOLDER], [FOLDER] is optional"   )
	fmt.Println("    run|full [FOLDER]         Run a test for each file in [FOLDER], [FOLDER] is optional" )
	fmt.Println("    help                      Print this help & exit"                                     )
}

func main() {
	if !(len(os.Args) > 1) {
		fmt.Fprintf(os.Stderr, "ERROR: No argument was provided\n")
		usage()
		os.Exit(4)
	}

	subcommand := os.Args[1]
	folder     := "./tests/"
	if len(os.Args) > 2 {
		folder  = os.Args[2]
	}

	wd, err := os.Getwd()
	if err != nil {}

	switch {
	case fileExists(wd + "/no"):    nobin = wd + "/no"
	case fileExists(wd + "/../no"): nobin = wd + "/../no"
	default:                        nobin = "no"
	}

	switch {
	case subcommand == "update":
		updateOutputForFolder(folder)
	case subcommand == "run" || subcommand == "full":
		runTestForFolder(folder)
	case subcommand == "help":
		usage()
		os.Exit(0)
	default:
		fmt.Printf("ERROR: Unknown subcommand: %s\n", subcommand)
		usage()
		os.Exit(4)
	}
}
