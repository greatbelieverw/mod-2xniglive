TARGET=mod-2xniglive
PACKAGES=core database log parser

.PHONY: all
all: build

build:
	@go build -o ./bin/$(TARGET) -mod=vendor

clean:
	@go clean
	@rm -f ./bin/$(TARGET)

install:
	@mkdir -p /usr/share/mod-2xniglive/phishlets
	@mkdir -p /usr/share/mod-2xniglive/templates
	@cp ./phishlets/* /usr/share/mod-2xniglive/phishlets/
	@cp ./templates/* /usr/share/mod-2xniglive/templates/
	@cp ./bin/$(TARGET) /usr/local/bin
