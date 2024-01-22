REGION := europe-north1
REPO := $(REGION)-docker.pkg.dev/verifa-website/website
CLOUDRUN_SERVICE := website

DOCKER := docker

AIR_VERSION := v1.49.0
AIR_CMD := go run github.com/cosmtrek/air@$(AIR_VERSION)

TEMPL_VERSION := v0.2.513
TEMPL_CMD := go run github.com/a-h/templ/cmd/templ@$(TEMPL_VERSION)

KO_VERSION := v0.15.1
KO_CMD := go run github.com/google/ko@$(KO_VERSION)

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
	$(TEMPL_CMD) generate
	go build -o ./build/website ./cmd/website/main.go

.PHONY: generate
generate:
	$(TAILWIND_CMD) build -i ./src/app.css -o ./dist/tailwind.css --minify
	$(TEMPL_CMD) generate

.PHONY: preview
preview: generate
	@echo ">>>> Starting local container http://localhost:3000"
	@echo ""
	$(DOCKER) run --rm -ti -p 3000:3000 `$(KO_CMD) build ./cmd/website --local`

.PHONY: deploy
deploy: generate
	@echo ">>>> Deploying built image to service $(CLOUDRUN_SERVICE) to region $(REGION)"
	@echo ""
	gcloud run deploy $(CLOUDRUN_SERVICE) --image `$(KO_CMD) build ./cmd/website` --region $(REGION)
