.DEFAULT_GOAL := help
VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")
MANPATH := $(HOME)/.local/share/man
KBDIR := $(HOME)/.klatterburton

# Go related variables.
GOFILES := $(wildcard *.go)

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# Use linker flags to provide version
LDFLAGS=-ldflags "-X=main.Version=$(VERSION)"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# Make sure manpath/man1 is created
_dummy := $(shell mkdir -p $(MANPATH)/man1)

## install: install app, runs 'go install' internally
install:
	go install $(LDFLAGS)
	cp klatter-burton.1.gz $(MANPATH)/man1

## update: updates the installed app to a new version. Basically the same as install without downloading the cert-files
update:
	go install $(LDFLAGS)
	cp klatter-burton.1.gz $(MANPATH)/man1

## build: build binary, runs 'go build' internally
build:
	go build $(LDFLAGS) -o $(PROJECTNAME)

## generate and gzip manppage from markdown, uses pandoc
manpage:
	pandoc klatterburton.md -s -t man -o klatter-burton.1
	gzip klatter-burton.1

## build and interact with docker dev environment
dev:
	docker build -t "kb" .
	docker run -it "kb" sh

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
