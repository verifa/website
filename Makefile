REGION := europe-north1
REPO := $(REGION)-docker.pkg.dev/verifa-website/website
CLOUDRUN_SERVICE := website

AIR_VERSION := v1.49.0
AIR_CMD := go run github.com/cosmtrek/air@$(AIR_VERSION)

TEMPL_VERSION := v0.2.513
TEMPL_CMD := go run github.com/a-h/templ/cmd/templ@$(TEMPL_VERSION)

.PHONY: build-tailwind
build-tailwind:
	tailwindcss build -i ./src/app.css -o ./dist/tailwind.css --minify
dev-tailwind:
	tailwindcss build -i ./src/app.css -o ./dist/tailwind.css --watch



.PHONY: dev
dev:
	$(AIR_CMD) -c .air.toml

.PHONY: generate
generate:
	tailwindcss build -i ./src/app.css -o ./dist/tailwind.css --minify
	$(TEMPL_CMD) generate

.PHONY: preview
preview: generate
	export KO_DOCKER_REPO=$(REPO)
	export IMAGE=$(ko build ./cmd/website --local)
	@echo ""
	@echo ">>>> Built image: ${IMAGE}"
	@echo ""
	@echo ">>>> Starting local container http://localhost:3000"
	@echo ""
	docker run --rm -ti -p 3000:3000 ${IMAGE}

.PHONY: deploy
deploy: generate
	export KO_DOCKER_REPO=$(REPO)
	export IMAGE=$(ko build ./cmd/website)
	@echo ""
	@echo ">>>> Built image: ${IMAGE}"
	@echo ""
	@echo ">>>> Deploying built image to service $(CLOUDRUN_SERVICE) to region $(REGION)"
	@echo ""
	gcloud run deploy $(CLOUDRUN_SERVICE) --image ${IMAGE} --region $(REGION)