KUBERNETES_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

##@ Application kubernetes manifests
.PHONY: build-staging-manifests build-prod-manifests deploy-staging deploy-prod

set-image-tag: ## Set kustomize image tag
	pushd ${KUBERNETES_DIR}/base && kustomize edit set image argocd-playground-app=ghcr.io/hpcsc/argocd-playground-app:$(TAG)

build-staging-manifests: ## Build kustomize for staging environment
	kustomize build ${KUBERNETES_DIR}/overlays/staging

build-prod-manifests: ## Build kustomize for production environment
	kustomize build ${KUBERNETES_DIR}/overlays/prod

deploy-staging: ## Build kustomize manifests and apply them in staging environment
	kustomize build ${KUBERNETES_DIR}/overlays/staging | kubectl apply -f -

deploy-prod: ## Build kustomize manifests and apply them in production environment
	kustomize build ${KUBERNETES_DIR}/overlays/prod | kubectl apply -f -
