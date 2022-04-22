GO=go

all: no

no: ./src/no.go ./src/compiler.go ./src/common.go
	$(GO) build ./src/no.go ./src/compiler.go ./src/common.go
