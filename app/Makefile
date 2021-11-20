APP_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
IMAGE = ghcr.io/hpcsc/argocd-playground-app

##@ Application
.PHONY: test run build-docker
test: ## Run tests
	cd ${APP_DIR} && go test -v ./...

run: ## Run application locally (require Go installed in host machine)
	cd ${APP_DIR} && ENVIRONMENT=local go run .

run-watch: ## Build docker image, run in local cluster and watch for changes using skaffold
	skaffold dev --trigger=manual

build-docker: require-tag ## Build docker image
	cd ${APP_DIR} && docker build -t ${IMAGE}:${TAG} .

push-docker: require-tag ## Push docker image
	docker push ${IMAGE}:${TAG}

update-commit: require-commit ## Set commit field in version.json
	cd ${APP_DIR} && COMMIT=${COMMIT} ./scripts/update-commit.sh
smoke-test: require-tag ## Run smoke test against deployed application
	./scripts/smoke-test.sh $(shell minikube ip):30099 ${TAG}
require-commit:
ifndef COMMIT
	$(error COMMIT variable is required)
endif
