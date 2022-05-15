package main

////////// GENERATOR //////////
import (
	"fmt"
	"os"
	"strconv"
)

// should be enough for everyone
const MEM_CAP int = 640000

// keywords
type Keyword int
const (
	KEYWORD_MACRO     Keyword = iota
	KEYWORD_CONST     Keyword = iota
	KEYWORD_INCLUDE   Keyword = iota
	KEYWORD_INCREMENT Keyword = iota
	KEYWORD_RESET     Keyword = iota
	KEYWORD_MEMORY    Keyword = iota
	KEYWORD_HERE      Keyword = iota
	KEYWORD_COUNT     Keyword = iota
)

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
	OP_PUSH_STR OpType = iota
	OP_PUSH_MEM OpType = iota

	// syscalls
	OP_SYSCALL0 OpType = iota
	OP_SYSCALL1 OpType = iota
	OP_SYSCALL2 OpType = iota
	OP_SYSCALL3 OpType = iota
	OP_SYSCALL4 OpType = iota
	OP_SYSCALL5 OpType = iota
	OP_SYSCALL6 OpType = iota

	// memory
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
	OP_AND      OpType = iota
	OP_NOT      OpType = iota

	// cmd line args
	OP_ARGV     OpType = iota
	OP_ARGC     OpType = iota

	// others
	OP_ERR      OpType = iota

	OP_COUNT    OpType = iota
)

type Operand int
type OperStr string
type Op struct {
	op      OpType
	operand Operand
	operstr OperStr
	loc     Location // only for operations that has a token equivalent
}

type Ctring struct {
	str string
	id  int
}

