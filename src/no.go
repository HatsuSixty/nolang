package main

func main() {
	program := compileFileIntoOps("test.no")
	generateYasmLinux_x86_64(program, "output.asm" )
	cmdRunEchoInfo("yasm -felf64 output.asm", false)
	cmdRunEchoInfo("ld -o output output.o",   false)
}
