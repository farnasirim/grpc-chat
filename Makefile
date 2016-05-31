GO ?= go
FORMAT := '{{ join .Deps " " }}'


.PHONY: help clean build dependencies

help:
	@echo "use \`make <target> where <target> is:\'"
	@echo "    help: display this help"
	@echo "    clean: clean the build directory"
	@echo "    build: create the executable"
	@echo "    dependencies: generate protobuf .go file"

clean:
	rm -f proto/chat.pb.go
	rm -f chat

dependencies: proto/chat.proto
	$(GO) get github.com/golang/protobuf/protoc-gen-go
	cd proto; go generate -v .
	$(GO) list -f=$(FORMAT) $(TARGET) | xargs $(GO) install

build: dependencies
	$(GO) build -o "chat"