func generateYasmLinux_x86_64(program []Op, output string) {
	if !(OP_COUNT == 49) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of ops in generateYasmLinux_x86_64\n")
		os.Exit(1)
	}

	f, err := os.OpenFile(output, os.O_RDWR | os.O_CREATE, 0644)
	if isError(err) {
		os.Exit(3)
	}

	strings := []Ctring{}
	strcnt  := 0

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
	f.WriteString("    mov [args_ptr], rsp\n"              )
	for i := range program {
		f.WriteString("addr_" + strconv.Itoa(i) + ":\n")
		switch program[i].op {
		case OP_PUSH_INT:
			f.WriteString("    ;; -- push --\n"      )
			f.WriteString("    mov rax, " + strconv.Itoa(int(program[i].operand)) + "\n")
			f.WriteString("    push rax\n"           )
		case OP_PUSH_STR:
			f.WriteString("    ;; -- push str --\n"  )
			id := strcnt
			strings = append(strings, Ctring{str: string(program[i].operstr), id: id})
			strcnt += 1
			f.WriteString("    push " + strconv.Itoa(len(program[i].operstr)) + "\n")
			f.WriteString("    push str_" + strconv.Itoa(id) + "\n")
		case OP_PUSH_MEM:
			f.WriteString("    ;; -- push mem --\n"  )
			f.WriteString("    push mem_" + strconv.Itoa(int(program[i].operand)) + "\n")
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
		case OP_AND:
			f.WriteString("    ;; -- and --\n"       )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    pop rbx\n"            )
			f.WriteString("    and rbx, rax\n"       )
			f.WriteString("    push rbx\n"           )
		case OP_NOT:
			f.WriteString("    ;; -- not --\n"       )
			f.WriteString("    pop rax\n"            )
			f.WriteString("    not rax\n"            )
			f.WriteString("    push rax\n"           )
		case OP_ARGV:
			f.WriteString("    ;; -- argv --\n"      )
			f.WriteString("    mov rax, [args_ptr]\n")
			f.WriteString("    add rax, 8\n"         )
			f.WriteString("    push rax\n"           )
		case OP_ARGC:
			f.WriteString("    ;; -- argc --\n"      )
			f.WriteString("    mov rax, [args_ptr]\n")
			f.WriteString("    mov rax, [rax]\n"     )
			f.WriteString("    push rax\n"           )
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
	for mem := range memorys {
		curm := memorys[mem]
		f.WriteString("mem_" + strconv.Itoa(curm.id) + ": resb " + strconv.Itoa(curm.alloc) + "\n")
	}
	f.WriteString("args_ptr: resb 8\n"           )
	f.WriteString("segment .data\n"              )
	for s := range strings {
		curs := strings[s]
		f.WriteString("str_" + strconv.Itoa(curs.id) + ": db ")
		sbytes := []byte(curs.str)
		for sb := range sbytes {
			f.WriteString(strconv.Itoa(int(sbytes[sb])))
			if sb != len(sbytes)-1 {
				f.WriteString(",")
			}
		}
		f.WriteString("\n")
	}

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
		if !(OP_COUNT == 49) {
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
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Using `end` to close blocks that are not `if`, `else`, `do`, `macro` or `const` is not allowed\n", loc.f, loc.r, loc.c)
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

var compiletimecount int = 0
func evaluateAtCompileTime(toks []Token, loc Location) int {
	stack := []int{}
	ret   := 0
	for tk := range toks {
		tok := toks[tk]
		op := tokenWordAsOp(tok).op
		a := 0
		b := 0
		switch {
		case op == OP_PLUS:
			if len(stack) >= 2 {
				stack, a = popInt(stack)
				stack, b = popInt(stack)
				stack = append(stack, a + b)
			}
		case op == OP_MINUS:
			if len(stack) >= 2 {
				stack, a = popInt(stack)
				stack, b = popInt(stack)
				stack = append(stack, b - a)
			}
		case op == OP_MULT:
			if len(stack) >= 2 {
				stack, a = popInt(stack)
				stack, b = popInt(stack)
				stack = append(stack, a * b)
			}
		case op == OP_DIVMOD:
			if len(stack) >= 2 {
				stack, a = popInt(stack)
				stack, b = popInt(stack)
				stack = append(stack, int(b / a))
				stack = append(stack, int(b % a))
			}
		case op == OP_DROP:
			if len(stack) > 0 {
				stack, a = popInt(stack)
			}
		case tok.scontent == "increment":
			if len(stack) < 1 {
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: `increment` requires 1 argument\n",
					tok.loc.f, tok.loc.r, tok.loc.c)
				os.Exit(1)
			}
			stack, a = popInt(stack)
			stack = append(stack, compiletimecount)
			compiletimecount += a
		case tok.scontent == "reset":
			stack = append(stack, compiletimecount)
			compiletimecount = 0
		default:
			if op == OP_ERR {
				err := true
				switch {
				case tok.kind == TOKEN_INT:
					stack = append(stack, tok.icontent)
					err = false
				case tok.kind == TOKEN_WORD:
					for co := 0; co < len(constants); co+=1 {
						curcon := constants[co]
						if curcon.name == tok.scontent {
							stack = append(stack, int(curcon.value))
							err = false
							break
						}
					}
				case tok.kind == TOKEN_STR:
					fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Strings are not allowed inside constant expressions\n",
						loc.f, loc.r, loc.c)
					os.Exit(1)
				}
				if err {
					fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unsupported word at compile time: %s\n",
						tok.loc.f, tok.loc.r, tok.loc.c, tok.scontent)
					os.Exit(1)
				}
			}
		}
	}

	if !(len(stack) == 1) {
		fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Compile time evaluation should produce only 1 result\n",
			loc.f, loc.r, loc.c)
		os.Exit(1)
	}

	ret = stack[0]
	return ret
}

func keywordAsString(key Keyword) string {
	if !(KEYWORD_COUNT == 7) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of keywords in keywordAsString\n")
		os.Exit(1)
	}

	switch key {

	case KEYWORD_CONST:     return "const"
	case KEYWORD_INCLUDE:   return "include"
	case KEYWORD_RESET:     return "reset"
	case KEYWORD_INCREMENT: return "increment"
	case KEYWORD_MACRO:     return "macro"
	case KEYWORD_MEMORY:    return "memory"
	case KEYWORD_HERE:      return "here"

	}
	return "unreachable"
}

func stringAsKeyword(str string) Keyword {
	if !(KEYWORD_COUNT == 7) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of keywords in stringAsKeyword\n")
		os.Exit(1)
	}

	switch str {

	case "const":     return KEYWORD_CONST
	case "include":   return KEYWORD_INCLUDE
	case "reset":     return KEYWORD_RESET
	case "increment": return KEYWORD_INCREMENT
	case "macro":     return KEYWORD_MACRO
	case "memory":    return KEYWORD_MEMORY
	case "here":      return KEYWORD_HERE
	default:          return Keyword(404)

	}
}

func tokenWordAsOp(token Token) Op {
	if !(OP_COUNT == 49) {
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
	case token.scontent == "and":
		return Op{op: OP_AND, loc: token.loc}
	case token.scontent == "not":
		return Op{op: OP_NOT, loc: token.loc}
	case token.scontent == "argv":
		return Op{op: OP_ARGV, loc: token.loc}
	case token.scontent == "argc":
		return Op{op: OP_ARGC, loc: token.loc}
	default:
		return Op{op: OP_ERR}
	}
}

