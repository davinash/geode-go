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
	wget http://apache.org/dyn/closer.cgi/geode/1.12.0/apache-geode-1.12.0.tgz
	ls -l .
	pwd
	tar xvfz apache-geode-1.12.0.tgz
	ls -l .
	GOFLAGS="-count=1" GO111MODULE=on $(GO_TEST) -timeout 50m github.com/davinash/geode-go/tests -v
