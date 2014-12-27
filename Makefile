.PHONY: fmt clean build test release archive

BIN        := mysite

VERSION	   := 0.0.0
COMMIT     := $(shell git rev-parse --short HEAD)

REMOTE     := localhost
REMOTE_DIR := fuga/

GO=go

GOOS   ?= $(shell $(GO) env GOOS)
GOARCH ?= $(shell $(GO) env GOARCH)

ARCHIVE := mysite.$(GOOS)-$(GOARCH).tar.gz
DISTDIR ?= $(CURDIR)/dist/$(GOOS)-$(GOARCH)

LDFLAGS := -ldflags \
           "-X main.Version $(VERSION) \
           -X main.CommitID $(COMMIT)"

clean:
	rm -rf $(BIN) $(DISTDIR)

fmt:
	$(GO) fmt ./...

test:
	$(GO) test ./...

build: $(BIN)

$(BIN): *.go Makefile fmt test
	$(GO) build -o $@ $(LDFLAGS)

archive: dist/$(ARCHIVE)

dist/$(ARCHIVE): $(DISTDIR)/$(BIN)
	tar -C $(DISTDIR) -czvf dit/$(ARCHIVE) .

release: REMOTE     ?= $(error "can't release, REMOTE not set")
release: REMOTE_DIR ?= $(error "can't release, REMOTE_DIR not set")
release: test dist/$(ARCHIVE)
	scp $< $(REMOTE):$(REMOTE_DIR)/