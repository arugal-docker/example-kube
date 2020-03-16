OS = $(shell uname)

GO = go
GO_PATH = $$($(GO) env GOPATH)
GO_GET = $(GO) get
GO_LINT = golangci-lint

.PHONY: lint
lint:
	$(GO_LINT) run -v ./...

.PHONY: fix
fix:
	$(GO_LINT) run -v --fix ./...

.PHONY: deps
deps:
	$(GO_GET) -v -t -d ./...