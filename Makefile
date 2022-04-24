GO    = go
FILES = ./src/no.go ./src/compiler.go ./src/common.go ./src/lexer.go

all: no

no: $(FILES)
	$(GO) build $(FILES)
