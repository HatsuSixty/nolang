package main

////////// GENERATOR //////////
import (
	"fmt"
	"os"
	"strconv"
)

type OpType int
const (
	OP_PLUS     OpType = iota
	OP_MINUS    OpType = iota
	OP_MULT     OpType = iota
	OP_DIVMOD   OpType = iota
	OP_DROP     OpType = iota
	OP_PRINT    OpType = iota
	OP_PUSH_INT OpType = iota
	OP_COUNT    OpType = iota
)

type Operand int
type Op struct {
	op      OpType
	operand Operand
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func generateYasmLinux_x86_64(program []Op, output string) {
	if !(OP_COUNT == 7) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of ops in generateYasmLinux_x86_64\n")
		os.Exit(1)
	}

	f, err := os.OpenFile(output, os.O_RDWR | os.O_CREATE, 0644)
	if isError(err) {
		os.Exit(3)
	}

	f.WriteString("BITS 64\n"                              )
	f.WriteString("segment .text\n"                        )
	f.WriteString("print:\n"                               )
	f.WriteString("    mov     r9, -3689348814741910323\n" )
	f.WriteString("    sub     rsp, 40\n"                  )
	f.WriteString("    mov     BYTE [rsp+31], 10\n"        )
	f.WriteString("    lea     rcx, [rsp+30]\n"            )
	f.WriteString(".L2:\n"                                 )
	f.WriteString("    mov     rax, rdi\n"                 )
	f.WriteString("    lea     r8, [rsp+32]\n"             )
	f.WriteString("    mul     r9\n"                       )
	f.WriteString("    mov     rax, rdi\n"                 )
	f.WriteString("    sub     r8, rcx\n"                  )
	f.WriteString("    shr     rdx, 3\n"                   )
	f.WriteString("    lea     rsi, [rdx+rdx*4]\n"         )
	f.WriteString("    add     rsi, rsi\n"                 )
	f.WriteString("    sub     rax, rsi\n"                 )
	f.WriteString("    add     eax, 48\n"                  )
	f.WriteString("    mov     BYTE [rcx], al\n"           )
	f.WriteString("    mov     rax, rdi\n"                 )
	f.WriteString("    mov     rdi, rdx\n"                 )
	f.WriteString("    mov     rdx, rcx\n"                 )
	f.WriteString("    sub     rcx, 1\n"                   )
	f.WriteString("    cmp     rax, 9\n"                   )
	f.WriteString("    ja      .L2\n"                      )
	f.WriteString("    lea     rax, [rsp+32]\n"            )
	f.WriteString("    mov     edi, 1\n"                   )
	f.WriteString("    sub     rdx, rax\n"                 )
	f.WriteString("    xor     eax, eax\n"                 )
	f.WriteString("    lea     rsi, [rsp+32+rdx]\n"        )
	f.WriteString("    mov     rdx, r8\n"                  )
	f.WriteString("    mov     rax, 1\n"                   )
	f.WriteString("    syscall\n"                          )
	f.WriteString("    add     rsp, 40\n"                  )
	f.WriteString("    ret\n"                              )
	f.WriteString("global _start\n"                        )
	f.WriteString("_start:\n"                              )
	for i := range program {
		switch program[i].op {
		case OP_PUSH_INT:
			f.WriteString("    ;; -- push --\n"   )
			f.WriteString("    mov rax, " + strconv.FormatUint(uint64(program[i].operand), 10) + "\n")
			f.WriteString("    push rax\n"        )
		case OP_PLUS:
			f.WriteString("    ;; -- plus --\n"   )
			f.WriteString("    pop rax\n"         )
			f.WriteString("    pop rbx\n"         )
			f.WriteString("    add rax, rbx\n"    )
			f.WriteString("    push rax\n"        )
		case OP_MINUS:
			f.WriteString("    ;; -- minus --\n"  )
			f.WriteString("    pop rax\n"         )
			f.WriteString("    pop rbx\n"         )
			f.WriteString("    sub rbx, rax\n"    )
			f.WriteString("    push rbx\n"        )
		case OP_MULT:
			f.WriteString("    ;; -- mult --\n"   )
			f.WriteString("    pop rax\n"         )
			f.WriteString("    pop rbx\n"         )
			f.WriteString("    mul rbx\n"         )
			f.WriteString("    push rax\n"        )
		case OP_DIVMOD:
			f.WriteString("    ;; -- divmod --\n" )
			f.WriteString("    xor rdx, rdx\n"    )
			f.WriteString("    pop rbx\n"         )
			f.WriteString("    pop rax\n"         )
			f.WriteString("    div rbx\n"         )
			f.WriteString("    push rax\n"        )
			f.WriteString("    push rdx\n"        )
		case OP_DROP:
			f.WriteString("    ;; -- drop --\n"   )
			f.WriteString("    pop rax\n"         )
		case OP_PRINT:
			f.WriteString("    ;; -- print --\n"  )
			f.WriteString("    pop rdi\n"         )
			f.WriteString("    call print\n"      )
		default:
			fmt.Fprintf(os.Stderr, "ERROR: Unreachable\n")
			os.Exit(2)
		}
	}
	f.WriteString("    ;; -- built-in exit --\n" )
	f.WriteString("    mov rax, 60\n"            )
	f.WriteString("    mov rdi, 0\n"             )
	f.WriteString("    syscall\n"                )

	f.Close()
}
///////////////////////////////

/////////// PARSER ///////////

func compileTokensIntoOps(tokens []Token) []Op {
	var ops []Op

	for i := range tokens {
		token := tokens[i]

		switch token.kind {
		case TOKEN_INT:
			ops = append(ops, Op{op: OP_PUSH_INT, operand: Operand(token.icontent)})
		case TOKEN_WORD:
			switch {
			case token.scontent == "+":
				ops = append(ops, Op{op: OP_PLUS})
			case token.scontent == "-":
				ops = append(ops, Op{op: OP_MINUS})
			case token.scontent == "*":
				ops = append(ops, Op{op: OP_MULT})
			case token.scontent == "divmod":
				ops = append(ops, Op{op: OP_DIVMOD})
			case token.scontent == "drop":
				ops = append(ops, Op{op: OP_DROP})
			case token.scontent == "print":
				ops = append(ops, Op{op: OP_PRINT})
			default:
				fmt.Fprintf(os.Stderr, "ERROR: Unknown word: %s", token.scontent)
				os.Exit(1)
			}
		default:
			fmt.Fprintf(os.Stderr, "ERROR: Unreachable\n")
			os.Exit(1)
		}
	}
	return ops
}

//////////////////////////////

////////// COMPILER //////////

func compileFileIntoOps(filepath string) []Op {
	tokens := lexfile(filepath)
	ops    := compileTokensIntoOps(tokens)
	return ops
}

//////////////////////////////
