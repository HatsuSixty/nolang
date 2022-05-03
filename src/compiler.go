package main

////////// GENERATOR //////////
import (
	"fmt"
	"os"
	"strconv"
)

// should be enough for everyone
const MEM_CAP int = 64000

type OpType int
const (
	// basic operators
	OP_PLUS     OpType = iota
	OP_MINUS    OpType = iota
	OP_MULT     OpType = iota
	OP_DIVMOD   OpType = iota
	OP_DROP     OpType = iota
	OP_PRINT    OpType = iota
	OP_PUSH_INT OpType = iota

	// syscalls
	OP_SYSCALL0 OpType = iota
	OP_SYSCALL1 OpType = iota
	OP_SYSCALL2 OpType = iota
	OP_SYSCALL3 OpType = iota
	OP_SYSCALL4 OpType = iota
	OP_SYSCALL5 OpType = iota
	OP_SYSCALL6 OpType = iota

	// memory
	OP_MEM      OpType = iota
	OP_LOAD8    OpType = iota
	OP_STORE8   OpType = iota
	OP_LOAD16   OpType = iota
	OP_STORE16  OpType = iota
	OP_LOAD32   OpType = iota
	OP_STORE32  OpType = iota
	OP_LOAD64   OpType = iota
	OP_STORE64  OpType = iota

	// stack
	OP_DUP      OpType = iota

	// logic (booleans)
	OP_EQ       OpType = iota
	OP_GT       OpType = iota
	OP_LT       OpType = iota
	OP_GE       OpType = iota
	OP_LE       OpType = iota
	OP_NE       OpType = iota

	// logic (conditions and loops)
	OP_IF       OpType = iota
	OP_ELSE     OpType = iota
	OP_END      OpType = iota
	OP_WHILE    OpType = iota
	OP_DO       OpType = iota

	OP_COUNT    OpType = iota
)

type Operand int
type Op struct {
	op      OpType
	operand Operand
}

