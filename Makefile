# SHELL := /bin/bash
TMPDIR := $(if $(TMPDIR),$(TMPDIR),"/tmp/")
GOPATH := $(shell go env GOPATH)

bin := $(GOPATH)/bin/linux-audit-exporter
chart-doc-gen := $(GOPATH)/bin/chart-doc-gen
gofiles := $(wildcard *.go **/*.go **/**/*.go **/**/**/*.go)
golangci-lint := $(GOPATH)/bin/golangci-lint
helm-chart-readme := ./deploy/helm-charts/linux-audit-exporter/README.md

.PHONY: all
all: build docs

.PHONY: build
build: build.source
	
.PHONY: build.source
build.source: $(bin)

.PHONY: docs
docs: docs.helm-chart-readme

.PHONY: docs.helm-chart-readme
docs.helm-chart-readme: $(chart-doc-gen)
	chart-doc-gen \
		-d=./deploy/helm-charts/linux-audit-exporter/doc.yaml \
	       	-v=./deploy/helm-charts/linux-audit-exporter/values.yaml \
	       	>  $(helm-chart-readme)

$(bin): $(gofiles)
	GO111MODULE=on go install ./...

.PHONY: ci
ci: build lint test

.PHONY: clean
clean:
	git clean -x -f

.PHONY: lint 
lint: lint.docs lint.golangci lint.tidy

.PHONY: lint.docs
lint.docs: lint.docs.helm-chart-readme

.PHONY: lint.docs.helm-chart-readme
lint.docs.helm-chart-readme: $(helm-chart-readme)
	git diff --exit-code -- $(helm-chart-readme) > /dev/null

.PHONY: lint.golangci
lint.golangci: $(golangci-lint)
	golangci-lint run ./...

.PHONY: lint.tidy
lint.tidy:
	go mod tidy
	git diff --exit-code -- go.mod go.sum

.PHONY: release.dryrun
release.dryrun:
	npm ci
	npx semantic-release -d

.PHONY: release
release: verify-github-creds
	npm ci
	npx semantic-release --ci=false

.PHONY: test
test: test.unit

.PHONY: test.unit
test.unit:
	GO111MODULE=on go test -cover -race ./...

verify-github-creds:
ifndef GITHUB_ACTOR
	$(error GITHUB_ACTOR environment variable undefined)
endif
ifndef GITHUB_TOKEN
	$(error GITHUB_TOKEN environment variable undefined)
endif

$(chart-doc-gen):
	cd $(TMPDIR) && GO111MODULE=on go get kubepack.dev/chart-doc-gen@v0.3.0

$(golangci-lint):
	cd $(TMPDIR) && GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.38.0
