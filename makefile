GO_TEST=go test
ifeq ($(OS), Windows_NT)
	EXE=.exe
else
	EXE=
endif

build-example:
	go build -o out/geode-client$(EXE) example/client.go 

generate:
	go generate ./...

test-intg:
	wget https://github.com/apache/geode/archive/rel/v1.12.0.tar.gz
	ls -l .
	pwd
	tar xvfz v1.12.0.tar.gz
	ls -l .
	GOFLAGS="-count=1" GO111MODULE=on $(GO_TEST) -timeout 50m github.com/davinash/geode-go/tests -v