func generateYasmLinux_x86_64(program []Op, output string) {
	if !(OP_COUNT == 35) {
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
		f.WriteString("addr_" + strconv.Itoa(i) + ":\n")
		switch program[i].op {
		case OP_PUSH_INT:
			f.WriteString("    ;; -- push --\n"      )
			f.WriteString("    mov rax, " + strconv.Itoa(int(program[i].operand)) + "\n")
			f.WriteString("    push rax\n"           )
		case OP_PLUS:
			f.WriteString("    ;; -- plus --\n"      )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    add rax, rbx\n"       )
			f.WriteString("    push rax\n"           )
		case OP_MINUS:
			f.WriteString("    ;; -- minus --\n"     )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    sub rbx, rax\n"       )
			f.WriteString("    push rbx\n"           )
		case OP_MULT:
			f.WriteString("    ;; -- mult --\n"      )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    mul rbx\n"            )
			f.WriteString("    push rax\n"           )
		case OP_DIVMOD:
			f.WriteString("    ;; -- divmod --\n"    )
			f.WriteString("    xor rdx, rdx\n"       )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    div rbx\n"            )
			f.WriteString("    push rax\n"           )
			f.WriteString("    push rdx\n"           )
		case OP_DROP:
			f.WriteString("    ;; -- drop --\n"      )
			f.WriteString("    pop rax\n"            )
		case OP_PRINT:
			f.WriteString("    ;; -- print --\n"     )
			f.WriteString("    pop rdi\n"            )
			f.WriteString("    call print\n"         )
		case OP_SYSCALL0:
			f.WriteString("    ;; -- syscall0 --\n"  )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    syscall\n"            )
			f.WriteString("    push rax\n"           )
		case OP_SYSCALL1:
			f.WriteString("    ;; -- syscall1 --\n"  )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rdi\n"            )
			f.WriteString("    syscall\n"            )
			f.WriteString("    push rax\n"           )
		case OP_SYSCALL2:
			f.WriteString("    ;; -- syscall2 --\n"  )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rdi\n"            )
			f.WriteString("    pop rsi\n"            )
			f.WriteString("    syscall\n"            )
			f.WriteString("    push rax\n"           )
		case OP_SYSCALL3:
			f.WriteString("    ;; -- syscall3 --\n"  )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rdi\n"            )
			f.WriteString("    pop rsi\n"            )
			f.WriteString("    pop rdx\n"            )
			f.WriteString("    syscall\n"            )
			f.WriteString("    push rax\n"           )
		case OP_SYSCALL4:
			f.WriteString("    ;; -- syscall4 --\n"  )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rdi\n"            )
			f.WriteString("    pop rsi\n"            )
			f.WriteString("    pop rdx\n"            )
			f.WriteString("    pop r10\n"            )
			f.WriteString("    syscall\n"            )
			f.WriteString("    push rax\n"           )
		case OP_SYSCALL5:
			f.WriteString("    ;; -- syscall5 --\n"  )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rdi\n"            )
			f.WriteString("    pop rsi\n"            )
			f.WriteString("    pop rdx\n"            )
			f.WriteString("    pop r10\n"            )
			f.WriteString("    pop r8\n"             )
			f.WriteString("    syscall\n"            )
			f.WriteString("    push rax\n"           )
		case OP_SYSCALL6:
			f.WriteString("    ;; -- syscall6 --\n"  )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rdi\n"            )
			f.WriteString("    pop rsi\n"            )
			f.WriteString("    pop rdx\n"            )
			f.WriteString("    pop r10\n"            )
			f.WriteString("    pop r8\n"             )
			f.WriteString("    pop r9\n"             )
			f.WriteString("    syscall\n"            )
			f.WriteString("    push rax\n"           )
		case OP_MEM:
			f.WriteString("    ;; -- mem --\n"       )
			f.WriteString("    push mem\n"           )
		case OP_LOAD8:
			f.WriteString("    ;; -- load8 --\n"     )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    xor rbx, rbx\n"       )
			f.WriteString("    mov bl, [rax]\n"      )
			f.WriteString("    push rbx\n"           )
		case OP_STORE8:
			f.WriteString("    ;; -- store8 --\n"    )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    mov [rax], bl\n"      )
		case OP_LOAD16:
			f.WriteString("    ;; -- load16 --\n"    )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    xor rbx, rbx\n"       )
			f.WriteString("    mov bx, [rax]\n"      )
			f.WriteString("    push rbx\n"           )
		case OP_STORE16:
			f.WriteString("    ;; -- store16 --\n"   )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    mov [rax], bx\n"      )
		case OP_LOAD32:
			f.WriteString("    ;; -- load32 --\n"    )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    xor rbx, rbx\n"       )
			f.WriteString("    mov ebx, [rax]\n"     )
			f.WriteString("    push rbx\n"           )
		case OP_STORE32:
			f.WriteString("    ;; -- store32 --\n"   )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    mov [rax], ebx\n"     )
		case OP_LOAD64:
			f.WriteString("    ;; -- load64 --\n"    )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    xor rbx, rbx\n"       )
			f.WriteString("    mov rbx, [rax]\n"     )
			f.WriteString("    push rbx\n"           )
		case OP_STORE64:
			f.WriteString("    ;; -- store64 --\n"   )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    mov [rax], rbx\n"     )
		case OP_EQ:
			f.WriteString("    ;; -- equal --\n"     )
			f.WriteString("    mov rcx, 0\n"         )
			f.WriteString("    mov rdx, 1\n"         )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    cmp rax, rbx\n"       )
			f.WriteString("    cmove rcx, rdx\n"     )
			f.WriteString("    push rcx\n"           )
		case OP_GT:
			f.WriteString("    ;; -- grtr than --\n" )
			f.WriteString("    mov rcx, 0\n"         )
			f.WriteString("    mov rdx, 1\n"         )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    cmp rax, rbx\n"       )
			f.WriteString("    cmovg rcx, rdx\n"     )
			f.WriteString("    push rcx\n"           )
		case OP_LT:
			f.WriteString("    ;; -- less than --\n" )
			f.WriteString("    mov rcx, 0\n"         )
			f.WriteString("    mov rdx, 1\n"         )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    cmp rax, rbx\n"       )
			f.WriteString("    cmovl rcx, rdx\n"     )
			f.WriteString("    push rcx\n"           )
		case OP_GE:
			f.WriteString("    ;; -- grtr equl --\n" )
			f.WriteString("    mov rcx, 0\n"         )
			f.WriteString("    mov rdx, 1\n"         )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    cmp rax, rbx\n"       )
			f.WriteString("    cmovge rcx, rdx\n"    )
			f.WriteString("    push rcx\n"           )
		case OP_LE:
			f.WriteString("    ;; -- less equl --\n" )
			f.WriteString("    mov rcx, 0\n"         )
			f.WriteString("    mov rdx, 1\n"         )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    cmp rax, rbx\n"       )
			f.WriteString("    cmovle rcx, rdx\n"    )
			f.WriteString("    push rcx\n"           )
		case OP_NE:
			f.WriteString("    ;; -- not equal --\n" )
			f.WriteString("    mov rcx, 0\n"         )
			f.WriteString("    mov rdx, 1\n"         )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    cmp rax, rbx\n"       )
			f.WriteString("    cmovne rcx, rdx\n"    )
			f.WriteString("    push rcx\n"           )
		case OP_IF:
			f.WriteString("    ;; -- if --\n"        )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    test rax, rax\n"      )
			f.WriteString("    jz addr_" + strconv.Itoa(int(program[i].operand)) + "\n")
		case OP_ELSE:
			f.WriteString("    ;; -- else --\n"      )
			f.WriteString("    jmp addr_" + strconv.Itoa(int(program[i].operand)) + "\n")
		case OP_END:
			f.WriteString("    ;; -- end --\n"       )
			if (i + 1) != int(program[i].operand) {
				f.WriteString("    jmp addr_" + strconv.Itoa(int(program[i].operand)) + "\n")
			}
			f.WriteString("addr_" + strconv.Itoa(i + 1) + ":\n")
		case OP_WHILE:
			f.WriteString("    ;; -- while --\n"     )
		case OP_DO:
			f.WriteString("    ;; -- do --\n"        )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    test rax, rax\n"      )
			f.WriteString("    jz addr_" + strconv.Itoa(int(program[i].operand)) + "\n")
		case OP_DUP:
			f.WriteString("    ;; -- dup --\n"       )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    push rax\n"           )
			f.WriteString("    push rax\n"           )
		default:
			fmt.Fprintf(os.Stderr, "ERROR: Unreachable\n")
			os.Exit(2)
		}
	}
	f.WriteString("    ;; -- built-in exit --\n" )
	f.WriteString("    mov rax, 60\n"            )
	f.WriteString("    mov rdi, 0\n"             )
	f.WriteString("    syscall\n"                )
	f.WriteString("segment .bss\n"               )
	f.WriteString("mem: resb " + strconv.Itoa(MEM_CAP) + "\n")

	f.Close()
}
///////////////////////////////

