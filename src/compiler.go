package main

import "fmt"

func pop(stack []uint) ([]uint, uint) {
	poppedVal := stack[len(stack)-1]
	slice     := stack[:len(stack)-1]
	return slice, poppedVal
}

func compile_program(program []Op) {
	stack := []uint{}
	for i := range program {
		switch program[i].op {
		case OP_PUSH_INT:
			stack     = append(stack, uint(program[i].operand))
		case OP_PLUS:
			var a, b uint
			stack, a = pop(stack)
			stack, b = pop(stack)
			stack    = append(stack, uint(a + b))
		case OP_PRINT:
			var a uint
			stack, a = pop(stack)
			fmt.Print(a)
		}
	}
}
