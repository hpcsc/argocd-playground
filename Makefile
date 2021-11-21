.DEFAULT_GOAL := help

SHELL = bash

include Makefile.utilities
include app/Makefile
include argocd/Makefile
include bootstrap/argocd/Makefile

##@ Help
.PHONY: help
help:  ## Display this help.
	./scripts/list-make-targets.sh $(MAKEFILE_LIST)

##@ Github Actions
.PHONY: act
act-cd:  ## Run github actions CD workflow locally using act
	act -W ./.github/workflows/cd.yaml --secret-file .env.cd

act-argocd-app:  ## Run github actions ArgoCD app workflow locally using act
	act -W ./.github/workflows/argocd-app.yaml --secret-file .env.argocd-app