/////////// PARSER ///////////

func crossreferenceBlocks(program []Op) []Op {
	mprogram := program
	var stack []int
	var ifIp int
	var blockIp int
	var whileIp int
	for i := range mprogram {
		if !(OP_COUNT == 35) {
			fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of ops in crossreferenceBlocks. Add here only operations that form blocks\n")
			os.Exit(1)
		}

		switch mprogram[i].op {
		case OP_IF:
			stack = append(stack, i)
		case OP_ELSE:
			stack, ifIp = popInt(stack)
			mprogram[ifIp].operand = Operand(i + 1)
			stack = append(stack, i)
		case OP_END:
			stack, blockIp = popInt(stack)
			switch {
			case (mprogram[blockIp].op == OP_IF) || (mprogram[blockIp].op == OP_ELSE):
				mprogram[blockIp].operand = Operand(i)
				mprogram[i].operand = Operand(i + 1)
			case mprogram[blockIp].op == OP_DO:
				mprogram[i].operand       = mprogram[blockIp].operand
				mprogram[blockIp].operand = Operand(i + 1)
			}
		case OP_WHILE:
			stack = append(stack, i)
		case OP_DO:
			stack, whileIp = popInt(stack)
			program[i].operand = Operand(whileIp)
			stack = append(stack, i)
		}
	}
	return mprogram
}

