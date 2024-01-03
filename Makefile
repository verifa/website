REGION := europe-north1
REPO := $(REGION)-docker.pkg.dev/verifa-website/website
CLOUDRUN_SERVICE := website

.PHONY: build-tailwind
build-tailwind:
	tailwindcss build -i ./src/app.css -o ./dist/tailwind.css --minify
dev-tailwind:
	tailwindcss build -i ./src/app.css -o ./dist/tailwind.css --watch

.PHONY: dev
dev:
	air -c .air.toml

.PHONY: preview
preview:
	go generate .
	export KO_DOCKER_REPO=$(REPO)
	export IMAGE=$(ko build ./cmd/website --local)
	@echo ""
	@echo ">>>> Built image: ${IMAGE}"
	@echo ""
	@echo ">>>> Starting local container http://localhost:3000"
	@echo ""
	docker run --rm -ti -p 3000:3000 ${IMAGE}

.PHONY: deploy
deploy:
	go generate .
	export KO_DOCKER_REPO=$(REPO)
	export IMAGE=$(ko build ./cmd/website)
	@echo ""
	@echo ">>>> Built image: ${IMAGE}"
	@echo ""
	@echo ">>>> Deploying built image to service $(CLOUDRUN_SERVICE) to region $(REGION)"
	@echo ""
	gcloud run deploy $(CLOUDRUN_SERVICE) --image ${IMAGE} --region $(REGION)