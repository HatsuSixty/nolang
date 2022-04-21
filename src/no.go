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
	compile_program(programm)
}
