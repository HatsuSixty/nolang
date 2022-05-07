package main

////////// GENERATOR //////////
import (
	"fmt"
	"os"
	"strconv"
)

// should be enough for everyone
const MEM_CAP int = 640000

// operations
type OpType int
const (
	// basic operators
	OP_PLUS     OpType = iota
	OP_MINUS    OpType = iota
	OP_MULT     OpType = iota
	OP_DIVMOD   OpType = iota
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
	OP_DROP     OpType = iota
	OP_SWAP     OpType = iota
	OP_OVER     OpType = iota
	OP_ROT      OpType = iota
	OP_2DUP     OpType = iota
	OP_2SWAP    OpType = iota

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

	// bitwise
	OP_SHL      OpType = iota
	OP_SHR      OpType = iota
	OP_OR       OpType = iota

	// others
	OP_ERR      OpType = iota

	OP_COUNT    OpType = iota
)

type Operand int
type Op struct {
	op      OpType
	operand Operand
	loc     Location // only for operations that has a token equivalent
}

func generateYasmLinux_x86_64(program []Op, output string) {
	if !(OP_COUNT == 44) {
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
			if len(program) <= (i + 1) {
				f.WriteString("addr_" + strconv.Itoa(i + 1) + ":\n")
			}
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
		case OP_DROP:
			f.WriteString("    ;; -- drop --\n"      )
			f.WriteString("    pop rax\n"            )
		case OP_SWAP:
			f.WriteString("    ;; -- swap --\n"      )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    push rax\n"           )
			f.WriteString("    push rbx\n"           )
		case OP_OVER:
			f.WriteString("    ;; -- over --\n"      )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    push rbx\n"           )
			f.WriteString("    push rax\n"           )
			f.WriteString("    push rbx\n"           )
		case OP_ROT:
			f.WriteString("    ;; -- rot --\n"       )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    pop rcx\n"            )
			f.WriteString("    push rbx\n"           )
			f.WriteString("    push rax\n"           )
			f.WriteString("    push rcx\n"           )
		case OP_2DUP:
			f.WriteString("    ;; -- 2dup --\n"      )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    push rbx\n"           )
			f.WriteString("    push rax\n"           )
			f.WriteString("    push rbx\n"           )
			f.WriteString("    push rax\n"           )
		case OP_2SWAP:
			f.WriteString("    ;; -- 2swap --\n"     )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    pop rcx\n"            )
			f.WriteString("    pop rdx\n"            )
			f.WriteString("    push rax\n"           )
			f.WriteString("    push rbx\n"           )
			f.WriteString("    push rcx\n"           )
			f.WriteString("    push rdx\n"           )
		case OP_SHL:
			f.WriteString("    ;; -- shl --\n"       )
			f.WriteString("    pop rcx\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    shl rbx, cl\n"        )
			f.WriteString("    push rbx\n"           )
		case OP_SHR:
			f.WriteString("    ;; -- shr --\n"       )
			f.WriteString("    pop rcx\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    shr rbx, cl\n"        )
			f.WriteString("    push rbx\n"           )
		case OP_OR:
			f.WriteString("    ;; -- or --\n"        )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    or rbx, rax\n"        )
			f.WriteString("    push rbx\n"           )
		default:
			fmt.Fprintf(os.Stderr, "ERROR: Unreachable (generateYasmLinux_x86_64)\n")
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
	var ifIp    int
	var blockIp int
	var whileIp int
	for i := range mprogram {
		if !(OP_COUNT == 44) {
			fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of ops in crossreferenceBlocks. Add here only operations that form blocks\n")
			os.Exit(1)
		}

		switch mprogram[i].op {
		case OP_IF:
			stack = append(stack, i)
		case OP_ELSE:
			if !(len(stack) > 0) {
				loc := mprogram[i].loc
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: `else` does not have any block to close\n",
					loc.f, loc.r, loc.c)
				os.Exit(1)
			}

			stack, ifIp = popInt(stack)
			if !(mprogram[ifIp].op == OP_IF) {
				loc := mprogram[i].loc
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Using `else` to close blocks that are not `if` is not allowed\n",
					loc.f, loc.r, loc.c)
				os.Exit(1)
			}

			mprogram[ifIp].operand = Operand(i + 1)
			stack = append(stack, i)
		case OP_END:
			if !(len(stack) > 0) {
				loc := mprogram[i].loc
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: `end` does not have any block to close\n",
					loc.f, loc.r, loc.c)
				os.Exit(1)
			}

			stack, blockIp = popInt(stack)

			switch {
			case (mprogram[blockIp].op == OP_IF) || (mprogram[blockIp].op == OP_ELSE):
				mprogram[blockIp].operand = Operand(i)
				mprogram[i].operand = Operand(i + 1)
			case mprogram[blockIp].op == OP_DO:
				mprogram[i].operand       = mprogram[blockIp].operand
				mprogram[blockIp].operand = Operand(i + 1)
			default:
				loc := mprogram[i].loc
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Using `end` to close blocks that are not `if`, `else`, `do` or `macro` is not allowed\n", loc.f, loc.r, loc.c)
				os.Exit(1)
			}
		case OP_WHILE:
			stack = append(stack, i)
		case OP_DO:
			if !(len(stack) > 0) {
				loc := mprogram[i].loc
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: `do` does not have any block to close\n",
					loc.f, loc.r, loc.c)
				os.Exit(1)
			}

			stack, whileIp = popInt(stack)
			program[i].operand = Operand(whileIp)
			stack = append(stack, i)
		}
	}
	if !(len(stack) == 0) {
		var unclosedBlock int
		stack, unclosedBlock = popInt(stack)
		loc := mprogram[unclosedBlock].loc
		fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unclosed block\n", loc.f, loc.r, loc.c)
		os.Exit(1)
	}

	return mprogram
}

