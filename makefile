GO_TEST=go test
ifeq ($(OS), Windows_NT)
	EXE=.exe
else
	EXE=
endif

build-all:
	go build -v .
	cd geode-func && mvn clean install

build-example:
	go build -o out/geode-client$(EXE) example/client.go 

generate:
	go generate ./...

test-intg: build-all
	GOFLAGS="-count=1" GO111MODULE=on $(GO_TEST) -timeout 50m github.com/davinash/geode-go/tests -v

test-intg-git-flow: build-all
	wget -q http://apachemirror.wuchna.com/geode/1.12.0/apache-geode-1.12.0.tgz
	ls -l .
	pwd
	tar xfz apache-geode-1.12.0.tgz
	ls -l .
	pwd
	which java
	java -version
	GEODE_HOME=$(CURDIR)/apache-geode-1.12.0 GOFLAGS="-count=1" GO111MODULE=on $(GO_TEST) -timeout 50m github.com/davinash/geode-go/tests -v

test-intg-single: build-all
	GOFLAGS="-count=1" GO111MODULE=on $(GO_TEST) -v github.com/davinash/geode-go/tests  -testify.m $(TEST_NAME)
