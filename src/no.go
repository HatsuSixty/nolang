package main

var programm = []Op{
	{op: OP_PUSH_INT, operand: 34},
	{op: OP_PUSH_INT, operand: 35},
	{op: OP_PLUS},
	{op: OP_PRINT},
	{op: OP_PUSH_INT, operand: 34},
	{op: OP_PUSH_INT, operand: 35},
	{op: OP_PLUS},
	{op: OP_PRINT},
}

func main() {
	generateYasmLinux_x86_64(programm, "output.asm")
	cmdRunEchoInfo("yasm -felf64 output.asm", false)
	cmdRunEchoInfo("ld -o output output.o",   false)
}