func wordExists(str string) {
	if (tokenWordAsOp(str).op == OP_ERR) && (stringAsKeyword(str) == Keyword(404)) {
		return false
	}
	return true
}

func expandMacro(macro Macro) []Op {
	opers := []Op{}
	if !(TOKEN_COUNT == 4) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of tokens while handling tokens for macro expansion\n")
		os.Exit(1)
	}

	m := 0
	for m < len(macro.toks) {
		switch macro.toks[m].kind {
		case TOKEN_INT:
			opers = append(opers, Op{op: OP_PUSH_INT, operand: Operand(macro.toks[m].icontent), loc: macro.toks[m].loc})
		case TOKEN_STR:
			opers = append(opers, Op{op: OP_PUSH_STR, operstr: OperStr(macro.toks[m].scontent), loc: macro.toks[m].loc})
		case TOKEN_WORD:
			opers = append(opers, handleWord(macro.toks[m])...)
		case TOKEN_KEYWORD:
			m, opers = handleKeyword(m, macro.toks, opers)
		}
		m += 1
	}
	return opers
}

func handleKeyword(i int, tokens []Token, ops []Op) (int, []Op) {
	token := tokens[i]
	if !(KEYWORD_COUNT == 7) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of keywords\n")
		os.Exit(1)
	}

	switch {
	case token.scontent == keywordAsString(KEYWORD_MACRO): // begin macro parsing
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

		checkNameRedefinition(macroName, tokens[i+1].loc)

		i += 2

		blockStack := []int{}
		pop        := 0
		if pop == 0 {}
		for i < len(tokens) {
			if tokenWordAsOp(tokens[i]).op == OP_END && len(blockStack) == 0 {
				macroClosed = true
				break
			}

			// @disallow inside macros
			if tokens[i].scontent == keywordAsString(KEYWORD_MACRO) {
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Creating macros inside macros is not allowed\n",
					tokens[i].loc.f, tokens[i].loc.r, tokens[i].loc.c)
				os.Exit(1)
			}

			if tokens[i].scontent == keywordAsString(KEYWORD_CONST) {
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Creating constants inside macros is not allowed\n",
					tokens[i].loc.f, tokens[i].loc.r, tokens[i].loc.c)
				os.Exit(1)
			}

			if tokens[i].scontent == keywordAsString(KEYWORD_MEMORY) {
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Creating memory blocks inside macros is not allowed\n",
					tokens[i].loc.f, tokens[i].loc.r, tokens[i].loc.c)
				os.Exit(1)
			}


			if !(OP_COUNT == 49) {
				fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of ops while parsing macro blocks. Add here only operations that are closed by `end`\n")
				os.Exit(1)
			}

			switch {
			case tokenWordAsOp(tokens[i]).op == OP_IF ||
				tokenWordAsOp(tokens[i]).op == OP_DO:
				blockStack = append(blockStack, i)
			case tokenWordAsOp(tokens[i]).op == OP_ELSE:
				if !(len(blockStack) > 0) {
					fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: `else` does not have any block to close\n",
						tokens[i].loc.f, tokens[i].loc.r, tokens[i].loc.c)
					os.Exit(1)
				}

				blockStack, pop = popInt(blockStack)
				if !(tokenWordAsOp(tokens[pop]).op == OP_IF) {
					fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Using `else` to close blocks that are not `if` is not allowed\n",
						tokens[i].loc.f, tokens[i].loc.r, tokens[i].loc.c)
					os.Exit(1)
				}

				blockStack = append(blockStack, i)
			case tokenWordAsOp(tokens[i]).op == OP_END:
				if !(len(blockStack) > 0) {
					fmt.Fprintf(os.Stderr, "%s:%s:%d: ERROR: `end` does not have any block to close\n",
						tokens[i].loc.f, tokens[i].loc.r, tokens[i].loc.c)
					os.Exit(1)
				}

				blockStack, pop = popInt(blockStack)

				if !((tokenWordAsOp(tokens[pop]).op == OP_IF)   ||
					(tokenWordAsOp(tokens[pop]).op == OP_ELSE) ||
					(tokenWordAsOp(tokens[pop]).op == OP_DO)) {
					fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Using `end` to close blocks that are not `if`, `else`, `do`, `macro` or `const` is not allowed\n",
						tokens[i].loc.f, tokens[i].loc.r, tokens[i].loc.c)
					os.Exit(1)
				}
			}

			macroToks = append(macroToks, tokens[i])
			i += 1
		}

		macros = append(macros, Macro{name: macroName, toks: macroToks})

		if !macroClosed {
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unclosed block\n", macroKeyLoc.f, macroKeyLoc.r, macroKeyLoc.c)
			os.Exit(1)
		}
	case token.scontent == keywordAsString(KEYWORD_INCLUDE): // end macro parsing
		if !((len(tokens)-1) >= i + 1) { // begin include parsing
			loc := token.loc
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Expected include path but got nothing\n",
				loc.f, loc.r, loc.c)
			os.Exit(1)
		}

		if !(tokens[i + 1].kind == TOKEN_STR) {
			loc := tokens[i + 1].loc
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Expected include path to be an string\n",
				loc.f, loc.r, loc.c)
			os.Exit(1)
		}

		wd, err := os.Getwd()
		if err != nil {}

		includepath := tokens[i + 1].scontent
		pathtofile  := ""

		switch {
		case fileExists(wd + "/" + includepath):       pathtofile = wd + "/" + includepath
		case fileExists(wd + "/std/" + includepath):   pathtofile = wd + "/std/" + includepath
		case fileExists(wd + "/../" + includepath):     pathtofile = wd + "/../" + includepath
		case fileExists(wd + "/../std/" + includepath): pathtofile = wd + "/../std/" + includepath
		default:
			pathtofile = includepath
		}

		includeops := compileFileIntoOps(pathtofile)
		ops = append(ops, includeops...)
		i += 1
	case token.scontent == keywordAsString(KEYWORD_CONST):  // end include parsing
		if !((len(tokens)-1) >= i + 1) { // begin const parsing
			loc := token.loc
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Expected const name but got nothing\n",
				loc.f, loc.r, loc.c)
			os.Exit(1)
		}

		if !(tokens[i + 1].kind == TOKEN_WORD) {
			loc := tokens[i + 1].loc
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Expected const name to be an word\n",
				loc.f, loc.r, loc.c)
			os.Exit(1)
		}

		constKeyLoc := tokens[i].loc
		constName   := tokens[i + 1].scontent
		constToks   := []Token{}
		constClosed := false

		checkNameRedefinition(constName, tokens[i+1].loc)

		i += 2

		for i < len(tokens) {
			if tokens[i].scontent == keywordAsString(KEYWORD_CONST) {
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Creating constants inside constants is not allowed\n",
					tokens[i].loc.f, tokens[i].loc.r, tokens[i].loc.c)
				os.Exit(1)
			}

			if tokenWordAsOp(tokens[i]).op == OP_END {
				constClosed = true
				break
			}

			constToks = append(constToks, tokens[i])
			i += 1
		}

		if !constClosed {
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unclosed block\n", constKeyLoc.f, constKeyLoc.r, constKeyLoc.c)
			os.Exit(1)
		}

		if len(constToks) == 0 {
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Expected expression before closing constant block\n",
				constKeyLoc.f, constKeyLoc.r, constKeyLoc.c)
			os.Exit(1)
		}

		constVal := evaluateAtCompileTime(constToks, constKeyLoc)
		constants = append(constants, Const{name: constName, value: constVal})
	case token.scontent == keywordAsString(KEYWORD_MEMORY): // end const parsing
		if !((len(tokens)-1) >= i + 1) { // begin memory parsing
			loc := token.loc
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Expected memory name but got nothing\n",
				loc.f, loc.r, loc.c)
			os.Exit(1)
		}

		if !(tokens[i + 1].kind == TOKEN_WORD) {
			loc := tokens[i + 1].loc
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Expected memory name to be an word\n",
				loc.f, loc.r, loc.c)
			os.Exit(1)
		}

		memoryKeyLoc := tokens[i].loc
		memoryName   := tokens[i + 1].scontent
		memoryToks   := []Token{}
		memoryClosed := false

		checkNameRedefinition(memoryName, tokens[i+1].loc)

		i += 2

		for i < len(tokens) {
			if tokens[i].scontent == keywordAsString(KEYWORD_MEMORY) {
				fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Creating memory blocks inside memory blocks is not allowed\n",
					tokens[i].loc.f, tokens[i].loc.r, tokens[i].loc.r)
				os.Exit(1)
			}

			if tokenWordAsOp(tokens[i]).op == OP_END {
				memoryClosed = true
				break
			}

			memoryToks = append(memoryToks, tokens[i])
			i += 1
		}

		if !memoryClosed {
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unclosed block\n",
				memoryKeyLoc.f, memoryKeyLoc.r, memoryKeyLoc.c)
			os.Exit(1)
		}

		if len(memoryToks) == 0 {
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Expected expression before closing memory block\n",
				memoryKeyLoc.f, memoryKeyLoc.r, memoryKeyLoc.c)
			os.Exit(1)
		}

		memoryVal := evaluateAtCompileTime(memoryToks, memoryKeyLoc)
		memorys = append(memorys, Memory{name: memoryName, id: memcnt, alloc: memoryVal})
		memcnt += 1
	case token.scontent == keywordAsString(KEYWORD_HERE): // end memory parsing
		// begin here parsing
		loct := token.loc.f + ":" + strconv.Itoa(token.loc.r) + ":" + strconv.Itoa(token.loc.c)
		ops = append(ops, Op{op: OP_PUSH_STR, operstr: OperStr(loct)})
		// end here parsing
	}
	return i, ops
}

