REGION := europe-north1
REPO := $(REGION)-docker.pkg.dev/verifa-website/website
CLOUDRUN_SERVICE_PROD := prod-website-service
CLOUDRUN_SERVICE_STAGING := staging-website-service

DOCKER := docker

AIR_CMD := go run -mod=mod github.com/cosmtrek/air

TEMPL_CMD := go run -mod=mod github.com/a-h/templ/cmd/templ

KO_CMD := go run -mod=mod github.com/google/ko

GOLANGCI_LINT_CMD := go run -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint

TAILWIND_CMD := npx tailwindcss

export KO_DOCKER_REPO := $(REPO)

.PHONY: init
init:
	npm ci

.PHONY: build-tailwind
build-tailwind:
	$(TAILWIND_CMD) build -i ./src/app.css -o ./dist/tailwind.css --minify
dev-tailwind:
	$(TAILWIND_CMD) build -i ./src/app.css -o ./dist/tailwind.css --watch

.PHONY: dev
dev:
	$(AIR_CMD) -c .air.toml

.PHONY: run
run:
	$(TAILWIND_CMD) build -i ./src/app.css -o ./dist/tailwind.css --minify
	$(TEMPL_CMD) generate
	go build -o ./build/website ./cmd/website/main.go

.PHONY: generate
generate:
	$(TAILWIND_CMD) build -i ./src/app.css -o ./dist/tailwind.css --minify
	$(TEMPL_CMD) generate

.PHONY: lint
lint:
	$(GOLANGCI_LINT_CMD) run -v --timeout 3m ./...

.PHONY: test
test:
	go test -v -coverpkg=./... ./...

.PHONY: ci
ci: generate lint test

.PHONY: preview
preview: generate
	$(KO_CMD) build ./cmd/website --local | tee image_name.tmp
	@echo ""
	@echo ">>>> Built image: $$(cat image_name.tmp)"
	@echo ""
	@echo ">>>> Starting local container http://localhost:3000"
	@echo ""
	$(DOCKER) run --rm -ti -p 3000:3000 `cat image_name.tmp`
	rm image_name.tmp

.PHONY: build
build:
	$(KO_CMD) build ./cmd/website | tee image_name.tmp
	@echo ""
	@echo ">>>> Built image: $$(cat image_name.tmp)"

.PHONY: deploy-prod
deploy-prod: generate build
	@echo ""
	@echo ">>>> Deploying built image to service $(CLOUDRUN_SERVICE_PROD) to region $(REGION)"
	@echo ""
	gcloud run deploy $(CLOUDRUN_SERVICE_PROD) --image $$(cat image_name.tmp) --region $(REGION)
	@echo ""
	@echo ">>>> Deployed"
	rm image_name.tmp

.PHONY: deploy-staging
deploy-staging: generate build
	@echo ""
	@echo ">>>> Deploying built image to service $(CLOUDRUN_SERVICE_STAGING) to region $(REGION)"
	@echo ""
	gcloud run deploy $(CLOUDRUN_SERVICE_STAGING) --image $$(cat image_name.tmp) --region $(REGION)
	@echo ""
	@echo ">>>> Deployed"
	rm image_name.tmp