type Macro struct {
	toks []Token
	name string
}

var macros []Macro

func tokenWordAsOp(token Token) Op {
	if !(OP_COUNT == 44) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of ops in tokenWordAsOp\n")
		os.Exit(1)
	}

	switch {
	case token.scontent == "+":
		return Op{op: OP_PLUS, loc: token.loc}
	case token.scontent == "-":
		return Op{op: OP_MINUS, loc: token.loc}
	case token.scontent == "*":
		return Op{op: OP_MULT, loc: token.loc}
	case token.scontent == "divmod":
		return Op{op: OP_DIVMOD, loc: token.loc}
	case token.scontent == "print":
		return Op{op: OP_PRINT, loc: token.loc}
	case token.scontent == "syscall0":
		return Op{op: OP_SYSCALL0, loc: token.loc}
	case token.scontent == "syscall1":
		return Op{op: OP_SYSCALL1, loc: token.loc}
	case token.scontent == "syscall2":
		return Op{op: OP_SYSCALL2, loc: token.loc}
	case token.scontent == "syscall3":
		return Op{op: OP_SYSCALL3, loc: token.loc}
	case token.scontent == "syscall4":
		return Op{op: OP_SYSCALL4, loc: token.loc}
	case token.scontent == "syscall5":
		return Op{op: OP_SYSCALL5, loc: token.loc}
	case token.scontent == "syscall6":
		return Op{op: OP_SYSCALL6, loc: token.loc}
	case token.scontent == "mem":
		return Op{op: OP_MEM, loc: token.loc}
	case token.scontent == "@8":
		return Op{op: OP_LOAD8, loc: token.loc}
	case token.scontent == "!8":
		return Op{op: OP_STORE8, loc: token.loc}
	case token.scontent == "@16":
		return Op{op: OP_LOAD16, loc: token.loc}
	case token.scontent == "!16":
		return Op{op: OP_STORE16, loc: token.loc}
	case token.scontent == "@32":
		return Op{op: OP_LOAD32, loc: token.loc}
	case token.scontent == "!32":
		return Op{op: OP_STORE32, loc: token.loc}
	case token.scontent == "@64":
		return Op{op: OP_LOAD64, loc: token.loc}
	case token.scontent == "!64":
		return Op{op: OP_STORE64, loc: token.loc}
	case token.scontent == "=":
		return Op{op: OP_EQ, loc: token.loc}
	case token.scontent == ">":
		return Op{op: OP_GT, loc: token.loc}
	case token.scontent == "<":
		return Op{op: OP_LT, loc: token.loc}
	case token.scontent == ">=":
		return Op{op: OP_GE, loc: token.loc}
	case token.scontent == "<=":
		return Op{op: OP_LE, loc: token.loc}
	case token.scontent == "!=":
		return Op{op: OP_NE, loc: token.loc}
	case token.scontent == "if":
		return Op{op: OP_IF, loc: token.loc}
	case token.scontent == "else":
		return Op{op: OP_ELSE, loc: token.loc}
	case token.scontent == "end":
		return Op{op: OP_END, loc: token.loc}
	case token.scontent == "while":
		return Op{op: OP_WHILE, loc: token.loc}
	case token.scontent == "do":
		return Op{op: OP_DO, loc: token.loc}
	case token.scontent == "dup":
		return Op{op: OP_DUP, loc: token.loc}
	case token.scontent == "drop":
		return Op{op: OP_DROP, loc: token.loc}
	case token.scontent == "swap":
		return Op{op: OP_SWAP, loc: token.loc}
	case token.scontent == "over":
		return Op{op: OP_OVER, loc: token.loc}
	case token.scontent == "rot":
		return Op{op: OP_ROT, loc: token.loc}
	case token.scontent == "2dup":
		return Op{op: OP_2DUP, loc: token.loc}
	case token.scontent == "2swap":
		return Op{op: OP_2SWAP, loc: token.loc}
	case token.scontent == "shl":
		return Op{op: OP_SHL, loc: token.loc}
	case token.scontent == "shr":
		return Op{op: OP_SHR, loc: token.loc}
	case token.scontent == "or":
		return Op{op: OP_OR, loc: token.loc}
	default:
		return Op{op: OP_ERR}
	}
}

