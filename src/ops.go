package main

type OpType int
const (
	OP_PLUS           OpType = iota
	OP_PRINT          /////////////
	OP_PUSH_INT       /////////////
	OP_COUNT          /////////////
)

type Operand int
type Op struct {
	op      OpType
	operand Operand
}
