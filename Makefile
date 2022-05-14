GO       = go
COMFILES = ./src/no.go ./src/compiler.go ./src/common.go ./src/lexer.go
TESFILES = ./src/test.go ./src/common.go

all: no test

no: $(COMFILES)
	$(GO) build $(COMFILES)

test: $(TESFILES)
	$(GO) build $(TESFILES)
