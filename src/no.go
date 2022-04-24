package main

var programm = []Op{
	{op: OP_PUSH_INT, operand: 34},
	{op: OP_PUSH_INT, operand: 35},
	{op: OP_PLUS},
	{op: OP_PRINT},
	{op: OP_PUSH_INT, operand: 34},
	{op: OP_PUSH_INT, operand: 35},
	{op: OP_PLUS},
	{op: OP_PUSH_INT, operand: 34},
	{op: OP_MINUS},
	{op: OP_PUSH_INT, operand: 34},
	{op: OP_PLUS},
	{op: OP_PUSH_INT, operand: 2},
	{op: OP_DIVMOD},
	{op: OP_DROP},
	{op: OP_PUSH_INT, operand: 2},
	{op: OP_MULT},
	{op: OP_PRINT},
}

func main() {
	lexfile("test.no")
	generateYasmLinux_x86_64(programm, "output.asm")
	cmdRunEchoInfo("yasm -felf64 output.asm", false)
	cmdRunEchoInfo("ld -o output output.o",   false)
}
