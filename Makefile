OS = $(shell uname)

GO = go
GO_PATH = $$($(GO) env GOPATH)
GO_LINT = golangci-lint

.PHONY: lint
lint:
	$(GO_LINT) run -v ./...

.PHONY: fix
fix:
	$(GO_LINT) run -v --fix ./...