func handleWord(token Token) []Op {
	opers := []Op{}
	optoadd := tokenWordAsOp(token)
	if optoadd.op == OP_ERR {
		err   := true

		for m := range macros {
			curmac := macros[m]

			if curmac.name == token.scontent {
				opers = append(opers, expandMacro(curmac)...)
				err = false
				break
			}
		}

		if err {
			for co := range constants {
				curcon := constants[co]

				if curcon.name == token.scontent {
					opers = append(opers, Op{op: OP_PUSH_INT, operand: Operand(curcon.value), loc: token.loc})
					err = false
					break
				}
			}
		}

		if err {
			for mem := range memorys {
				curmem := memorys[mem]

				if curmem.name == token.scontent {
					opers = append(opers, Op{op: OP_PUSH_MEM, operand: Operand(curmem.id), loc: token.loc})
					err = false
					break
				}
			}
		}

		if err {
			loc := token.loc
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Unknown word: %s\n", loc.f, loc.r, loc.c,
				token.scontent)
			os.Exit(1)
		}
	} else {
		opers = append(opers, optoadd)
	}
	return opers
}

func checkNameRedefinition(name string, loc Location) {
	if wordExists(name) {
		fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Redefition of already existing word: %s\n",
			loc.f, loc.r, loc.c, name)
		os.Exit(1)
	}

	for co := range constants {
		if constants[co].name == name {
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Redefition of already existing constant: %s\n",
				loc.f, loc.r, loc.c, name)
			os.Exit(1)
		}
	}

	for mem := range memorys {
		if memorys[mem].name == name {
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Redefition of already existing memory block: %s\n",
				loc.f, loc.r, loc.c, name)
			os.Exit(1)
		}
	}

	for m := range macros {
		if macros[m].name == name {
			fmt.Fprintf(os.Stderr, "%s:%d:%d: ERROR: Redefition of already existing macro: %s\n",
				loc.f, loc.r, loc.c, name)
			os.Exit(1)
		}
	}
}

