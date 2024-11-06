
CI_CMD := go run ./cmd/ci/ci.go
WEBSITE_CMD := go run ./cmd/website/website.go

.PHONY: init
init:
	npm ci

.PHONY: dev
dev:
	$(CI_CMD) -dev

.PHONY: run
run:
	$(WEBSITE_CMD)

.PHONY: pr
pr:
	$(CI_CMD) -pr

.PHONY: lint
lint:
	$(CI_CMD) -lint

.PHONY: test
test:
	$(CI_CMD) -test

.PHONY: ci
ci: pr

.PHONY: preview
preview:
	$(CI_CMD) -preview

.PHONY: deploy-prod
deploy-prod:
	$(CI_CMD) -deploy=prod