func expandMacro(macro Macro) []Op {
	opers := []Op{}
	if !(TOKEN_COUNT == 2) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of tokens while handling tokens for macro expansion\n")
		os.Exit(1)
	}

	for m := range macro.toks {
		optoadd := tokenWordAsOp(macro.toks[m]) // words
		if optoadd.op == OP_ERR {
			switch {
			case macro.toks[m].kind == TOKEN_INT: // integers
				opers = append(opers, Op{op: OP_PUSH_INT,
					operand: Operand(macro.toks[m].icontent),
					loc: macro.toks[m].loc})
			default:
				err := true
				for mm := range macros { // expand macros that are inside macros
					mn := macros[mm].name
					if mn == macro.toks[m].scontent {
						opers = append(opers, expandMacro(macros[mm])...)
						err = false
						break
					}
				}
				if err {
					loc := macro.toks[m].loc
					fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unknown word: %s\n", loc.f, loc.r, loc.c,
							macro.toks[m].scontent)
					os.Exit(2)
				}
			}
		} else {
			opers = append(opers, optoadd)
		}
	}
	return opers
}

func compileTokensIntoOps(tokens []Token) []Op {
	var ops    []Op

	if !(TOKEN_COUNT == 2) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of Tokens in compileTokensIntoOps\n")
		os.Exit(1)
	}

	i := 0
	for i < len(tokens) {
		token := tokens[i]

		switch token.kind {
		case TOKEN_INT:
			ops = append(ops, Op{op: OP_PUSH_INT, operand: Operand(token.icontent), loc: token.loc})
		case TOKEN_WORD:
			optoadd := tokenWordAsOp(token)
			if optoadd.op == OP_ERR {
				switch {
				case token.scontent == "macro":
					if !((len(tokens)-1) >= i + 1) {
						loc := token.loc
						fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Expected macro name but got nothing\n",
							loc.f, loc.r, loc.c)
						os.Exit(1)
					}

					if !(tokens[i + 1].kind == TOKEN_WORD) {
						loc := tokens[i + 1].loc
						fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Expected macro name to be an word\n",
							loc.f, loc.r, loc.c)
						os.Exit(1)
					}

					macroKeyLoc := token.loc
					macroName   := tokens[i + 1].scontent
					macroToks   := []Token{}
					macroClosed := false

					if in(macroName, builtinWordsNames) {
						fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Redefition of already existing word: %s\n",
							tokens[i+1].loc.f, tokens[i+1].loc.r, tokens[i+1].loc.c, macroName)
						os.Exit(1)
					}

					for m := range macros {
						if macros[m].name == macroName {
							fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Redefition of already existing macro: %s\n",
								tokens[i+1].loc.f, tokens[i+1].loc.r, tokens[i+1].loc.c, macroName)
							os.Exit(1)
						}
					}

					i += 2

					blockStack := []int{}
					pop        := 0
					if pop == 0 {}
					for i < len(tokens) {
						if tokens[i].scontent == "end" && len(blockStack) == 0 {
							macroClosed = true
							break
						}

						if tokens[i].scontent == "macro" {
							fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Creating macros inside macros is not allowed\n",
								tokens[i].loc.f, tokens[i].loc.r, tokens[i].loc.c)
							os.Exit(1)
						}

						if !(OP_COUNT == 44) {
							fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of ops while parsing macro blocks. Add here only operations that are closed by `end`\n")
							os.Exit(1)
						}

						switch {
						case (tokens[i].scontent == "if")  ||
							(tokens[i].scontent == "else") ||
							(tokens[i].scontent == "do"):
							blockStack = append(blockStack, 1)
						case tokens[i].scontent == "end":
							blockStack, pop = popInt(blockStack)
						}

						macroToks = append(macroToks, tokens[i])
						i += 1
					}

					macros = append(macros, Macro{name: macroName, toks: macroToks})

					if !macroClosed {
						fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unclosed block\n", macroKeyLoc.f, macroKeyLoc.r, macroKeyLoc.c)
						os.Exit(1)
					}
				default:
					err := true

					for m := range macros {
						curmac := macros[m]

						if curmac.name == token.scontent {
							ops = append(ops, expandMacro(curmac)...)
							err = false
							break
						}
					}

					if err {
						loc := token.loc
						fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unknown word: %s\n", loc.f, loc.r, loc.c,
							token.scontent)
						os.Exit(1)
					}
				}
			} else {
				ops = append(ops, optoadd)
			}
		default:
			fmt.Fprintf(os.Stderr, "ERROR: Unreachable (compileTokensIntoOps)\n")
			os.Exit(1)
		}
		i += 1
	}
	return crossreferenceBlocks(ops)
}

//////////////////////////////

////////// COMPILER //////////

var builtinWordsNames []string = []string{
	"+",
	"-",
	"*",
	"divmod",
	"print",
	"syscall0",
	"syscall1",
	"syscall2",
	"syscall3",
	"syscall4",
	"syscall5",
	"syscall6",
	"mem",
	"@8",
	"!8",
	"@16",
	"!16",
	"@32",
	"!32",
	"@64",
	"!64",
	"dup",
	"drop",
	"swap",
	"over",
	"rot",
	"2dup",
	"2swap",
	"=",
	">",
	"<",
	">=",
	"<=",
	"!=",
	"if",
	"else",
	"end",
	"while",
	"do",
	"macro",
	"shl",
	"shr",
	"or",
}

func compileFileIntoOps(filepath string) []Op {
	if !(OP_COUNT == 44) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of ops in builtInWordsNames\n")
		os.Exit(1)
	}

	tokens := lexfile(filepath)
	ops    := compileTokensIntoOps(tokens)
	return ops
}

//////////////////////////////