type Const struct {
	value int
	name  string
}
var constants []Const

type Macro struct {
	toks []Token
	name string
}
var macros []Macro

type Memory struct {
	name  string
	id    int
	alloc int
}
var memorys []Memory
var memcnt  int = 0

func compileTokensIntoOps(tokens []Token) []Op {
	var ops []Op

	if !(TOKEN_COUNT == 4) {
		fmt.Fprintf(os.Stderr, "Assertion Failed: Exhaustive handling of Tokens in compileTokensIntoOps\n")
		os.Exit(1)
	}

	i := 0
	for i < len(tokens) {
		token := tokens[i]

		switch token.kind {
		case TOKEN_INT:
			ops = append(ops, Op{op: OP_PUSH_INT, operand: Operand(token.icontent), loc: token.loc})
		case TOKEN_STR:
			ops = append(ops, Op{op: OP_PUSH_STR, operstr: OperStr(token.scontent), loc: token.loc})
		case TOKEN_WORD:
			ops = append(ops, handleWord(token)...)
		case TOKEN_KEYWORD:
			i, ops = handleKeyword(i, tokens, ops)
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

func compileFileIntoOps(filepath string) []Op {
	tokens := lexfile(filepath)
	ops    := compileTokensIntoOps(tokens)
	return ops
}

//////////////////////////////
