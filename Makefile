
CI_CMD := go run ./cmd/ci/ci.go

.PHONY: init
init:
	npm ci

.PHONY: dev
dev:
	$(CI_CMD) -dev

.PHONY: run
run:
	$(CI_CMD) -run

.PHONY: pr
pr:
	$(CI_CMD) -pr

.PHONY: lint
lint:
	$(GOLANGCI_LINT_CMD) run -v --timeout 3m ./...

.PHONY: test
test:
	go test -v -coverpkg=./... ./...

.PHONY: ci
ci: generate lint test

.PHONY: preview
preview:
	$(CI_CMD) -preview

.PHONY: deploy-staging
deploy-staging:
	$(CI_CMD) -deploy=staging

.PHONY: deploy-prod
deploy-prod:
	$(CI_CMD) -deploy=prod
