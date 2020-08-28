ifeq ($(OS), Windows_NT)
	EXE=.exe
else
	EXE=
endif

build-example:
	go build -o out/geode-client$(EXE) example/client.go 

generate:
	go generate ./...

