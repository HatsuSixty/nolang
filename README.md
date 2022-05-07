# **Nolang**

[Concatenative](https://concatenative.org) [stack-based](https://en.wikipedia.org/wiki/Stack-oriented_programming) [programming language](https://en.wikipedia.org/wiki/Programming_language) [design](https://en.wikipedia.org/wiki/Design)ed for [writing](https://en.wikipedia.org/wiki/Writing) [programs](https://en.wikipedia.org/wiki/Computer_program) for [computers](https://en.wikipedia.org/wiki/Computer). It definitely is *no*t Forth, but it's written in go*lang*.

# Quick Start

You will need to have the [go](https://go.dev) compiler and the [yasm](https://yasm.tortall.net/) assembler installed.
```console
$ echo "34 35 + print" > sum.no
$ make
$ ./no -c sum.no -r
```