func compileTokensIntoOps(tokens []Token) []Op {
	var ops []Op

	if !(TOKEN_COUNT == 2) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of Tokens in compileTokensIntoOps\n")
		os.Exit(1)
	}

	for i := range tokens {
		token := tokens[i]

		switch token.kind {
		case TOKEN_INT:
			ops = append(ops, Op{op: OP_PUSH_INT, operand: Operand(token.icontent)})
		case TOKEN_WORD:
			if !(OP_COUNT == 35) {
				fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of ops in compileTokensIntoOps\n")
				os.Exit(1)
			}

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
			case token.scontent == "syscall0":
				ops = append(ops, Op{op: OP_SYSCALL0})
			case token.scontent == "syscall1":
				ops = append(ops, Op{op: OP_SYSCALL1})
			case token.scontent == "syscall2":
				ops = append(ops, Op{op: OP_SYSCALL2})
			case token.scontent == "syscall3":
				ops = append(ops, Op{op: OP_SYSCALL3})
			case token.scontent == "syscall4":
				ops = append(ops, Op{op: OP_SYSCALL4})
			case token.scontent == "syscall5":
				ops = append(ops, Op{op: OP_SYSCALL5})
			case token.scontent == "syscall6":
				ops = append(ops, Op{op: OP_SYSCALL6})
			case token.scontent == "mem":
				ops = append(ops, Op{op: OP_MEM})
			case token.scontent == "@8":
				ops = append(ops, Op{op: OP_LOAD8})
			case token.scontent == "!8":
				ops = append(ops, Op{op: OP_STORE8})
			case token.scontent == "@16":
				ops = append(ops, Op{op: OP_LOAD16})
			case token.scontent == "!16":
				ops = append(ops, Op{op: OP_STORE16})
			case token.scontent == "@32":
				ops = append(ops, Op{op: OP_LOAD32})
			case token.scontent == "!32":
				ops = append(ops, Op{op: OP_STORE32})
			case token.scontent == "@64":
				ops = append(ops, Op{op: OP_LOAD64})
			case token.scontent == "!64":
				ops = append(ops, Op{op: OP_STORE64})
			case token.scontent == "=":
				ops = append(ops, Op{op: OP_EQ})
			case token.scontent == ">":
				ops = append(ops, Op{op: OP_GT})
			case token.scontent == "<":
				ops = append(ops, Op{op: OP_LT})
			case token.scontent == ">=":
				ops = append(ops, Op{op: OP_GE})
			case token.scontent == "<=":
				ops = append(ops, Op{op: OP_LE})
			case token.scontent == "!=":
				ops = append(ops, Op{op: OP_NE})
			case token.scontent == "if":
				ops = append(ops, Op{op: OP_IF})
			case token.scontent == "else":
				ops = append(ops, Op{op: OP_ELSE})
			case token.scontent == "end":
				ops = append(ops, Op{op: OP_END})
			case token.scontent == "while":
				ops = append(ops, Op{op: OP_WHILE})
			case token.scontent == "do":
				ops = append(ops, Op{op: OP_DO})
			case token.scontent == "dup":
				ops = append(ops, Op{op: OP_DUP})
			default:
				loc := token.loc
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unknown word: %s\n", loc.f, loc.r, loc.c,
					token.scontent)
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
	ops    := crossreferenceBlocks(compileTokensIntoOps(tokens))
	return ops
}

//////////////////////////////
