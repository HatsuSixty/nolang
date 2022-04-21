package main

import (
	"fmt"
	"os"
	"strconv"
)

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

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func generateYasmLinux_x86_64(program []Op, output string) {
	if !(OP_COUNT == 3) {
		println("Assertion Failed: Exhaustive handling of operations in generateyasmlinux_x86_64()")
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
			f.WriteString("    ;; -- push --\n"  )
			f.WriteString("    mov rax, " + strconv.FormatUint(uint64(program[i].operand), 10) + "\n")
			f.WriteString("    push rax\n"       )
		case OP_PLUS:
			f.WriteString("    ;; -- plus --\n"  )
			f.WriteString("    pop rax\n"        )
			f.WriteString("    pop rbx\n"        )
			f.WriteString("    add rax, rbx\n"   )
			f.WriteString("    push rax\n"       )
		case OP_PRINT:
			f.WriteString("    ;; -- print --\n" )
			f.WriteString("    pop rdi\n"        )
			f.WriteString("    call print\n"     )
		}
	}
	f.WriteString("    ;; -- built-in exit --\n" )
	f.WriteString("    mov rax, 60\n"            )
	f.WriteString("    mov rdi, 0\n"             )
	f.WriteString("    syscall\n"                )

	f.Close()
}
