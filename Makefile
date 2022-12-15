DEFAULT_GOAL: help

.PHONY: build
build: ## Builds the project
	go build -ldflags="-s -w" github.com/SUNET/nglue-go/cmd/nglue

.PHONY: clean
clean: ## Removes nglue binary from current directory
	rm -rf ./nglue

.PHONY: install
install: ## Installs the nglue binary
	install ./nglue /usr/local/bin/nglue

.PHONY: help
help: ## Prints this message and exits
	@printf "Makefile for building nglue binary\n"
	@printf "Subcommands:\n\n"
	@perl -F':.*##\s+' -lanE '$$F[1] and say "\033[36m$$F[0]\033[0m : $$F[1]"' $(MAKEFILE_LIST) \
		| sort \
		| column -s ':' -t
