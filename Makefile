ifeq ($(origin GITROOT), undefined)
GITROOT := $(shell git rev-parse --show-toplevel)
endif

# Commands
GO           ?= go
GOIMPORTS    ?= goimports
GOLANCI_LINT ?= $(or $(wildcard $(GITROOT)/bin/golangci-lint), $(shell which golangci-lint))
GOLINT       ?= $(GOLANCI_LINT) run

GO_IMPORT_PATH = github.com/popodidi/log

define INSTALL_RULE
install-$1:
ifeq (,$(shell which $1))
	$2
else
	@echo "$1 is installed"
endif
endef

define GO_INSTALL_RULE
$(eval $(call INSTALL_RULE,$1,$(GO) get $2))
endef

GO_DEPS = \
	goimports;golang.org/x/tools/cmd/goimports

.PHONY: deps precommit format lint vet test tidy

deps: install-golangci-lint $(foreach DEP,$(GO_DEPS), $(eval CMD = $(word 1,$(subst ;, ,$(DEP)))) install-$(CMD))
$(foreach DEP,$(GO_DEPS), \
	$(eval CMD = $(word 1,$(subst ;, ,$(DEP)))) \
	$(eval SRC = $(word 2,$(subst ;, ,$(DEP)))) \
	$(eval $(call GO_INSTALL_RULE,$(CMD),$(SRC))) \
)

$(eval \
	$(call INSTALL_RULE,golangci-lint, \
	@wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.23.8)\
)

precommit: tidy format lint test

lint:
	$(GOLINT) $(GITROOT)/...

format:
	@find . -name "*.go" | xargs $(GOIMPORTS) -w -local $(GO_IMPORT_PATH)

test:
	$(GO) test $(GITROOT)/...

tidy:
	$(GO) mod tidy
