REGION := europe-north1
REPO := $(REGION)-docker.pkg.dev/verifa-website/website
CLOUDRUN_SERVICE := website

DOCKER := docker

AIR_CMD := go run -mod=mod github.com/cosmtrek/air

TEMPL_CMD := go run -mod=mod github.com/a-h/templ/cmd/templ

KO_CMD := go run -mod=mod github.com/google/ko

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

.PHONY: deploy
deploy: generate
	$(KO_CMD) build ./cmd/website | tee image_name.tmp
	@echo ""
	@echo ">>>> Built image: $$(cat image_name.tmp)"
	@echo ""
	@echo ">>>> Deploying built image to service $(CLOUDRUN_SERVICE) to region $(REGION)"
	@echo ""
	gcloud run deploy $(CLOUDRUN_SERVICE) --image `cat image_name.tmp` --region $(REGION)
	@echo ""
	@echo ">>>> Deployed"
	rm image_name.tmp
